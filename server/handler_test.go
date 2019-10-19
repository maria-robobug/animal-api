package server_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	clientMock "github.com/maria-robobug/animal-api/internal/mock"
	testutils "github.com/maria-robobug/animal-api/internal/testutil"
	"github.com/maria-robobug/animal-api/server"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomDog(t *testing.T) {
	logger := logrus.New()
	specs := []struct {
		Title string
		Exp   *server.DogResponse
	}{
		{
			Title: "get random dog",
			Exp: &server.DogResponse{
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
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		mockedDogAPI := new(clientMock.DogAPI)
		mockedCatAPI := new(clientMock.CatAPI)
		mockedDogAPI.On("GetRandomDogInfo").Return(nil)
		serv := &server.AnimalAPIServer{
			DogAPIClient: mockedDogAPI,
			CatAPIClient: mockedCatAPI,
			Server:       &http.Server{},
			Logger:       logger,
		}

		for _, spec := range specs {
			rr, r := testutils.MakeRequest("GET", "/api/v1/dogs/random", nil)
			testHandler := http.HandlerFunc(serv.GetRandomDog)
			testHandler.ServeHTTP(rr, r)

			body := &server.DogResponse{}
			if err := json.Unmarshal(rr.Body.Bytes(), body); err != nil {
				t.Fatalf("unable to read response: %s", err)
			}

			mockedDogAPI.AssertNumberOfCalls(t, "GetRandomDogInfo", 1)
			assert.True(t, rr.Code == http.StatusOK)
			assert.Equal(t, body, spec.Exp)
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Run("internal server error", func(t *testing.T) {
			mockedDogAPI := new(clientMock.DogAPI)
			mockedCatAPI := new(clientMock.CatAPI)
			mockedDogAPI.On("GetRandomDogInfo").Return(errors.New("Internal Server Error"))
			serv := &server.AnimalAPIServer{
				DogAPIClient: mockedDogAPI,
				CatAPIClient: mockedCatAPI,
				Server:       &http.Server{},
				Logger:       logger,
			}

			rr, r := testutils.MakeRequest("GET", "/api/v1/dogs/random", nil)
			testHandler := http.HandlerFunc(serv.GetRandomDog)
			testHandler.ServeHTTP(rr, r)

			mockedDogAPI.AssertNumberOfCalls(t, "GetRandomDogInfo", 1)
			assert.True(t, rr.Code == http.StatusInternalServerError)
		})
	})
}
