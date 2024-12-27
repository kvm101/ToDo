package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/todo"
)

func main() {
	http.HandleFunc("/add", todo.HandlerAdd)
	http.HandleFunc("/read", todo.HandlerRead)
	http.HandleFunc("/update", todo.HandlerUpdate)
	http.HandleFunc("/delete", todo.HandlerDelete)

	fmt.Println("Server is running on socket: 128.0.0.1:7654")
	log.Fatal(http.ListenAndServe(":7654", nil))
}
