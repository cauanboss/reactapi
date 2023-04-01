package main

import (
	"api/interface/routing"
	"log"
	"net/http"
)

func handleRequest() {
	http.Handle("/", routing.GetRoutes())
	log.Fatal(http.ListenAndServe(":8101", nil))
}

func main() {
	handleRequest()
}
