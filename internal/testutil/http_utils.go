package testutils

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// MakeRequest - helper method for making http requests in tests
func MakeRequest(method, url string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	return httptest.NewRecorder(), httptest.NewRequest(method, url, body)
}
