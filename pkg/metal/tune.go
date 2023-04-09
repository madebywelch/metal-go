package metal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TuneRequest represents the data required for the /tune POST request.
type TuneRequest struct {
	App   string `json:"app"`
	IDA   string `json:"idA"`
	IDB   string `json:"idB"`
	Label int    `json:"label"`
}

// TuneData represents the data field in the TuneResponse.
type TuneData struct {
	ID    string `json:"id"`
	App   string `json:"app"`
	IDA   string `json:"idA"`
	IDB   string `json:"idB"`
	Label int    `json:"label"`
}

// TuneResponse represents the response from the /tune POST request.
type TuneResponse struct {
	Data TuneData `json:"data"`
}

// Tune - Tune your embeddings for better results.
func (c *Client) Tune(req TuneRequest) (*TuneResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/tune", c.baseURL)
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

	var tuneResponse TuneResponse
	err = json.Unmarshal(body, &tuneResponse)
	if err != nil {
		return nil, err
	}

	return &tuneResponse, nil
}
