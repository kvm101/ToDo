package todo

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := fmt.Sprint(r.Header)
		user_id := getID(r)

		log.Printf("USER_ID %d %s: %s %s %s %s\n", user_id, r.RemoteAddr, r.Method, r.RequestURI, r.Proto, header[4:(len(header)-1)])
		next.ServeHTTP(w, r)
	})
}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		basic, found := strings.CutPrefix(auth, "Basic ")
		if found == false {
			return
		}

		dbasic, err := b64.StdEncoding.DecodeString(basic)
		if err != nil {
			log.Println(err)
		}

		split_dbase := strings.Split(string(dbasic), ":")
		authentification := authentification(split_dbase[0], split_dbase[1])

		if authentification == false {
			return
		}

		next.ServeHTTP(w, r)
	})
}
