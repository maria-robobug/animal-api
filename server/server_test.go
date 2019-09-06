package server_test

import (
	"testing"

	"github.com/maria-robobug/animal-api/server"
	"github.com/sirupsen/logrus"

	"github.com/maria-robobug/animal-api/internal/mock"

	"github.com/stretchr/testify/assert"
)

func TestNewServer_ValidConfig(t *testing.T) {
	mockClient := new(mock.DogAPI)
	logger := logrus.New()

	// given
	cnfg := &server.Config{
		DogAPIClient: mockClient,
		Addr:         ":9000",
		Logger:       logger,
	}

	// when
	serv, err := server.New(cnfg)

	// then
	assert.Nil(t, err, "error is not nil")
	assert.NotNil(t, serv.DogAPIClient, "client is nil")
	assert.NotNil(t, serv.Logger, "logger is nil")
}
func TestNewServer_MissingClient(t *testing.T) {
	logger := logrus.New()

	// given
	cnfg := &server.Config{
		Addr:   ":9000",
		Logger: logger,
	}

	// when
	_, err := server.New(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: nil client")
}

func TestNewServer_MissingAddrPort(t *testing.T) {
	mockClient := new(mock.DogAPI)
	logger := logrus.New()

	// given
	cnfg := &server.Config{
		DogAPIClient: mockClient,
		Logger:       logger,
	}

	// when
	_, err := server.New(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: missing address port")
}

func TestNewServer_MissingLogger(t *testing.T) {
	mockClient := new(mock.DogAPI)

	// given
	cnfg := &server.Config{DogAPIClient: mockClient, Addr: ":8000"}

	// when
	_, err := server.New(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: logger missing")
}
