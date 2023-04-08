package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, but got %s", r.Method)
		}

		if r.URL.Path != "/apps" {
			t.Errorf("Expected path to be /apps, but got %s", r.URL.Path)
		}

		var req CreateAppRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			t.Error(err)
		}

		if req.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		resp := CreateAppResponse{
			Data: CreateAppData{
				ID:   "1",
				Name: req.Name,
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

	t.Run("returns an error if app name is empty", func(t *testing.T) {
		_, err := client.CreateApp(CreateAppRequest{Name: ""})
		if err == nil {
			t.Error("Expected an error when app name is empty")
		}
	})

	t.Run("creates an app with the given name", func(t *testing.T) {
		appName := "Test App"
		resp, err := client.CreateApp(CreateAppRequest{Name: appName})

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp.Data.Name != appName {
			t.Errorf("Expected app name to be %s, but got %s", appName, resp.Data.Name)
		}

		if resp.Data.ID == "" {
			t.Error("Expected app ID to be non-empty")
		}
	})
}
