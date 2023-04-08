package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, but got %s", r.Method)
		}

		if r.URL.Path != "/indexes/test_index_id" {
			t.Errorf("Expected path to be /indexes/test_index_id, but got %s", r.URL.Path)
		}

		resp := GetAppResponse{
			Data: GetAppData{
				ID:   "test_index_id",
				Name: "Test App",
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

	t.Run("gets an app by given index ID", func(t *testing.T) {
		req := GetAppRequest{
			IndexID: "test_index_id",
		}

		resp, err := client.GetApp(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if resp.Data.ID != "test_index_id" {
			t.Errorf("Expected app ID to be test_index_id, but got %s", resp.Data.ID)
		}

		if resp.Data.Name != "Test App" {
			t.Errorf("Expected app name to be Test App, but got %s", resp.Data.Name)
		}
	})
}
