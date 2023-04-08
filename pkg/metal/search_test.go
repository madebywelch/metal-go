package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, but got %s", r.Method)
		}

		if r.URL.Path != "/search" {
			t.Errorf("Expected path to be /search, but got %s", r.URL.Path)
		}

		var req SearchRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			t.Error("Failed to decode request body")
		}

		if req.App != "test_app" {
			t.Errorf("Expected app to be 'test_app', but got %s", req.App)
		}

		resp := SearchResponse{
			Data: []SearchResult{
				{
					ID:       "test_document_id",
					Dist:     "0.123",
					Metadata: map[string]interface{}{"key": "value"},
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

	t.Run("searches for documents", func(t *testing.T) {
		req := SearchRequest{
			App:  "test_app",
			Text: "test_text",
		}

		resp, err := client.Search(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if len(resp.Data) != 1 {
			t.Errorf("Expected 1 search result, but got %d", len(resp.Data))
		}

		result := resp.Data[0]
		if result.ID != "test_document_id" {
			t.Errorf("Expected document ID to be 'test_document_id', but got %s", result.ID)
		}

		if result.Dist != "0.123" {
			t.Errorf("Expected distance to be '0.123', but got %s", result.Dist)
		}
	})
}
