package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// RoundTripFunc .
type RoundTripFunc func(*http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// MockClient - returns a mock client for testing
func MockClient(status int, body string) (client *http.Client) {
	client = newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: status,
			Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
			Header:     make(http.Header),
		}
	})

	return client
}
