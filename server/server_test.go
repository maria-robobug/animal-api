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
			mockClient := new(mock.DogAPI)
			logger := logrus.New()
			cnfg := &server.Config{
				DogAPIClient: mockClient,
				Addr:         ":9000",
				Logger:       logger,
			}

			serv, err := server.New(cnfg)

			assert.Nil(t, err, "error is not nil")
			assert.NotNil(t, serv.DogAPIClient, "client is nil")
			assert.NotNil(t, serv.Logger, "logger is nil")
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("missing client", func(t *testing.T) {
			logger := logrus.New()
			cnfg := &server.Config{
				Addr:   ":9000",
				Logger: logger,
			}

			_, err := server.New(cnfg)

			assert.NotNil(t, err, "error is nil")
			assert.Equal(t, err.Error(), "invalid config: nil client")
		})

		t.Run("missing addr port", func(t *testing.T) {
			mockClient := new(mock.DogAPI)
			logger := logrus.New()
			cnfg := &server.Config{
				DogAPIClient: mockClient,
				Logger:       logger,
			}

			_, err := server.New(cnfg)

			assert.NotNil(t, err, "error is nil")
			assert.Equal(t, err.Error(), "invalid config: missing address port")
		})

		t.Run("missing logger", func(t *testing.T) {
			mockClient := new(mock.DogAPI)
			cnfg := &server.Config{DogAPIClient: mockClient, Addr: ":8000"}

			_, err := server.New(cnfg)

			assert.NotNil(t, err, "error is nil")
			assert.Equal(t, err.Error(), "invalid config: logger missing")
		})
	})
}
