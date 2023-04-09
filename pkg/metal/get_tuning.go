package metal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetTuningResponse represents the response from the /tunings/{tuningId} GET request.
type GetTuningResponse struct {
	Data struct {
		ID     string `json:"id"`
		App    string `json:"app"`
		IDA    string `json:"idA"`
		IDB    string `json:"idB"`
		Result string `json:"result"`
	} `json:"data"`
}

// GetTuning - This endpoint gets a single Tuning.
func (c *Client) GetTuning(tuningID string) (*GetTuningResponse, error) {
	url := fmt.Sprintf("%s/tunings/%s", c.baseURL, tuningID)
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

	var getTuningResponse GetTuningResponse
	err = json.Unmarshal(body, &getTuningResponse)
	if err != nil {
		return nil, err
	}

	return &getTuningResponse, nil
}
