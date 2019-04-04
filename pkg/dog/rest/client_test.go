package rest_test

import (
	"net/http"
	"testing"

	"github.com/maria-robobug/animal-api/pkg/dog/rest"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	_, err := rest.NewClient("", "", &http.Client{})

	assert.Nil(t, err)
}

func TestInvalidClient(t *testing.T) {
	invalidClientError := "invalid client: nil client provided"
	_, err := rest.NewClient("", "", nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), invalidClientError)
}
