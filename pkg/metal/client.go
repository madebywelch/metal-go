package metal

import (
	"errors"
	"net/http"
)

// Client represents the Metal API client and its configuration.
type Client struct {
	httpClient *http.Client
	apiKey     string
	clientID   string
	baseURL    string
}

// NewClient initializes a new Metal API client with the required headers.
func NewClient(apiKey, clientID string) (*Client, error) {
	if apiKey == "" || clientID == "" {
		return nil, errors.New("apiKey and clientID are required")
	}

	client := &Client{
		httpClient: &http.Client{},
		apiKey:     apiKey,
		clientID:   clientID,
		baseURL:    "https://api.getmetal.io/v1",
	}

	return client, nil
}
