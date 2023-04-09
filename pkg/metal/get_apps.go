package metal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAppsDataItem represents a single app item in the GetAppsResponse.
type GetAppsDataItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetAppsResponse represents the response from the /apps GET request.
type GetAppsResponse struct {
	Data []GetAppsDataItem `json:"data"`
}

// GetApps - This endpoint gets a list of apps.
func (c *Client) GetApps() (*GetAppsResponse, error) {
	url := fmt.Sprintf("%s/apps", c.baseURL)
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

	var getAppsResponse GetAppsResponse
	err = json.Unmarshal(body, &getAppsResponse)
	if err != nil {
		return nil, err
	}

	return &getAppsResponse, nil
}
