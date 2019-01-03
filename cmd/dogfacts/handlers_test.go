package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/maria-robobug/dogfacts/pkg/dogapi"
	"github.com/maria-robobug/dogfacts/pkg/utils"
)

const baseURL = "http://dummy.com"

var mockClientValid = utils.MockClient(200, `[
  {
    "breeds": [
      {
        "bred_for": "Ratting, Companionship",
        "breed_group": "Non-Sporting",
        "height": {
          "imperial": "16 - 17 inches at the withers",
          "metric": "41 - 43 cm at the withers"
        },
        "id": 53,
        "life_span": "11 - 13 years",
        "name": "Boston Terrier",
        "temperament": "Friendly, Lively, Intelligent",
        "weight": {
          "imperial": "10 - 25 lbs",
          "metric": "5 - 11 kgs"
        }
      }
    ],
    "categories": [],
    "id": "rkZRggqVX",
    "url": "https://somecdn.com/images/blah.jpg"
  }
]`)

var app = &application{
	errorLog: log.New(os.Stderr, "", 0),
	infoLog:  log.New(os.Stdin, "", 0),
}

func TestGetRandomDog_Valid(t *testing.T) {
	app.client = &dogapi.DogClient{
		BaseURL: &url.URL{Path: baseURL},
		Client:  mockClientValid,
	}

	expected := &response{
		ImageURL:    "https://somecdn.com/images/blah.jpg",
		Name:        "Boston Terrier",
		Height:      "41 - 43 cm at the withers",
		Weight:      "5 - 11 kgs",
		Lifespan:    "11 - 13 years",
		Temperament: "Friendly, Lively, Intelligent",
	}

	rec, req, err := makeRequest("GET", "http://localhost:9090/api/v1/dog", nil)
	if err != nil {
		t.Errorf("unable to make request: %v", err)
	}

	app.getRandomDog(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", rec.Code)
	}

	body := &response{}
	err = json.Unmarshal(rec.Body.Bytes(), body)
	if err != nil {
		t.Errorf("unable to read response: %s", err)
	}

	if expected != body {
		t.Errorf("response body does not match; expected: %+v\n got %+v", expected, body)
	}
}

func makeRequest(method, url string, body io.Reader) (*httptest.ResponseRecorder, *http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	return httptest.NewRecorder(), req, nil
}
