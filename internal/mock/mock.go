package mock

import (
	"encoding/json"

	"github.com/maria-robobug/animal-api/internal/client"
	"github.com/stretchr/testify/mock"
)

var defaultResp = `[
    {
      "breeds": [
        {
          "bred_for": "Ratting, Companionship",
          "breed_group": "Non-Sporting",
          "height": {
            "imperial": "16 - 17",
            "metric": "41 - 43"
          },
					"id": 53,
					"name": "Boston Terrier",
					"breed_group": "Non-Sporting",
          "life_span": "11 - 13 years",
          "temperament": "Friendly, Lively, Intelligent",
          "weight": {
            "imperial": "10 - 25",
            "metric": "5 - 11"
          }
        }
      ],
      "categories": [],
      "id": "rkZRggqVX",
      "url": "https://somecdn.com/images/blah.jpg"
    }
  ]`

// DogAPI mocks the api
type DogAPI struct {
	mock.Mock
}

// GetRandomDogInfo mocks the method to get a random dog
func (a *DogAPI) GetRandomDogInfo() ([]client.DogInfo, error) {
	args := a.Mock.Called()

	d := []client.DogInfo{}
	err := json.Unmarshal([]byte(defaultResp), &d)
	if err != nil {
		return d, err
	}

	return d, args.Error(0)
}
