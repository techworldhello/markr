package api

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func CreateServer(c *Controller) *http.ServeMux {
	server := http.ServeMux{}

	server.HandleFunc("/import", c.storeResults)
	return &server
}

type Controller struct{}

func (c Controller) storeResults(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method != http.MethodPost:
		log.Printf("protocol %s not supported", r.Method)
		handleIncorrectProtocol(w, r)
	default:
		handleResp(w, r)
	}
}

func handleIncorrectProtocol(w http.ResponseWriter, r *http.Request) {
	writeResp(w, http.StatusForbidden, fmt.Sprintf("Protocol %s not supported for endpoint %s", r.Method, r.RequestURI))
}

func handleResp(w http.ResponseWriter, r *http.Request) {
	if ct := r.Header.Get("Content-Type"); ct != "text/xml+markr" {
		writeResp(w, http.StatusUnsupportedMediaType, fmt.Sprintf("Content-Type %s not supported", ct))
		return
	}
	var m McqTestResults
	decoder := xml.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		log.Fatal(err)
	}
	writeResp(w, http.StatusOK, fmt.Sprintf("Record successfully saved"))
}

func writeResp(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, err := fmt.Fprint(w, fmt.Sprintf(`{"statusCode": %d, "message": %s}`, statusCode, message))
	if err != nil {
		log.Fatalf("error writing to stream: %v", err)
	}
}
