package main

import (
	"github.com/techworldhello/markr/pkg/api"
	"log"
	"net/http"
)

func main() {
	server := api.CreateServer(&api.Controller{})

	log.Print("Starting server on port 4567..")
	if err := http.ListenAndServe(":4567", server); err != nil {
		log.Fatal(err)
	}
}