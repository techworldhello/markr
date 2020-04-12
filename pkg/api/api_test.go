package api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var body = `
<mcq-test-results>
    <mcq-test-result scanned-on="2017-12-04T12:12:10+11:00">
        <first-name>Jane</first-name>
        <last-name>Austen</last-name>
        <student-number>521585128</student-number>
        <test-id>1234</test-id>
        <summary-marks available="20" obtained="13" />
    </mcq-test-result>
</mcq-test-results>`

func TestStoreResultsReturns200(t *testing.T) {
	recorder := httptest.NewRecorder()

	testRequest, _ := http.NewRequest("POST", "/import", bytes.NewBuffer([]byte(body)))
	testRequest.Header.Add("Content-Type", "text/xml+markr")

	c := Controller{}
	c.storeResults(recorder, testRequest)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, `{"statusCode": 200, "message": "Record successfully saved"}`, recorder.Body.String())
}

func TestStoreResultsFail(t *testing.T) {
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
			request := httptest.NewRequest(expect.protocol, expect.url, bytes.NewBuffer([]byte(body)))
			c.storeResults(recorder, request)

			assert.Equal(t, expect.statusCode, recorder.Code)
			assert.Equal(t, expect.resp, recorder.Body.String())
		})
	}
}
