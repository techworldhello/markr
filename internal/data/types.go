package data

import (
	"time"
)

type McqTestResults struct {
	Results []*TestResult `xml:"mcq-test-result"`
}

type TestResult struct {
	Id int
	CreatedAt time.Time
	ScannedOn     time.Time `xml:"scanned-on,attr"`
	FirstName     string    `xml:"first-name"`
	LastName      string    `xml:"last-name"`
	StudentNumber int       `xml:"student-number"`
	TestID        int       `xml:"test-id"`
	SummaryMarks *SummaryMarks `xml:"summary-marks"`
}

type SummaryMarks struct {
	Available int `xml:"available,attr"`
	Obtained  int `xml:"obtained,attr"`
}

type Aggregate struct {
	Mean   float64 `json:"mean"`
	Stddev float64 `json:"stddev"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	P25    float64 `json:"p25"`
	P50    float64 `json:"p50"`
	P75    float64 `json:"p75"`
	Count  int     `json:"count"`
}
