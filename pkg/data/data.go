package data

import (
	"time"
)

type McqTestResults struct {
	Results []*TestResult `xml:"mcq-test-result"`
}

type TestResult struct {
	ScannedOn     time.Time `xml:"scanned-on"`
	FirstName     string    `xml:"first-name"`
	LastName      string    `xml:"last-name"`
	StudentNumber int       `xml:"student-number"`
	TestID        int       `xml:"test-id"`
	*SummaryMarks `xml:"summary-marks"`
}

type SummaryMarks  struct {
	Available int `xml:"available,attr"`
	Obtained  int `xml:"obtained,attr"`
}
