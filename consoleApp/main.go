package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to ToDo Scketch!!!")
	client := http.DefaultClient

	request, err := client.Get("https://privatbank.ua/rates-archive")

	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := io.ReadAll(request.Body)

	defer request.Body.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	fixed_data := html.UnescapeString(string(data))

	// start_index := strings.Index(fixed_data, "")
	// end_index := start_index + 4

	// fmt.Println(string(fixed_data[start_index:end_index]))
	//
	after, _ := strings.CutPrefix(fixed_data, "<span class=\"compare-currency\">UAH</span>")
	before, _ := strings.CutSuffix(after, "</span>")

	fmt.Println(before)
}
