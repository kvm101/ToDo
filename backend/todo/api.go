package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var temp_data task

type readFilter struct {
	Section string `json:"section"`
	Sortf   string `json:"sortf"`
}

func readRequest(r *http.Request) *task {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data task
	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err)
	}

	return &data
}

func HandlerAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		temp_data = *readRequest(r)
		err := addTask(temp_data.Head, temp_data.Description, temp_data.Complexity, temp_data.Importance)
		if err != nil {
			log.Println(err)
		}

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerRead(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		var data []byte
		var filter readFilter

		err = json.Unmarshal(body, &filter)
		if err != nil {
			fmt.Println(err)
		}

		tasks := readTasks(filter.Section, filter.Sortf)
		if err != nil {
			log.Println(err)
		}

		data, err = json.Marshal(tasks)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprint(w, string(data))

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		temp_data = *readRequest(r)
		err := editTask(temp_data.Task_id, temp_data.Head, temp_data.Description, temp_data.Complexity, temp_data.Importance)
		if err != nil {
			log.Println(err)
		}

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		temp_data = *readRequest(r)
		deleteTask(temp_data.Task_id)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerDone(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		temp_data = *readRequest(r)
		doneTask(temp_data.Task_id, temp_data.Done)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}
