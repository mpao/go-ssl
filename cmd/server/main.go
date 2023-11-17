package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	log.Println("starting server")
	if err := http.ListenAndServeTLS(":8080", "cert/localhost.crt", "cert/localhost.key", nil); err != nil {
		log.Fatal(err)
	}
}
