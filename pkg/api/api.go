package api

import (
	"encoding/xml"
	"fmt"
	"github.com/techworldhello/markr/pkg/data"
	"github.com/techworldhello/markr/pkg/db"
	"log"
	"net/http"
)

func CreateServer(c *Controller) *http.ServeMux {
	server := http.ServeMux{}

	server.HandleFunc("/import", c.saveResult)
	return &server
}

type Controller struct{
	db.Database
}

func (c Controller) saveResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method != http.MethodPost:
		log.Printf("protocol %s not supported", r.Method)
		handleIncorrectProtocol(w, r)
	default:
		c.handleSave(w, r)
	}
}

func handleIncorrectProtocol(w http.ResponseWriter, r *http.Request) {
	writeResp(w, http.StatusForbidden, fmt.Sprintf("Protocol %s not supported for endpoint %s", r.Method, r.RequestURI))
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
		writeResp(w, http.StatusInternalServerError, "Error saving record/s - please try again later.")
		return
	}

	writeResp(w, http.StatusOK, "Record successfully saved")
}
