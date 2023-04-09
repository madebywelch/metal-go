package metal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetIndexesRequest represents the data required for the /apps/{appId}/indexes GET request.
type GetIndexesRequest struct {
	AppID string
}

// IndexData represents the data of a single index in the GetIndexesResponse.
type IndexData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	App        string `json:"app"`
	Model      string `json:"model"`
	Dimensions int    `json:"dimensions"`
}

// GetIndexesResponse represents the response from the /apps/{appId}/indexes GET request.
type GetIndexesResponse struct {
	Data []IndexData `json:"data"`
}

// GetIndexes - This endpoint gets a list of indexes for an app.
func (c *Client) GetIndexes(req GetIndexesRequest) (*GetIndexesResponse, error) {
	url := fmt.Sprintf("%s/apps/%s/indexes", c.baseURL, req.AppID)
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

	var getIndexesResponse GetIndexesResponse
	err = json.Unmarshal(body, &getIndexesResponse)
	if err != nil {
		return nil, err
	}

	return &getIndexesResponse, nil
}
