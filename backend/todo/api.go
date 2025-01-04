package todo

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var temp_data task

type readFilter struct {
	Section string `json:"section"`
	Sortf   string `json:"sortf"`
}

func getID(r *http.Request) uint32 {
	auth := r.Header.Get("Authorization")
	basic, found := strings.CutPrefix(auth, "Basic ")
	if found == false {
		return 0
	}

	dbasic, err := b64.StdEncoding.DecodeString(basic)
	if err != nil {
		log.Println(err)
	}

	split_dbase := strings.Split(string(dbasic), ":")
	db := getDB()
	defer db.Close()

	var user_id uint32
	err = db.QueryRow("SELECT user_id from users where login=$1 and password=$2", split_dbase[0], split_dbase[1]).Scan(&user_id)
	if err != nil {
		log.Println(err)
	}

	return user_id
}

func readRequest(r *http.Request) *task {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
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
		user_id := getID(r)

		err := addTask(user_id, temp_data.Head, temp_data.Description, temp_data.Complexity, temp_data.Importance)
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
			log.Println(err)
		}

		user_id := getID(r)
		tasks, err := readTasks(user_id, filter.Section, filter.Sortf)
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

		user_id := getID(r)
		err := editTask(user_id, temp_data.Task_id, temp_data.Head, temp_data.Description, temp_data.Complexity, temp_data.Importance)
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

		user_id := getID(r)
		deleteTask(user_id, temp_data.Task_id)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerDone(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		temp_data = *readRequest(r)

		user_id := getID(r)
		doneTask(user_id, temp_data.Task_id, temp_data.Done)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}
