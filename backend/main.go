package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/todo"
)

func main() {
	http.Handle("/add", todo.LogRequest(http.HandlerFunc(todo.HandlerAdd)))
	http.Handle("/read", todo.LogRequest(http.HandlerFunc(todo.HandlerRead)))
	http.Handle("/update", todo.LogRequest(http.HandlerFunc(todo.HandlerUpdate)))
	http.Handle("/delete", todo.LogRequest(http.HandlerFunc(todo.HandlerDelete)))

	fmt.Println("Server is running on: http://localhost:7655")
	log.Fatal(http.ListenAndServe(":7655", nil))
}
