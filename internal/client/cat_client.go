package client

import (
	"encoding/json"
	"net/http"
)

type CatAPI interface {
	GetRandomCatInfo() ([]CatInfo, error)
}

type CatAPIClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

type CatInfo struct {
	Breeds []CatBreed `json:"breeds"`
	URL    string     `json:"url"`
	Width  int64      `json:"width"`
	Height int64      `json:"height"`
}

type CatBreed struct {
	Name        string  `json:"name"`
	Weight      Measure `json:"weight"`
	Lifespan    string  `json:"life_span"`
	Temperament string  `json:"temperament"`
	Description string  `json:"description"`
}

func NewCatAPI(baseURL, apiKey string, httpClient *http.Client) (*CatAPIClient, error) {
	if httpClient == nil {
		return nil, errInvalidClient
	}

	return &CatAPIClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client:  httpClient,
	}, nil
}

func (c *CatAPIClient) GetRandomCatInfo() ([]CatInfo, error) {
	const endpoint = "/images/search?size=small&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return []CatInfo{}, err
	}
	req.Header.Add("x-api-key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return []CatInfo{}, err
	}
	defer resp.Body.Close()

	catInfo := []CatInfo{}
	err = json.NewDecoder(resp.Body).Decode(&catInfo)

	return catInfo, err
}
