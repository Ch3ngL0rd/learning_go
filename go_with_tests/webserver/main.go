package main

import (
	"log"
	"net/http"
)

func main() {
	server := PlayerServer{&StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd": 10,
		},
	}}
	log.Fatal(http.ListenAndServe(":5000", &server))
}