package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/todo"
)

func HandlerAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		todo.Add(r.Body)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerRead(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		todo.Read(r.Body)
	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		todo.Update(r.Body)
	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		todo.Delete(r.Body)
	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func main() {
	http.HandleFunc("/add", HandlerAdd)
	http.HandleFunc("/read", HandlerRead)
	http.HandleFunc("/update", HandlerUpdate)
	http.HandleFunc("/delete", HandlerDelete)

	fmt.Println("Server is running on 128.0.0.1:7654")
	log.Fatal(http.ListenAndServe(":7654", nil))
}
