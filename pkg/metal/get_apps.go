package metal

import (
	"encoding/json"
	"errors"
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

	var getAppsResponse GetAppsResponse
	err = json.Unmarshal(body, &getAppsResponse)
	if err != nil {
		return nil, err
	}

	return &getAppsResponse, nil
}
