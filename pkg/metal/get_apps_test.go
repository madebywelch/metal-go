package metal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApps(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, but got %s", r.Method)
		}

		if r.URL.Path != "/apps" {
			t.Errorf("Expected path to be /apps, but got %s", r.URL.Path)
		}

		resp := GetAppsResponse{
			Data: []GetAppsDataItem{
				{
					ID:   "app1",
					Name: "App 1",
				},
				{
					ID:   "app2",
					Name: "App 2",
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

	t.Run("gets a list of apps", func(t *testing.T) {
		resp, err := client.GetApps()

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if resp == nil {
			t.Fatalf("Expected response, but got nil")
		}

		if len(resp.Data) != 2 {
			t.Errorf("Expected 2 apps, but got %d", len(resp.Data))
		}

		if resp.Data[0].ID != "app1" || resp.Data[0].Name != "App 1" {
			t.Errorf("Expected first app to have ID app1 and name App 1, but got ID %s and name %s", resp.Data[0].ID, resp.Data[0].Name)
		}

		if resp.Data[1].ID != "app2" || resp.Data[1].Name != "App 2" {
			t.Errorf("Expected second app to have ID app2 and name App 2, but got ID %s and name %s", resp.Data[1].ID, resp.Data[1].Name)
		}
	})
}
