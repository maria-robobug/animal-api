package dogapi

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/maria-robobug/dogfacts/pkg/utils"
)

func TestNewDogClient(t *testing.T) {
	c := &http.Client{
		Timeout: time.Second * 5,
	}
	_, err := NewDogClient("", c)

	if err != nil {
		t.Error("expected client to create successfully")
	}
}

func TestInvalidDogClient(t *testing.T) {
	_, err := NewDogClient("", nil)

	if err == nil {
		t.Error("expected error for invalid client")
	}

	if err.Error() != errInvalidClient.Error() {
		t.Errorf("expected error: %s", errInvalidClient.Error())
	}
}

func TestGetRandomDogInfo(t *testing.T) {
	c := utils.MockClient(200, `[
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
