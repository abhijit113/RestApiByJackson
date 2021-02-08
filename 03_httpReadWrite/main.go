package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Println("welcome")

		d, err := ioutil.ReadAll(r.Body)

		if err!=nil{
			/*w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Oops"))
			*/
			http.Error(w,"Oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data is %s", d)

		fmt.Fprintf(w,"Hello, %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("goodbye")
	})

	http.ListenAndServe(":9090", nil)
}
