package main

import (
	"database/sql"
	"github.com/techworldhello/markr/internal/api"
	"github.com/techworldhello/markr/internal/db"
	"log"
	"net/http"
)

var sqlDb *sql.DB

func init() {
	db, err := db.OpenConnection()
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	sqlDb = db
}

func main() {
	server := api.CreateServer(&api.Controller{db.New(sqlDb)})

	log.Print("Starting server on port 4567..")
	if err := http.ListenAndServe(":4567", server); err != nil {
		log.Fatal(err)
	}
}