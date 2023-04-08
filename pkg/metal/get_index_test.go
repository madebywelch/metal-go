package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIndex(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, but got %s", r.Method)
		}

		if r.URL.Path != "/indexes/test_index_id" {
			t.Errorf("Expected path to be /indexes/test_index_id, but got %s", r.URL.Path)
		}

		resp := GetIndexResponse{
			Data: GetIndexData{
				ID:         "test_index_id",
				Name:       "Test index",
				App:        "test_app",
				Model:      "test_model",
				Dimensions: 128,
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

	t.Run("gets an index", func(t *testing.T) {
		req := GetIndexRequest{
			IndexID: "test_index_id",
		}

		resp, err := client.GetIndex(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if resp.Data.ID != "test_index_id" {
			t.Errorf("Expected index ID to be test_index_id, but got %s", resp.Data.ID)
		}

		if resp.Data.Name != "Test index" {
			t.Errorf("Expected index name to be 'Test index', but got '%s'", resp.Data.Name)
		}

		if resp.Data.App != "test_app" {
			t.Errorf("Expected index app to be 'test_app', but got '%s'", resp.Data.App)
		}

		if resp.Data.Model != "test_model" {
			t.Errorf("Expected index model to be 'test_model', but got '%s'", resp.Data.Model)
		}

		if resp.Data.Dimensions != 128 {
			t.Errorf("Expected index dimensions to be 128, but got %d", resp.Data.Dimensions)
		}
	})
}
