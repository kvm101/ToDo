package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type task struct {
	Task_id     int8   `json:"task_id"`
	Head        string `json:"head"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("Welcome to ToDo Scketch!!!")
	client := http.DefaultClient
	var resp *http.Response
	var err error
	var data []byte
	var tasks []task = make([]task, 0)

	// for {
	resp, err = client.Get("http://localhost:7655/read")
	if err != nil {
		log.Println(err)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Println(err)
	}

	for _, v := range tasks {
		fmt.Println(v.Task_id, v.Head, v.Description)
	}

	fmt.Print("Commands info: \n ")
}

// }
