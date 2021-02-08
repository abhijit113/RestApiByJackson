package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/welcome", func(http.ResponseWriter, *http.Request) {
		log.Println("welcome")
		})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye")
		})

	http.ListenAndServe(":9090", nil)
}
