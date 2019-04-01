package mock

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"

	"github.com/maria-robobug/animal-api/pkg/http/rest"
	"github.com/stretchr/testify/mock"
)

var defaultResp = []byte(`[
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
  ]`)

// DogAPI mocks the api
type DogAPI struct {
	mock.Mock
}

// GetRandomDogInfo mocks the method to get a random dog
func (m *DogAPI) GetRandomDogInfo() ([]rest.DogInfo, error) {
	args := m.Mock.Called()

	d := []rest.DogInfo{}
	err := json.Unmarshal(defaultResp, &d)
	if err != nil {
		return d, err
	}

	return d, args.Error(0)
}

// TestingHTTPClient mocks the http client
func TestingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}
