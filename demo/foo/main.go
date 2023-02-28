package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("starting foo\n")
	log.Printf("listening on 8080\n")
	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
