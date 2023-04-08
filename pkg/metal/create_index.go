package metal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateIndexRequest represents the data required for the /apps/{appId}/indexes POST request.
type CreateIndexRequest struct {
	AppID      string `json:"app"`
	Name       string `json:"name"`
	Model      string `json:"model"`
	Dimensions int    `json:"dimensions,omitempty"` // optional
}

// CreateIndexResponse represents the response from the /apps/{appId}/indexes POST request.
type CreateIndexResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AppID      string `json:"app"`
	Model      string `json:"model"`
	Dimensions int    `json:"dimensions"`
}

// CreateIndex - This endpoint creates an index for an app.
func (c *Client) CreateIndex(req CreateIndexRequest) (*CreateIndexResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/apps/%s/indexes", c.baseURL, req.AppID)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-metal-api-key", c.apiKey)
	request.Header.Set("x-metal-client-id", c.clientID)

	response, err := c.doRequest(request)
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

	var createIndexResponse CreateIndexResponse
	err = json.Unmarshal(body, &createIndexResponse)
	if err != nil {
		return nil, err
	}

	return &createIndexResponse, nil
}
