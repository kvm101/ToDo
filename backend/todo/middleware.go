package todo

import (
	"fmt"
	"log"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := fmt.Sprint(r.Header)

		log.Printf("user: test %s %s %s %s", r.Method, r.RequestURI, r.Proto, header[4:(len(header)-1)])
		next.ServeHTTP(w, r)
	})
}
