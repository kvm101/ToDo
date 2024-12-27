package todo

import (
	"fmt"
	"net/http"
)

func HandlerAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		Add()

	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerRead(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Read()
	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		Edit()
	default:
		fmt.Fprint(w, "Not correct request!")
	}
}

func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		Delete()
	default:
		fmt.Fprint(w, "Not correct request!")
	}
}
