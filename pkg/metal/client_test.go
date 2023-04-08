package metal

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("returns an error if apiKey or clientID is empty", func(t *testing.T) {
		_, err := NewClient("", "")
		if err == nil {
			t.Error("Expected an error when apiKey and clientID are empty")
		}

		_, err = NewClient("api_key", "")
		if err == nil {
			t.Error("Expected an error when clientID is empty")
		}

		_, err = NewClient("", "client_id")
		if err == nil {
			t.Error("Expected an error when apiKey is empty")
		}
	})

	t.Run("creates a new client with the given apiKey and clientID", func(t *testing.T) {
		apiKey := "api_key"
		clientID := "client_id"
		client, err := NewClient(apiKey, clientID)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if client.apiKey != apiKey {
			t.Errorf("Expected apiKey to be %s, but got %s", apiKey, client.apiKey)
		}

		if client.clientID != clientID {
			t.Errorf("Expected clientID to be %s, but got %s", clientID, client.clientID)
		}

		if client.baseURL != "https://api.getmetal.io/v1" {
			t.Errorf("Expected baseURL to be https://api.getmetal.io/v1, but got %s", client.baseURL)
		}
	})
}
