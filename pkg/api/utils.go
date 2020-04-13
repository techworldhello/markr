package api

import (
	"fmt"
	"github.com/techworldhello/markr/pkg/data"
	"log"
	"net/http"
	"time"
)

func writeResp(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, err := fmt.Fprint(w, fmt.Sprintf(`{"statusCode": %d, "message": "%s"}`, statusCode, message))
	if err != nil {
		log.Fatalf("error writing to stream: %v", err)
	}
}

func fieldsAreMissing(m data.McqTestResults) bool {
	if m.Results == nil {
		return true
	}
	emptyTime := time.Time{}
	for _, result := range m.Results {
		if result.StudentNumber == 0 || result.TestID == 0 ||
			result.FirstName == "" || result.LastName == "" ||
			result.SummaryMarks.Obtained == 0 || result.SummaryMarks.Available == 0 ||
			result.ScannedOn == emptyTime {
			return true
		}
	}
	return false
}
