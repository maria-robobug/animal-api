package rest_test

import (
	"log"
	"os"
	"testing"

	"github.com/maria-robobug/animal-api/pkg/dog/rest"

	"github.com/maria-robobug/animal-api/pkg/mock"

	"github.com/stretchr/testify/assert"
)

func TestNewServer_ValidConfig(t *testing.T) {
	infoLog := log.New(os.Stdin, "", 0)
	errorLog := log.New(os.Stderr, "", 0)
	mockClient := new(mock.DogAPI)

	// given
	cnfg := &rest.Config{
		Client:   mockClient,
		Addr:     ":9000",
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	// when
	serv, err := rest.NewServer(cnfg)

	// then
	assert.Nil(t, err, "error is not nil")
	assert.NotNil(t, serv.Client, "client is nil")
	assert.NotNil(t, serv.InfoLog, "info log is nil")
	assert.NotNil(t, serv.ErrorLog, "error log is nil")
}
func TestNewServer_MissingClient(t *testing.T) {
	// given
	cnfg := &rest.Config{}

	// when
	_, err := rest.NewServer(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: nil client")
}

func TestNewServer_MissingAddrPort(t *testing.T) {
	mockClient := new(mock.DogAPI)

	// given
	cnfg := &rest.Config{Client: mockClient}

	// when
	_, err := rest.NewServer(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: missing address port")
}

func TestNewServer_MissingLogger(t *testing.T) {
	mockClient := new(mock.DogAPI)

	// given
	cnfg := &rest.Config{Client: mockClient, Addr: ":8000"}

	// when
	_, err := rest.NewServer(cnfg)

	// then
	assert.NotNil(t, err, "error is nil")
	assert.Equal(t, err.Error(), "invalid config: logger missing")
}
