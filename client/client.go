package client

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errInvalidClient = errors.New("invalid client: nil client provided")
)

type DogInfo struct {
	Breeds []Breed `json:"breeds"`
	URL    string  `json:"url"`
}

type Breed struct {
	Height      Measure `json:"height"`
	Weight      Measure `json:"weight"`
	LifeSpan    string  `json:"life_span"`
	Name        string  `json:"name"`
	Temperament string  `json:"temperament"`
}

type Measure struct {
	Metric string `json:"metric"`
}

type DogAPI interface {
	GetRandomDogInfo() ([]DogInfo, error)
}

type Client struct {
	BaseURL string
	Client  *http.Client
}

func New(baseURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		return nil, errInvalidClient
	}

	return &Client{
		BaseURL: baseURL,
		Client:  httpClient,
	}, nil
}

func (c *Client) GetRandomDogInfo() ([]DogInfo, error) {
	const endpoint = "/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	resp, err := c.Client.Get(c.BaseURL + endpoint)
	if err != nil {
		return []DogInfo{}, err
	}
	defer resp.Body.Close()

	dogInfo := []DogInfo{}
	json.NewDecoder(resp.Body).Decode(&dogInfo)

	return dogInfo, err
}
