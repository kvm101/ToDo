package todo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 2345
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type task struct {
	Task_id     int8   `json:"task_id"`
	Head        string `json:"head"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func getDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}

func addTask(head, description string) {
	db := getDB()
	defer db.Close()

	listQuery := `CREATE TABLE IF NOT EXISTS list(
		task_id SERIAL PRIMARY KEY,
		head VARCHAR(50) NOT NULL,
		description VARCHAR(255),
		done BOOLEAN DEFAULT FALSE
	)`

	insertTask := `insert into "list"("head", "description") values($1, $2)`

	_, err := db.Exec(listQuery)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(insertTask, head, description)
	if err != nil {
		log.Println(err)
	}
}

func deleteTask(task_id int8) {
	db := getDB()
	defer db.Close()

	deleteTask := `delete from "list" where task_id=$1`

	_, err := db.Exec(deleteTask, task_id)
	if err != nil {
		log.Println(err)
	}
}

func editTask(task_id int8, head, description string) {
	db := getDB()
	defer db.Close()
	updateTask := `update "list" set "head"=$1, "description"=$2 where "task_id"=$3`

	_, err := db.Exec(updateTask, head, description, task_id)
	if err != nil {
		log.Println(err)
	}
}

func doneTask(task_id int8, done bool) {
	db := getDB()
	defer db.Close()

	doneTask := `update "list" set "done"=$1 where "task_id"=$2`

	_, err := db.Exec(doneTask, done, task_id)
	if err != nil {
		log.Println(err)
	}
}

func readTask(section string) []task {
	var tasks []task = make([]task, 0)
	db := getDB()
	defer db.Close()

	var section_tasks string

	switch section {
	case "all":
		section_tasks = `select * from "list"`
	case "done":
		section_tasks = `select * from "list" where "done"=TRUE`

	case "undone":
		section_tasks = `select * from "list" where "done"=FALSE`

	default:
		section_tasks = `select * from "list"`
	}

	rows, err := db.Query(section_tasks)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table_id int8
		var head string
		var description string
		var done bool
		rows.Scan(&table_id, &head, &description, &done)

		tasks = append(tasks, task{table_id, head, description, done})
	}

	if err != nil {
		log.Println(err)
	}

	return tasks
}
