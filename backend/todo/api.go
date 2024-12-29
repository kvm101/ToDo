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
		var data task
		json.Unmarshal(body, &data)

		if err != nil {
			fmt.Println(err)
		}

		Add(data.Head, data.Description)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerRead(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := json.Marshal()
		if err != nil {
			log.Println(err)
		}

		fmt.Fprint(w, data)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		body, err := io.ReadAll(r.Body)
		var data task
		json.Unmarshal(body, &data)

		if err != nil {
			fmt.Println(err)
		}

		Edit(data.Task_id, data.Head, data.Description)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		body, err := io.ReadAll(r.Body)
		var data task
		json.Unmarshal(body, &data)

		if err != nil {
			fmt.Println(err)
		}

		Delete(data.Task_id)

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}
