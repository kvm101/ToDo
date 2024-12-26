package todo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func StartDB(host string, port int, user, password, dbname string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database => successfully connected!")

}

func Add() {

}

func Delete() {

}

func Update() {

}

func Read() {

}
