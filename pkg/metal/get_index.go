package metal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetIndexRequest represents the data required for the /indexes/{indexId} GET request.
type GetIndexRequest struct {
	IndexID string
}

// GetIndexData represents the data field in the GetIndexResponse.
type GetIndexData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	App        string `json:"app"`
	Model      string `json:"model"`
	Dimensions int    `json:"dimensions"`
}

// GetIndexResponse represents the response from the /indexes/{indexId} GET request.
type GetIndexResponse struct {
	Data GetIndexData `json:"data"`
}

// GetIndex - This endpoint gets a single index.
func (c *Client) GetIndex(req GetIndexRequest) (*GetIndexResponse, error) {
	url := fmt.Sprintf("%s/indexes/%s", c.baseURL, req.IndexID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
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

	var getIndexResponse GetIndexResponse
	err = json.Unmarshal(body, &getIndexResponse)
	if err != nil {
		return nil, err
	}

	return &getIndexResponse, nil
}
