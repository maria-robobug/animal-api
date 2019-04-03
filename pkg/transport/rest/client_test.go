package rest_test

import (
	"net/http"
	"testing"

	"github.com/maria-robobug/animal-api/pkg/mock"
	"github.com/maria-robobug/animal-api/pkg/transport/rest"
	"github.com/stretchr/testify/assert"
)

const (
	okResponse = `[
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
  ]`
)

func TestNewClient(t *testing.T) {
	_, err := rest.NewClient("", "", &http.Client{})

	assert.Nil(t, err)
}

func TestInvalidClient(t *testing.T) {
	invalidClientError := "invalid client: nil client provided"
	_, err := rest.NewClient("", "", nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), invalidClientError)
}

func TestGetRandomDogInfo(t *testing.T) {
	httpClient, teardown := createTestClient([]byte(okResponse))
	defer teardown()

	expected := []rest.DogInfo{
		{
			Breeds: []rest.Breed{
				rest.Breed{
					"Boston Terrier",
					rest.Measure{Metric: "41 - 43 cm at the withers"},
					rest.Measure{Metric: "5 - 11 kgs"},
					"Non-Sporting",
					"11 - 13 years",
					"Friendly, Lively, Intelligent",
				},
			},
			URL: "https://somecdn.com/images/blah.jpg",
		},
	}

	cli, _ := rest.NewClient("http://test.com", "1234", httpClient)
	body, _ := cli.GetRandomDogInfo()

	assert.Equal(t, expected, body)
}

func createTestClient(resp []byte) (client *http.Client, teardown func()) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(resp)
	})

	return mock.TestingHTTPClient(h)
}
