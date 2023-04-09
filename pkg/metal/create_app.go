package metal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateAppRequest represents the data required for the /apps POST request.
type CreateAppRequest struct {
	Name string `json:"name"`
}

// CreateAppData represents the data field in the CreateAppResponse.
type CreateAppData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateAppResponse represents the response from the /apps POST request.
type CreateAppResponse struct {
	Data CreateAppData `json:"data"`
}

// CreateApp - This endpoint creates an app.
func (c *Client) CreateApp(req CreateAppRequest) (*CreateAppResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/apps", c.baseURL)
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

	var createAppResponse CreateAppResponse
	err = json.Unmarshal(body, &createAppResponse)
	if err != nil {
		return nil, err
	}

	return &createAppResponse, nil
}
