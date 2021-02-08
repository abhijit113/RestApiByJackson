package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Println("welcome")

		d, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data is %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("goodbye")
	})

	http.ListenAndServe(":9090", nil)
}
