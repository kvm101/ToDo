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
	Complexity  byte   `json:"complexity"`
	Importance  byte   `json:"importance"`
	Date        string `json:"date"`
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

func addTask(head, description string, complexity, importance byte) error {
	if complexity < 3 {
		return fmt.Errorf("have(complexity = %d) \nwant (complexity < 3)", complexity)
	}

	if importance < 3 {
		return fmt.Errorf("have(importance = %d) \nwant (importance < 3)", importance)
	}

	db := getDB()
	defer db.Close()

	listQuery := `CREATE TABLE IF NOT EXISTS list(
		task_id SERIAL PRIMARY KEY,
		head VARCHAR(50) NOT NULL,
		description VARCHAR(255),
		done BOOLEAN DEFAULT FALSE,
		complexity NUMERIC(1,0) DEFAULT 0,
		importance NUMERIC(1,0) DEFAULT 0,
		date DATE NOT NULL DEFAULT CURRENT_DATE
	)`

	insertTask := `insert into "list"("head", "description", "complexity", "importance") values($1, $2, $3, $4)`

	_, err := db.Exec(listQuery)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(insertTask, head, description, complexity, importance)
	if err != nil {
		log.Println(err)
	}

	return nil
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

func editTask(task_id int8, head, description string, complexity, importance byte) error {
	if complexity < 3 {
		return fmt.Errorf("have(complexity = %d) \nwant (complexity < 3)", complexity)
	}

	if importance < 3 {
		return fmt.Errorf("have(importance = %d) \nwant (importance < 3)", importance)
	}

	db := getDB()
	defer db.Close()
	updateTask := `update "list" set "head"=$1, "description"=$2, "complexity"=$3, "importance"=$4 where "task_id"=$5`

	_, err := db.Exec(updateTask, head, description, complexity, importance, task_id)
	if err != nil {
		log.Println(err)
	}

	return nil
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

func readTasks(section string, sortf string) []task {
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

	switch sortf {
	case "date_asc":
		section_tasks += ` ORDER BY "date" ASC`

	case "date_desc":
		section_tasks += ` ORDER BY "date" DESC`

	case "head_asc":
		section_tasks += ` ORDER BY "head" ASC`

	case "head_desc":
		section_tasks += ` ORDER BY "head" DESC`

	case "complexity_asc":
		section_tasks += ` ORDER BY "complexity"::NUMERIC ASC`

	case "complexity_desc":
		section_tasks += ` ORDER BY "complexity"::NUMERIC DESC`

	case "importance_asc":
		section_tasks += ` ORDER BY "importance"::NUMERIC ASC`

	case "importance_desc":
		section_tasks += ` ORDER BY "importance"::NUMERIC DESC`

	default:
		section_tasks += ``
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
		var complexity byte
		var importance byte
		var date string
		rows.Scan(&table_id, &head, &description, &done, &complexity, &importance, &date)
		tasks = append(tasks, task{table_id, head, description, done, complexity, importance, date})
	}

	if err != nil {
		log.Println(err)
	}

	return tasks
}
