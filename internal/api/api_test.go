package api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/markr/internal/data"
	"github.com/techworldhello/markr/internal/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveResultReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()

	testRequest, _ := http.NewRequest("POST", "/import", bytes.NewBuffer([]byte(data.RequestBody)))
	testRequest.Header.Add("Content-Type", "text/xml+markr")

	c := Controller{MockStore{}}
	c.saveResult(recorder, testRequest)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, `{"statusCode": 200, "message": "Record/s successfully saved"}`, recorder.Body.String())
}

func TestSaveResultWithIncompleteReqBody(t *testing.T) {
	recorder := httptest.NewRecorder()

	testRequest, _ := http.NewRequest("POST", "/import", bytes.NewBuffer([]byte(data.IncompleteRequestBody)))
	testRequest.Header.Add("Content-Type", "text/xml+markr")

	c := Controller{MockStore{}}
	c.saveResult(recorder, testRequest)

	assert.Equal(t, 422, recorder.Code)
	assert.Equal(t, `{"statusCode": 422, "message": "Incomplete data - please check that all fields are fulfilled."}`, recorder.Body.String())
}

func TestSaveResultFail(t *testing.T) {
	expectations := []struct {
		name       string
		url        string
		protocol   string
		statusCode int
		resp       string
	}{
		{
			name:       "missing_content_type",
			url:        "/import",
			protocol:   "POST",
			statusCode: 415,
			resp:       `{"statusCode": 415, "message": "Content-Type  not supported"}`,
		},
		{
			name:       "unsupported_protocol",
			url:        "/import",
			protocol:   "GET",
			statusCode: 403,
			resp:       `{"statusCode": 403, "message": "Protocol GET not supported for endpoint /import"}`,
		},
	}
	c := Controller{}
	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(expect.protocol, expect.url, bytes.NewBuffer([]byte(data.RequestBody)))
			c.saveResult(recorder, request)

			assert.Equal(t, expect.statusCode, recorder.Code)
			assert.Equal(t, expect.resp, recorder.Body.String())
		})
	}
}

func TestGetAggregateReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()

	testRequest, _ := http.NewRequest("GET", "/results/1234/aggregate", nil)

	c := Controller{MockStore{}}
	c.getAggregate(recorder, testRequest)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, `{"mean":65,"stddev":0,"min":65,"max":65,"p25":65,"p50":65,"p75":65,"count":1}`, recorder.Body.String())
}

type MockStore struct {}

func (m MockStore) SaveResults(data data.McqTestResults) error {
	return nil
}

func (m MockStore) RetrieveMarks(testId string) ([]db.DBMarksRecord, error) {
	return []db.DBMarksRecord{{"1234", 20, 13}}, nil
}
