package metal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SearchRequest represents the data required for the /search POST request.
type SearchRequest struct {
	App  string `json:"app"`
	Text string `json:"text"`
}

// SearchResult represents a single search result item.
type SearchResult struct {
	ID       string                 `json:"id"`
	Dist     string                 `json:"dist"`
	Metadata map[string]interface{} `json:"metadata"`
}

// SearchResponse represents the response from the /search POST request.
type SearchResponse struct {
	Data []SearchResult `json:"data"`
}

// Search - This endpoint searches for the closest Documents(embeddings)
func (c *Client) Search(req SearchRequest) (*SearchResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/search", c.baseURL)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-metal-api-key", c.apiKey)
	request.Header.Set("x-metal-client-id", c.clientID)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var searchResponse SearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	return &searchResponse, nil
}