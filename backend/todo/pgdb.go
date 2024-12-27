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

func StartDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func Add(head, description string) {
	db := StartDB()
	defer db.Close()
	insertTask := `insert into "List"("head", "description") values($1, $2)`

	result, err := db.Exec(insertTask, head, description)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(result)
}

func Delete(task_id int8) {
	db := StartDB()
	defer db.Close()
	deleteTask := `delete from "List" where task_id=$1`

	result, err := db.Exec(deleteTask, task_id)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(result)
}

func Edit(task_id int8, head, description string) {
	db := StartDB()
	defer db.Close()
	updateTask := `update "List" set "head=$1", "description"=$2 where "task_id"=$3`

	result, err := db.Exec(updateTask, task_id, head, description)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(result)
}

func Read() {
	db := StartDB()
	defer db.Close()
	readTask := `select * from "List"`

	result, err := db.Exec(readTask)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(result)
}
