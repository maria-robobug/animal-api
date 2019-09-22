package client_test

import (
	"net/http"
	"testing"

	"github.com/maria-robobug/animal-api/internal/client"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, err := client.New("", "", &http.Client{})

		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		invalidClientError := "invalid client: nil client provided"
		_, err := client.New("", "", nil)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), invalidClientError)
	})
}
