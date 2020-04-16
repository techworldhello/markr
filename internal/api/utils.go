package api

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/techworldhello/markr/internal/data"
	"net/http"
	"time"
)

func writeResp(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, err := fmt.Fprint(w, fmt.Sprintf(`{"statusCode": %d, "message": "%s"}`, statusCode, message))
	if err != nil {
		log.Errorf("error writing to stream: %v", err)
	}
}

func writeResultResp(w http.ResponseWriter, result data.Aggregate) {
	resultBytes, err := json.Marshal(&result)
	if err != nil {
		log.Errorf("error marshalling struct to json: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, string(resultBytes))
	if err != nil {
		log.Errorf("error writing to stream: %v", err)
	}
}

func handleIncorrectProtocol(w http.ResponseWriter, r *http.Request) {
	writeResp(w, http.StatusForbidden, fmt.Sprintf("Protocol %s not supported for endpoint %s", r.Method, r.RequestURI))
}

func handleDbProcessingError(w http.ResponseWriter) {
	writeResp(w, http.StatusInternalServerError, "Error processing record/s - please try again later.")
}

func fieldsAreMissing(m data.McqTestResults) bool {
	if m.Results == nil {
		return true
	}
	for _, result := range m.Results {
		if result.StudentNumber == 0 || result.TestID == 0 ||
			result.FirstName == "" || result.LastName == "" ||
			result.SummaryMarks.Obtained == 0 || result.SummaryMarks.Available == 0 ||
			result.ScannedOn == (time.Time{}) {
			return true
		}
	}
	return false
}
