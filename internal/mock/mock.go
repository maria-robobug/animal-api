package mock

import (
	"encoding/json"

	"github.com/maria-robobug/animal-api/model"

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
func (a *DogAPI) GetRandomDogInfo() ([]model.DogInfo, error) {
	args := a.Mock.Called()

	d := []model.DogInfo{}
	err := json.Unmarshal([]byte(defaultResp), &d)
	if err != nil {
		return d, err
	}

	return d, args.Error(0)
}

// Cache mocks a generic cache
type Cache struct {
	mock.Mock
}

// Get mocks the method to get a cache key value
func (c *Cache) Get(key string) (interface{}, error) {
	args := c.Mock.Called()

	data := make(map[string]string)
	data["key"] = "value"

	return data, args.Error(0)
}
