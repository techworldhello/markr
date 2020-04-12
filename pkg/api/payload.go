package api

import (
	"encoding/xml"
	"time"
)

type McqTestResults struct {
	XMLName       xml.Name `xml:"mcq-test-results"`
	Text          string   `xml:",chardata"`
	McqTestResult []struct {
		ScannedOn     time.Time `xml:"scanned-on"`
		FirstName     string    `xml:"first-name"`
		LastName      string    `xml:"last-name"`
		StudentNumber int       `xml:"student-number"`
		TestID        int       `xml:"test-id"`
		SummaryMarks  struct {
			Available int `xml:"available,attr"`
			Obtained  int `xml:"obtained,attr"`
		} `xml:"summary-marks"`
	} `xml:"mcq-test-result"`
}
