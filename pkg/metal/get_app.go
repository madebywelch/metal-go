package metal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAppRequest represents the data required for the /indexes/{indexId} GET request.
type GetAppRequest struct {
	IndexID string
}

// GetAppData represents the data field in the GetAppResponse.
type GetAppData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetAppResponse represents the response from the /indexes/{indexId} GET request.
type GetAppResponse struct {
	Data GetAppData `json:"data"`
}

// GetApp - This endpoint gets a single index.
func (c *Client) GetApp(req GetAppRequest) (*GetAppResponse, error) {
	url := fmt.Sprintf("%s/indexes/%s", c.baseURL, req.IndexID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
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

	var getAppResponse GetAppResponse
	err = json.Unmarshal(body, &getAppResponse)
	if err != nil {
		return nil, err
	}

	return &getAppResponse, nil
}
