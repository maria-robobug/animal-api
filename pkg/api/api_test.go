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

func TestNewApplication(t *testing.T) {
	infoLog := log.New(os.Stdin, "", 0)
	errorLog := log.New(os.Stderr, "", 0)

	// given
	mockClient := new(testutils.MockDogApi)

	// when
	app := api.New(mockClient, ":9000", infoLog, errorLog)

	// then
	assert.NotNil(t, app)
}

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
	app := &api.Application{
		Client:   mockClient,
		Server:   &http.Server{},
		InfoLog:  log.New(os.Stdin, "", 0),
		ErrorLog: log.New(os.Stderr, "", 0),
	}

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
	app := &api.Application{
		Client:   mockClient,
		Server:   &http.Server{},
		InfoLog:  log.New(os.Stdin, "", 0),
		ErrorLog: log.New(os.Stderr, "", 0),
	}

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
