package dogapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const dogAPIKey = "09b9779d-8aea-4c93-8a55-973cefcfc69f"

type DogInfo []struct {
	ImageURL string  `json:"image"`
	Breeds   []Breed `json:"breeds"`
}

type Breed struct {
	Height struct {
		Metric string `json:"metric"`
	} `json:"height"`
	Weight struct {
		Metric string `json:"metric"`
	} `json:"weight"`
	LifeSpan    string `json:"life_span"`
	Name        string `json:"name"`
	Temperament string `json:"temperament"`
}

type DogClient struct {
	baseURL *url.URL
	client  *http.Client
}

func NewDogClient(httpClient *http.Client) *DogClient {
	return &DogClient{
		baseURL: &url.URL{Path: "https://api.thedogapi.com/v1"},
		client:  httpClient,
	}
}

func (c *DogClient) GetRandomDogInfo() (DogInfo, error) {
	const endpoint = "/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	resp, err := c.client.Get(c.baseURL.Path + endpoint)
	if err != nil {
		return DogInfo{}, err
	}

	dogInfo := DogInfo{}
	json.NewDecoder(resp.Body).Decode(&dogInfo)

	return dogInfo, err
}
