package metal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// IndexRequest represents the data required for the /index POST request.
type IndexRequest struct {
	App      string                 `json:"app"`
	Text     string                 `json:"text,omitempty"`     // optional
	ImageURL string                 `json:"imageUrl,omitempty"` // optional
	ID       string                 `json:"id,omitempty"`       // optional
	Metadata map[string]interface{} `json:"metadata,omitempty"` // optional
}

// IndexResponseData represents the data field in the IndexResponse.
type IndexResponseData struct {
	CreatedAt time.Time              `json:"createdAt"`
	ID        string                 `json:"id"`
	Metadata  map[string]interface{} `json:"metadata"`
	Text      string                 `json:"text"`
}

// IndexResponse represents the response from the /index POST request.
type IndexResponse struct {
	Data IndexResponseData `json:"data"`
}

// Index - This endpoint generates and stores a Document(embedding) with the inputted data.
func (c *Client) Index(req IndexRequest) (*IndexResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/index", c.baseURL)
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

	var indexResponse IndexResponse
	err = json.Unmarshal(body, &indexResponse)
	if err != nil {
		return nil, err
	}

	return &indexResponse, nil
}
