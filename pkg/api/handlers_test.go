package api_test

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/maria-robobug/dogfacts/pkg/api"
	"github.com/maria-robobug/dogfacts/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

var (
	app = &api.Application{
		Server:   &http.Server{},
		InfoLog:  log.New(os.Stdin, "", 0),
		ErrorLog: log.New(os.Stderr, "", 0),
	}
)

func TestGetRandomDog_Valid(t *testing.T) {
	// initialise mocks and data
	expected := &api.Response{
		ImageURL:    "https://somecdn.com/images/blah.jpg",
		Name:        "Boston Terrier",
		Height:      "41 - 43 cm at the withers",
		Weight:      "5 - 11 kgs",
		Lifespan:    "11 - 13 years",
		Temperament: "Friendly, Lively, Intelligent",
	}
	mockClient := new(testutils.MockDogApi)
	mockClient.On("GetRandomDogInfo").Return(nil)
	app.Client = mockClient

	// given
	rr, r := makeRequest("GET", "/api/v1/dogs", nil)
	testHandler := http.HandlerFunc(app.GetRandomDog)

	// when
	testHandler.ServeHTTP(rr, r)

	body := &api.Response{}
	err := json.Unmarshal(rr.Body.Bytes(), body)
	if err != nil {
		t.Errorf("unable to read response: %s", err)
	}

	// then
	assert.True(t, rr.Code == http.StatusOK)
	assert.Equal(t, body, expected)
}

func TestGetRandomDog_InternalServerError(t *testing.T) {
	// initialise mocks and data
	mockClient := new(testutils.MockDogApi)
	mockClient.On("GetRandomDogInfo").Return(errors.New("Internal Server Error"))
	app.Client = mockClient

	// given
	rr, r := makeRequest("GET", "/api/v1/dogs", nil)
	testHandler := http.HandlerFunc(app.GetRandomDog)

	// when
	testHandler.ServeHTTP(rr, r)

	// then
	assert.True(t, rr.Code == http.StatusInternalServerError)
}

func makeRequest(method, url string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	return httptest.NewRecorder(), httptest.NewRequest(method, url, body)
}
