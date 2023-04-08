package metal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetDocumentRequest represents the data required for the /documents/{documentId} GET request.
type GetDocumentRequest struct {
	DocumentID string
}

// GetDocumentData represents the data field in the GetDocumentResponse.
type GetDocumentData struct {
	ID        string                 `json:"id"`
	Text      string                 `json:"text"`
	CreatedAt time.Time              `json:"createdAt"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// GetDocumentResponse represents the response from the /documents/{documentId} GET request.
type GetDocumentResponse struct {
	Data GetDocumentData `json:"data"`
}

// GetDocument - This endpoint gets an embedding document.
func (c *Client) GetDocument(req GetDocumentRequest) (*GetDocumentResponse, error) {
	url := fmt.Sprintf("%s/documents/%s", c.baseURL, req.DocumentID)
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

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var getDocumentResponse GetDocumentResponse
	err = json.Unmarshal(body, &getDocumentResponse)
	if err != nil {
		return nil, err
	}

	return &getDocumentResponse, nil
}
