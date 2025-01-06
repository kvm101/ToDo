package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {

	var number int

	for {
		number = rand.IntN(100)
		if number > 50 {
			_, err := http.Get("http://localhost:7655/read")
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
