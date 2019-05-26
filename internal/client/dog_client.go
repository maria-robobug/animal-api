package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/maria-robobug/animal-api/model"
)

var (
	errInvalidClient = errors.New("invalid client: nil client provided")
)

type DogAPI interface {
	GetRandomDogInfo() ([]model.DogInfo, error)
}

type DogAPIClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
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

func (c *DogAPIClient) GetRandomDogInfo() ([]model.DogInfo, error) {
	const endpoint = "/images/search?size=small&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return []model.DogInfo{}, err
	}
	req.Header.Add("x-api-key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return []model.DogInfo{}, err
	}
	defer resp.Body.Close()

	dogInfo := []model.DogInfo{}
	err = json.NewDecoder(resp.Body).Decode(&dogInfo)

	return dogInfo, err
}
