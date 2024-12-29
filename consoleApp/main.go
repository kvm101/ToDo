package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type task struct {
	task_id     int8
	head        string
	description string
}

func main() {
	fmt.Println("Welcome to ToDo Scketch!!!")
	client := http.DefaultClient
	var resp *http.Response
	var err error
	var data []byte
	var tasks []task

	// for {
	resp, err = client.Get("http://localhost:7655/read")
	data, err = io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &tasks)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range tasks {
		fmt.Println(v)
	}
	fmt.Print("Commands info: \n ")
	// }
}
