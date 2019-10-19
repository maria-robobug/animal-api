package server_test

import (
	"testing"

	"github.com/maria-robobug/animal-api/internal/mock"
	"github.com/maria-robobug/animal-api/server"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Run("valid config", func(t *testing.T) {
			mockDogClient := new(mock.DogAPI)
			mockCatClient := new(mock.CatAPI)
			logger := logrus.New()
			cnfg := &server.Config{
				DogAPIClient: mockDogClient,
				CatAPIClient: mockCatClient,
				Addr:         ":9000",
				Logger:       logger,
			}

			serv, err := server.New(cnfg)

			assert.Nil(t, err, "error is not nil")
			assert.NotNil(t, serv.DogAPIClient, "dog-api client is nil")
			assert.NotNil(t, serv.CatAPIClient, "cat-api client is nil")
			assert.NotNil(t, serv.Logger, "logger is nil")
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("missing dog-api client", func(t *testing.T) {
			logger := logrus.New()
			cnfg := &server.Config{
				CatAPIClient: new(mock.CatAPI),
				Addr:         ":9000",
				Logger:       logger,
			}

			_, err := server.New(cnfg)

			assert.NotNil(t, err, "error is nil")
			assert.Equal(t, err.Error(), "invalid config: nil dog-api client")
		})

		t.Run("missing cat-api client", func(t *testing.T) {
			logger := logrus.New()
			cnfg := &server.Config{
				DogAPIClient: new(mock.DogAPI),
				Addr:         ":9000",
				Logger:       logger,
			}

			_, err := server.New(cnfg)

			assert.NotNil(t, err, "error is nil")
			assert.Equal(t, err.Error(), "invalid config: nil cat-api client")
		})

		t.Run("missing addr port", func(t *testing.T) {
			mockDogClient := new(mock.DogAPI)
			mockCatClient := new(mock.CatAPI)
			logger := logrus.New()
			cnfg := &server.Config{
				DogAPIClient: mockDogClient,
				CatAPIClient: mockCatClient,
				Logger:       logger,
			}

			_, err := server.New(cnfg)

			assert.NotNil(t, err, "error is nil")
			assert.Equal(t, err.Error(), "invalid config: missing address port")
		})

		t.Run("missing logger", func(t *testing.T) {
			mockDogClient := new(mock.DogAPI)
			mockCatClient := new(mock.CatAPI)
			cnfg := &server.Config{
				DogAPIClient: mockDogClient,
				CatAPIClient: mockCatClient,
				Addr:         ":8000",
			}

			_, err := server.New(cnfg)

			assert.NotNil(t, err, "error is nil")
			assert.Equal(t, err.Error(), "invalid config: logger missing")
		})
	})
}
