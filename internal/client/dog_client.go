package client

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errInvalidClient = errors.New("invalid client: nil client provided")
)

type DogAPI interface {
	GetRandomDogInfo() ([]DogInfo, error)
}

type DogAPIClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

type DogInfo struct {
	Breeds []Breed `json:"breeds"`
	URL    string  `json:"url"`
}

type Breed struct {
	Name        string  `json:"name"`
	Height      Measure `json:"height"`
	Weight      Measure `json:"weight"`
	BreedGroup  string  `json:"breed_group"`
	LifeSpan    string  `json:"life_span"`
	Temperament string  `json:"temperament"`
}

type Measure struct {
	Metric string `json:"metric"`
}

func New(baseURL, apiKey string, httpClient *http.Client) (*DogAPIClient, error) {
	if httpClient == nil {
		return nil, errInvalidClient
	}

	return &DogAPIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client:  httpClient,
	}, nil
}

func (c *DogAPIClient) GetRandomDogInfo() ([]DogInfo, error) {
	const endpoint = "/images/search?size=small&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return []DogInfo{}, err
	}
	req.Header.Add("x-api-key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return []DogInfo{}, err
	}
	defer resp.Body.Close()

	dogInfo := []DogInfo{}
	err = json.NewDecoder(resp.Body).Decode(&dogInfo)

	return dogInfo, err
}
