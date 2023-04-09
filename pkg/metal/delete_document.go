package metal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeleteDocumentRequest represents the data required for the /documents/{documentId} DELETE request.
type DeleteDocumentRequest struct {
	DocumentID string
}

// DeleteDocumentResponse represents the response from the /documents/{documentId} DELETE request.
type DeleteDocumentResponse struct {
	Data struct{} `json:"data"`
}

// DeleteDocument - This endpoint deletes an embedding document.
func (c *Client) DeleteDocument(req DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	url := fmt.Sprintf("%s/documents/%s", c.baseURL, req.DocumentID)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
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

	var deleteDocumentResponse DeleteDocumentResponse
	err = json.Unmarshal(body, &deleteDocumentResponse)
	if err != nil {
		return nil, err
	}

	return &deleteDocumentResponse, nil
}
