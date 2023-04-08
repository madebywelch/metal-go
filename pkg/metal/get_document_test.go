package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetDocument(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, but got %s", r.Method)
		}

		if r.URL.Path != "/documents/test_document_id" {
			t.Errorf("Expected path to be /documents/test_document_id, but got %s", r.URL.Path)
		}

		resp := GetDocumentResponse{
			Data: GetDocumentData{
				ID:        "test_document_id",
				Text:      "Test document",
				CreatedAt: time.Now(),
				Metadata: map[string]interface{}{
					"example": "metadata",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	client, err := NewClient("test_api_key", "test_client_id")
	if err != nil {
		t.Fatal(err)
	}
	client.baseURL = ts.URL

	t.Run("gets a document", func(t *testing.T) {
		req := GetDocumentRequest{
			DocumentID: "test_document_id",
		}

		resp, err := client.GetDocument(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if resp.Data.ID != "test_document_id" {
			t.Errorf("Expected document ID to be test_document_id, but got %s", resp.Data.ID)
		}

		if resp.Data.Text != "Test document" {
			t.Errorf("Expected document text to be 'Test document', but got '%s'", resp.Data.Text)
		}

		if _, ok := resp.Data.Metadata["example"]; !ok {
			t.Errorf("Expected document metadata to contain 'example' key, but it was not present")
		}
	})
}
