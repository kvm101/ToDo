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

	fmt.Println("Server is running on: http://localhost:7655")
	log.Fatal(http.ListenAndServe(":7655", nil))
}
