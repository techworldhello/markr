package api

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/techworldhello/markr/internal/aggregate"
	"github.com/techworldhello/markr/internal/data"
	"github.com/techworldhello/markr/internal/db"
	"net/http"
	"strings"
)

type DatabaseManager interface {
	SaveResults(data data.McqTestResults) error
	RetrieveMarks(testId string) ([]db.DBMarksRecord, error)
}

func CreateServer(c *Controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/import", c.saveResult)
	router.HandleFunc("/results/{test-id}/aggregate", c.getAggregate)
	router.HandleFunc("/status", checkStatus)
	return router
}

type Controller struct{
	DatabaseManager
}

func (c Controller) saveResult(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method != http.MethodPost:
		log.Errorf("protocol %s not supported", r.Method)
		handleIncorrectProtocol(w, r)
	default:
		c.handleSave(w, r)
	}
}

func (c Controller) getAggregate(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method != http.MethodGet:
		log.Errorf("protocol %s not supported", r.Method)
		handleIncorrectProtocol(w, r)
	default:
		c.handleAggregate(w, r)
	}
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, `{"status": "healthy"}`)
	if err != nil {
		log.Errorf("error writing to stream: %v", err)
	}
}

func (c Controller) handleSave(w http.ResponseWriter, r *http.Request) {
	if ct := r.Header.Get("Content-Type"); ct != "text/xml+markr" {
		writeResp(w, http.StatusUnsupportedMediaType, fmt.Sprintf("Content-Type %s not supported", ct))
		return
	}

	var testResults data.McqTestResults

	decoder := xml.NewDecoder(r.Body)
	if err := decoder.Decode(&testResults); err != nil {
		log.Errorf("error unmarshalling request body: %v", err)
	}

	log.Printf("testResults: %+v", testResults.Results)


	if missing := fieldsAreMissing(testResults); missing != false {
		log.Warnf("field/s are missing from result data: %+v", testResults)
		writeResp(w, http.StatusUnprocessableEntity, "Incomplete data - please check that all fields are fulfilled.")
		return
	}

	if err := c.SaveResults(testResults); err != nil {
		log.Errorf("error savings results: %v", err)
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

	records, err := c.RetrieveMarks(testID)
	if err != nil {
		log.Errorf("error retrieving marks for test ID %s: %v", testID, err)
		handleDbProcessingError(w)
		return
	}

	if len(records) == 0 {
		writeResp(w, http.StatusNotFound, fmt.Sprintf("No results were found for Test ID %s", testID))
		return
	}

	writeResultResp(w, aggregate.CalculateAverage(records))
}
