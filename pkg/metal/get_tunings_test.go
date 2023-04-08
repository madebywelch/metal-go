package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTunings(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, but got %s", r.Method)
		}

		if r.URL.Path != "/apps/test_app_id/tunings" {
			t.Errorf("Expected path to be /apps/test_app_id/tunings, but got %s", r.URL.Path)
		}

		resp := GetTuningsResponse{
			Data: []TuningItem{
				{
					ID:     "test_tuning_id1",
					App:    "test_app_id",
					IDA:    "test_idA1",
					IDB:    "test_idB1",
					Result: "test_result1",
				},
				{
					ID:     "test_tuning_id2",
					App:    "test_app_id",
					IDA:    "test_idA2",
					IDB:    "test_idB2",
					Result: "test_result2",
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

	t.Run("gets tunings", func(t *testing.T) {
		resp, err := client.GetTunings("test_app_id")

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if len(resp.Data) != 2 {
			t.Errorf("Expected 2 tuning items, but got %d", len(resp.Data))
		}

		if resp.Data[0].ID != "test_tuning_id1" {
			t.Errorf("Expected first tuning ID to be test_tuning_id1, but got %s", resp.Data[0].ID)
		}

		if resp.Data[1].ID != "test_tuning_id2" {
			t.Errorf("Expected second tuning ID to be test_tuning_id2, but got %s", resp.Data[1].ID)
		}
	})
}
