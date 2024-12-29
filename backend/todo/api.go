package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandlerAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		var data Task
		json.Unmarshal(body, &data)

		if err != nil {
			fmt.Println(err)
		}

		Add(data.head, data.description)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerRead(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := Read()
		s_data, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprint(w, s_data)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		var data Task
		json.Unmarshal(body, &data)

		if err != nil {
			fmt.Println(err)
		}

		Edit(data.task_id, data.head, data.description)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		body, err := io.ReadAll(r.Body)
		var data Task
		json.Unmarshal(body, &data)

		if err != nil {
			fmt.Println(err)
		}

		Delete(data.task_id)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}
