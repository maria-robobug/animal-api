package dogapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

// RoundTripFunc .
type RoundTripFunc func(*http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestNewDogClient(t *testing.T) {
	c := &http.Client{
		Timeout: time.Second * 5,
	}
	_, err := NewDogClient(c)

	if err != nil {
		t.Error("expected client to create successfully")
	}
}

func TestInvalidDogClient(t *testing.T) {
	_, err := NewDogClient(nil)

	if err == nil {
		t.Error("expected error for invalid client")
	}

	if err.Error() != errInvalidClient.Error() {
		t.Errorf("expected error: %s", errInvalidClient.Error())
	}
}

func TestGetRandomDogInfo(t *testing.T) {
	c := mockClient(`[
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

	expected := []DogInfo{
		{
			Breeds: []Breed{
				Breed{
					Measure{Metric: "41 - 43 cm at the withers"},
					Measure{Metric: "5 - 11 kgs"},
					"11 - 13 years",
					"Boston Terrier",
					"Friendly, Lively, Intelligent",
				},
			},
			URL: "https://somecdn.com/images/blah.jpg",
		},
	}

	cli := DogClient{&url.URL{Path: "http://dummy.com"}, c}
	body, _ := cli.GetRandomDogInfo()

	if reflect.DeepEqual(body, expected) != true {
		t.Errorf("expected: %+v but got %+v", expected, body)
	}
}

func mockClient(body string) (client *http.Client) {
	client = NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
			Header:     make(http.Header),
		}
	})

	return client
}
