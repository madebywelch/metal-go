package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTune(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, but got %s", r.Method)
		}

		if r.URL.Path != "/tune" {
			t.Errorf("Expected path to be /tune, but got %s", r.URL.Path)
		}

		var req TuneRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			t.Error("Failed to decode request body")
		}

		if req.App != "test_app" {
			t.Errorf("Expected app to be 'test_app', but got %s", req.App)
		}

		resp := TuneResponse{
			Data: TuneData{
				ID:    "test_tuning_id",
				App:   "test_app",
				IDA:   "test_idA",
				IDB:   "test_idB",
				Label: 1,
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

	t.Run("tunes embeddings", func(t *testing.T) {
		req := TuneRequest{
			App:   "test_app",
			IDA:   "test_idA",
			IDB:   "test_idB",
			Label: 1,
		}

		resp, err := client.Tune(req)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if resp.Data.ID != "test_tuning_id" {
			t.Errorf("Expected tuning ID to be 'test_tuning_id', but got %s", resp.Data.ID)
		}

		if resp.Data.App != "test_app" {
			t.Errorf("Expected app to be 'test_app', but got %s", resp.Data.App)
		}

		if resp.Data.IDA != "test_idA" {
			t.Errorf("Expected idA to be 'test_idA', but got %s", resp.Data.IDA)
		}

		if resp.Data.IDB != "test_idB" {
			t.Errorf("Expected idB to be 'test_idB', but got %s", resp.Data.IDB)
		}

		if resp.Data.Label != 1 {
			t.Errorf("Expected label to be 1, but got %d", resp.Data.Label)
		}
	})
}
