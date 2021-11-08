package main

import (
	"log"
	"net/http"
)

type server struct {}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &server{}))
}
