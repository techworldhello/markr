package api

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/techworldhello/markr/internal/data"
	"github.com/techworldhello/markr/internal/db"
	"log"
	"net/http"
	"strings"
)

func CreateServer(c *Controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/import", c.saveResult)
	router.HandleFunc("/results/{test-id}/aggregate", c.getAggregate)
	return router
}

type Controller struct{
	db.Database
}

func (c Controller) saveResult(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method != http.MethodPost:
		log.Printf("protocol %s not supported", r.Method)
		handleIncorrectProtocol(w, r)
	default:
		c.handleSave(w, r)
	}
}

func (c Controller) getAggregate(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method != http.MethodGet:
		log.Printf("protocol %s not supported", r.Method)
		handleIncorrectProtocol(w, r)
	default:
		c.handleAggregate(w, r)
	}
}

func (c Controller) handleSave(w http.ResponseWriter, r *http.Request) {
	if ct := r.Header.Get("Content-Type"); ct != "text/xml+markr" {
		writeResp(w, http.StatusUnsupportedMediaType, fmt.Sprintf("Content-Type %s not supported", ct))
		return
	}

	var data data.McqTestResults

	decoder := xml.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Fatal(err)
	}

	if missing := fieldsAreMissing(data); missing != false {
		log.Printf("field/s are missing from result data: %+v", data)
		writeResp(w, http.StatusUnprocessableEntity, "Incomplete data - please check that all fields are fulfilled.")
		return
	}

	if err := c.Save(data); err != nil {
		log.Printf("error savings results: %v", err)
		handleDbProcessingError(w)
		return
	}

	writeResp(w, http.StatusOK, "Record successfully saved")
}

func (c Controller) handleAggregate(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	testID := p[2]
	if testID == "" {
		writeResp(w, http.StatusUnprocessableEntity, "Test ID must be supplied in params.")
		return
	}

	scores, err := c.RetrieveScores(testID)
	if err != nil {
		handleDbProcessingError(w)
		return
	}

	if len(scores) == 0 {
		writeResp(w, http.StatusNotFound, fmt.Sprintf("No results were found for Test ID %s", testID))
		return
	}

	writeResults(w, scores)
}
