package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIndexes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, but got %s", r.Method)
		}

		if r.URL.Path != "/apps/test_app_id/indexes" {
			t.Errorf("Expected path to be /apps/test_app_id/indexes, but got %s", r.URL.Path)
		}

		resp := GetIndexesResponse{
			Data: []IndexData{
				{
					ID:         "index1",
					Name:       "Test index 1",
					App:        "test_app_id",
					Model:      "test_model",
					Dimensions: 128,
				},
				{
					ID:         "index2",
					Name:       "Test index 2",
					App:        "test_app_id",
					Model:      "test_model",
					Dimensions: 128,
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

	t.Run("gets indexes", func(t *testing.T) {
		req := GetIndexesRequest{
			AppID: "test_app_id",
		}

		resp, err := client.GetIndexes(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if len(resp.Data) != 2 {
			t.Errorf("Expected 2 indexes, but got %d", len(resp.Data))
		}

		if resp.Data[0].ID != "index1" {
			t.Errorf("Expected index ID to be index1, but got %s", resp.Data[0].ID)
		}

		if resp.Data[0].Name != "Test index 1" {
			t.Errorf("Expected index name to be 'Test index 1', but got '%s'", resp.Data[0].Name)
		}

		if resp.Data[1].ID != "index2" {
			t.Errorf("Expected index ID to be index2, but got %s", resp.Data[1].ID)
		}

		if resp.Data[1].Name != "Test index 2" {
			t.Errorf("Expected index name to be 'Test index 2', but got '%s'", resp.Data[1].Name)
		}
	})
}
