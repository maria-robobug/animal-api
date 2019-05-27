package server_test

import (
	"log"
	"os"
	"testing"

	"github.com/maria-robobug/animal-api/server"

	"github.com/maria-robobug/animal-api/internal/mock"

	"github.com/stretchr/testify/assert"
)

var (
	infoLog    = log.New(os.Stdin, "", 0)
	errorLog   = log.New(os.Stderr, "", 0)
	mockClient = new(mock.DogAPI)
	mockCache  = new(mock.Cache)
)

func TestNewServer_ValidConfig(t *testing.T) {
	// given
	cnfg := &server.Config{
		Cache:        mockCache,
		DogAPIClient: mockClient,
		Addr:         ":9000",
		InfoLog:      infoLog,
		ErrorLog:     errorLog,
	}

	// when
	serv, err := server.New(cnfg)

	// then
	assert.Nil(t, err, "error is not nil")
	assert.NotNil(t, serv.DogAPIClient, "client is nil")
	assert.NotNil(t, serv.InfoLog, "info log is nil")
	assert.NotNil(t, serv.ErrorLog, "error log is nil")
}
func TestNewServer_MissingClient(t *testing.T) {
	// given
	cnfg := &server.Config{
		Cache:    mockCache,
		Addr:     ":9000",
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	// when
	_, err := server.New(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: nil client")
}

func TestNewServer_MissingCache(t *testing.T) {
	// given
	cnfg := &server.Config{
		DogAPIClient: mockClient,
		Addr:         ":9000",
		InfoLog:      infoLog,
		ErrorLog:     errorLog,
	}

	// when
	_, err := server.New(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: nil cache")
}

func TestNewServer_MissingAddrPort(t *testing.T) {
	// given
	cnfg := &server.Config{
		Cache:        mockCache,
		DogAPIClient: mockClient,
		InfoLog:      infoLog,
		ErrorLog:     errorLog,
	}

	// when
	_, err := server.New(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: missing address port")
}

func TestNewServer_MissingLogger(t *testing.T) {
	// given
	cnfg := &server.Config{
		Cache:        mockCache,
		DogAPIClient: mockClient,
		Addr:         ":8000",
	}

	// when
	_, err := server.New(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: logger missing")
}
