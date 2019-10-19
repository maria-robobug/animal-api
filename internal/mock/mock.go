package mock

import (
	"encoding/json"

	"github.com/maria-robobug/animal-api/internal/client"
	"github.com/stretchr/testify/mock"
)

var (
	defaultDogResp = `[
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
      "url": "https://somecdn.com/images/blah.jpg",
      "width": 500,
      "height": 200
    }
  ]`

	defaultCatResp = `[
    {
        "breeds": [
            {
                "weight": {
                    "imperial": "5 - 11",
                    "metric": "2 - 5"
                },
                "id": "rblu",
                "name": "Russian Blue",
                "cfa_url": "http://cfa.org/Breeds/BreedsKthruR/RussianBlue.aspx",
                "vetstreet_url": "http://www.vetstreet.com/cats/russian-blue-nebelung",
                "vcahospitals_url": "https://vcahospitals.com/know-your-pet/cat-breeds/russian-blue",
                "temperament": "Active, Dependent, Easy Going, Gentle, Intelligent, Loyal, Playful, Quiet",
                "origin": "Russia",
                "country_codes": "RU",
                "country_code": "RU",
                "description": "Russian Blues are very loving and reserved. They do not like noisy households but they do like to play and can be quite active when outdoors. They bond very closely with their owner and are known to be compatible with other pets.",
                "life_span": "10 - 16",
                "indoor": 0,
                "lap": 1,
                "alt_names": "Archangel Blue, Archangel Cat",
                "adaptability": 3,
                "affection_level": 3,
                "child_friendly": 3,
                "dog_friendly": 3,
                "energy_level": 3,
                "grooming": 3,
                "health_issues": 1,
                "intelligence": 3,
                "shedding_level": 3,
                "social_needs": 3,
                "stranger_friendly": 1,
                "vocalisation": 1,
                "experimental": 0,
                "hairless": 0,
                "natural": 1,
                "rare": 0,
                "rex": 0,
                "suppressed_tail": 0,
                "short_legs": 0,
                "wikipedia_url": "https://en.wikipedia.org/wiki/Russian_Blue",
                "hypoallergenic": 1
            }
        ],
        "id": "zK-7AFYVn",
        "url": "https://cdn2.thecatapi.com/images/zK-7AFYVn.jpg",
        "width": 1080,
        "height": 1080
    }
]`
)

// DogAPI mocks the dog api
type DogAPI struct {
	mock.Mock
}

// CatAPI mocks the cat api
type CatAPI struct {
	mock.Mock
}

// GetRandomDogInfo mocks the method to get a random dog
func (a *DogAPI) GetRandomDogInfo() ([]client.DogInfo, error) {
	args := a.Mock.Called()

	d := []client.DogInfo{}
	err := json.Unmarshal([]byte(defaultDogResp), &d)
	if err != nil {
		return d, err
	}

	return d, args.Error(0)
}

func (a *CatAPI) GetRandomCatInfo() ([]client.CatInfo, error) {
	args := a.Mock.Called()

	c := []client.CatInfo{}
	err := json.Unmarshal([]byte(defaultCatResp), &c)
	if err != nil {
		return c, err
	}

	return c, args.Error(0)
}
