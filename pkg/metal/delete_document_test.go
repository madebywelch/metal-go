package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteDocument(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected method DELETE, but got %s", r.Method)
		}

		if r.URL.Path != "/documents/test_document_id" {
			t.Errorf("Expected path to be /documents/test_document_id, but got %s", r.URL.Path)
		}

		resp := DeleteDocumentResponse{
			Data: struct{}{},
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

	t.Run("deletes a document by given ID", func(t *testing.T) {
		req := DeleteDocumentRequest{
			DocumentID: "test_document_id",
		}

		resp, err := client.DeleteDocument(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}
	})
}
