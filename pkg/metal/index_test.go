package metal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, but got %s", r.Method)
		}

		if r.URL.Path != "/index" {
			t.Errorf("Expected path to be /index, but got %s", r.URL.Path)
		}

		var req IndexRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			t.Error("Failed to decode request body")
		}

		if req.App != "test_app" {
			t.Errorf("Expected app to be 'test_app', but got %s", req.App)
		}

		resp := IndexResponse{
			CreatedAt: time.Now(),
			ID:        "test_document_id",
			Metadata:  req.Metadata,
			Text:      req.Text,
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

	t.Run("indexes a document", func(t *testing.T) {
		req := IndexRequest{
			App:  "test_app",
			Text: "test_text",
			ID:   "test_id",
			Metadata: map[string]interface{}{
				"key": "value",
			},
		}

		resp, err := client.Index(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if resp.ID != "test_document_id" {
			t.Errorf("Expected document ID to be 'test_document_id', but got %s", resp.ID)
		}

		if !bytes.Equal([]byte(resp.Text), []byte(req.Text)) {
			t.Errorf("Expected document text to be '%s', but got '%s'", req.Text, resp.Text)
		}
	})
}
