package metal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TuningItem represents a single tuning item in the GetTuningsResponse.
type TuningItem struct {
	ID     string `json:"id"`
	App    string `json:"app"`
	IDA    string `json:"idA"`
	IDB    string `json:"idB"`
	Result string `json:"result"`
}

// GetTuningsResponse represents the response from the /apps/{appId}/tunings GET request.
type GetTuningsResponse struct {
	Data []TuningItem `json:"data"`
}

// GetTunings - This endpoint gets a list of tunings for an app.
func (c *Client) GetTunings(appID string) (*GetTuningsResponse, error) {
	url := fmt.Sprintf("%s/apps/%s/tunings", c.baseURL, appID)
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

	var getTuningsResponse GetTuningsResponse
	err = json.Unmarshal(body, &getTuningsResponse)
	if err != nil {
		return nil, err
	}

	return &getTuningsResponse, nil
}
