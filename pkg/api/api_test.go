package api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/markr/pkg/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStore struct {}

func (m MockStore) Save(data.McqTestResults) error {
	return nil
}

func TestSaveResultReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()

	testRequest, _ := http.NewRequest("POST", "/import", bytes.NewBuffer([]byte(data.RequestBody)))
	testRequest.Header.Add("Content-Type", "text/xml+markr")

	c := Controller{MockStore{}}
	c.saveResult(recorder, testRequest)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, `{"statusCode": 200, "message": "Record successfully saved"}`, recorder.Body.String())
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
