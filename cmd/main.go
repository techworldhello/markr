package main

import (
	"database/sql"
	"github.com/techworldhello/markr/pkg/api"
	"github.com/techworldhello/markr/pkg/db"
	"log"
	"net/http"
)


func initDB() *sql.DB {
	db, err := db.OpenConnection()
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	return db
}

func main() {
	server := api.CreateServer(&api.Controller{*db.New(initDB())})

	log.Print("Starting server on port 4567..")
	if err := http.ListenAndServe(":4567", server); err != nil {
		log.Fatal(err)
	}
}