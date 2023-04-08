package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, but got %s", r.Method)
		}

		if r.URL.Path != "/apps/test_app_id/indexes" {
			t.Errorf("Expected path to be /apps/test_app_id/indexes, but got %s", r.URL.Path)
		}

		var req CreateIndexRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			t.Error(err)
		}

		resp := CreateIndexResponse{
			Data: IndexData{
				ID:         "1",
				Name:       req.Name,
				App:        req.AppID,
				Model:      req.Model,
				Dimensions: req.Dimensions,
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

	t.Run("creates an index with the given properties", func(t *testing.T) {
		req := CreateIndexRequest{
			AppID:      "test_app_id",
			Model:      "custom",
			Name:       "Test Index",
			Dimensions: 512,
		}

		resp, err := client.CreateIndex(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp.Data.App != req.AppID {
			t.Errorf("Expected AppID to be %s, but got %s", req.AppID, resp.Data.App)
		}

		if resp.Data.Model != req.Model {
			t.Errorf("Expected Model to be %s, but got %s", req.Model, resp.Data.Model)
		}

		if resp.Data.Name != req.Name {
			t.Errorf("Expected Name to be %s, but got %s", req.Name, resp.Data.Name)
		}

		if resp.Data.Dimensions != req.Dimensions {
			t.Errorf("Expected Dimensions to be %d, but got %d", req.Dimensions, resp.Data.Dimensions)
		}

		if resp.Data.ID == "" {
			t.Error("Expected index ID to be non-empty")
		}
	})
}
