package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTuning(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, but got %s", r.Method)
		}

		if r.URL.Path != "/tunings/test_tuning_id" {
			t.Errorf("Expected path to be /tunings/test_tuning_id, but got %s", r.URL.Path)
		}

		resp := GetTuningResponse{
			Data: struct {
				ID     string `json:"id"`
				App    string `json:"app"`
				IDA    string `json:"idA"`
				IDB    string `json:"idB"`
				Result string `json:"result"`
			}{
				ID:     "test_tuning_id",
				App:    "test_app_id",
				IDA:    "test_idA",
				IDB:    "test_idB",
				Result: "test_result",
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

	t.Run("gets tuning", func(t *testing.T) {
		resp, err := client.GetTuning("test_tuning_id")

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if resp.Data.ID != "test_tuning_id" {
			t.Errorf("Expected tuning ID to be test_tuning_id, but got %s", resp.Data.ID)
		}

		if resp.Data.App != "test_app_id" {
			t.Errorf("Expected app ID to be test_app_id, but got %s", resp.Data.App)
		}

		if resp.Data.IDA != "test_idA" {
			t.Errorf("Expected IDA to be test_idA, but got %s", resp.Data.IDA)
		}

		if resp.Data.IDB != "test_idB" {
			t.Errorf("Expected IDB to be test_idB, but got %s", resp.Data.IDB)
		}

		if resp.Data.Result != "test_result" {
			t.Errorf("Expected result to be test_result, but got %s", resp.Data.Result)
		}
	})
}
