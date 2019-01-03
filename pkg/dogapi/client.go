package dogapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

var (
	errInvalidClient = errors.New("invalid client, provide at least custom timeout")
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

type DogClient struct {
	BaseURL *url.URL
	Client  *http.Client
}

func NewDogClient(basePath string, httpClient *http.Client) (*DogClient, error) {
	if httpClient == nil {
		return nil, errInvalidClient
	}

	return &DogClient{
		BaseURL: &url.URL{Path: basePath},
		Client:  httpClient,
	}, nil
}

func (cli *DogClient) GetRandomDogInfo() ([]DogInfo, error) {
	const endpoint = "/images/search?size=med&mime_types=jpg&format=json&has_breeds=true&order=RANDOM&page=0&limit=1"

	resp, err := cli.Client.Get(cli.BaseURL.Path + endpoint)
	if err != nil {
		return []DogInfo{}, err
	}
	defer resp.Body.Close()

	dogInfo := []DogInfo{}
	json.NewDecoder(resp.Body).Decode(&dogInfo)

	return dogInfo, err
}
