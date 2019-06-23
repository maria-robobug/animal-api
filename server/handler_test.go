package server_test

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/maria-robobug/animal-api/server"

	"github.com/maria-robobug/animal-api/internal/mock"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomDog_Valid(t *testing.T) {
	// initialise mocks and data
	expected := &server.Response{
		Image: server.Image{
			URL:    "https://somecdn.com/images/blah.jpg",
			Width:  500,
			Height: 200,
		},
		Name:        "Boston Terrier",
		Height:      "41 - 43 cm",
		Weight:      "5 - 11 kgs",
		Lifespan:    "11 - 13 years",
		Temperament: "Friendly, Lively, Intelligent",
		BreedGroup:  "Non-Sporting",
	}

	mockClient := new(mock.DogAPI)
	mockClient.On("GetRandomDogInfo").Return(nil)

	serv := &server.AnimalAPIServer{
		DogAPIClient: mockClient,
		Server:       &http.Server{},
		InfoLog:      log.New(os.Stdin, "", 0),
		ErrorLog:     log.New(os.Stderr, "", 0),
	}

	// given
	rr, r := makeRequest("GET", "/api/v1/dogs/random", nil)
	testHandler := http.HandlerFunc(serv.GetRandomDog)

	// when
	testHandler.ServeHTTP(rr, r)
	body := &server.Response{}
	if err := json.Unmarshal(rr.Body.Bytes(), body); err != nil {
		t.Errorf("unable to read response: %s", err)
	}

	// then
	assert.True(t, rr.Code == http.StatusOK)
	assert.Equal(t, body, expected)
}

func TestGetRandomDog_InternalServerError(t *testing.T) {
	// initialise mocks and data
	mockClient := new(mock.DogAPI)
	mockClient.On("GetRandomDogInfo").Return(errors.New("Internal Server Error"))
	serv := &server.AnimalAPIServer{
		DogAPIClient: mockClient,
		Server:       &http.Server{},
		InfoLog:      log.New(os.Stdin, "", 0),
		ErrorLog:     log.New(os.Stderr, "", 0),
	}

	// given
	rr, r := makeRequest("GET", "/api/v1/dogs/random", nil)
	testHandler := http.HandlerFunc(serv.GetRandomDog)

	// when
	testHandler.ServeHTTP(rr, r)

	// then
	assert.True(t, rr.Code == http.StatusInternalServerError)
}

func makeRequest(method, url string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	return httptest.NewRecorder(), httptest.NewRequest(method, url, body)
}
