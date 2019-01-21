package testutils

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"

	"github.com/maria-robobug/dogfacts/server/client"
	"github.com/stretchr/testify/mock"
)

var defaultResp = []byte(`[
    {
      "breeds": [
        {
          "bred_for": "Ratting, Companionship",
          "breed_group": "Non-Sporting",
          "height": {
            "imperial": "16 - 17 inches at the withers",
            "metric": "41 - 43 cm at the withers"
          },
          "id": 53,
          "life_span": "11 - 13 years",
          "name": "Boston Terrier",
          "temperament": "Friendly, Lively, Intelligent",
          "weight": {
            "imperial": "10 - 25 lbs",
            "metric": "5 - 11 kgs"
          }
        }
      ],
      "categories": [],
      "id": "rkZRggqVX",
      "url": "https://somecdn.com/images/blah.jpg"
    }
  ]`)

type MockDogAPI struct {
	mock.Mock
}

func (m *MockDogAPI) GetRandomDogInfo() ([]client.DogInfo, error) {
	args := m.Mock.Called()

	d := []client.DogInfo{}
	err := json.Unmarshal(defaultResp, &d)
	if err != nil {
		return d, err
	}

	return d, args.Error(0)
}

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
