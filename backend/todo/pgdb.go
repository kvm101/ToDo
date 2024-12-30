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
}

func getDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database => successfully connected!")
	return db
}

func Add(head, description string) {
	db := getDB()
	defer db.Close()

	listQuery := `CREATE TABLE IF NOT EXISTS list(
		task_id SERIAL PRIMARY KEY,
		head VARCHAR(50) NOT NULL,
		description VARCHAR(255)
	)`

	insertTask := `insert into "list"("head", "description") values($1, $2)`

	_, err := db.Exec(listQuery)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec(insertTask, head, description)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Delete(task_id int8) {
	db := getDB()
	defer db.Close()
	deleteTask := `delete from "list" where task_id=$1`

	_, err := db.Exec(deleteTask, task_id)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Edit(task_id int8, head, description string) {
	db := getDB()
	defer db.Close()
	updateTask := `update "list" set "head"=$1, "description"=$2 where "task_id"=$3`

	_, err := db.Exec(updateTask, head, description, task_id)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Read() []task {
	var tasks []task = make([]task, 0)
	db := getDB()
	defer db.Close()

	rows, err := db.Query(`select * from "list"`)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var table_id int8
		var head string
		var description string
		rows.Scan(&table_id, &head, &description)

		tasks = append(tasks, task{table_id, head, description})
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)
	return tasks
}
