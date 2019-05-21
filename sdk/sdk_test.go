package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_Search(t *testing.T) {
	fake := Operation{}

	okServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "something went wrong", http.StatusOK)
		}),
	)
	defer okServer.Close()

	jsonServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := json.NewEncoder(w).Encode(&fake); err != nil {
				t.Fail()
			}
		}),
	)
	defer jsonServer.Close()

	t.Run("server is not running", func(t *testing.T) {
		api := New("http://invalid")
		_, err := api.Search("AX1234BT")

		if err == nil {
			t.Errorf("should throw an error")
		}
	})

	t.Run("invalid response body", func(t *testing.T) {
		_, err := New(okServer.URL).Search("AX1234BT")
		if err.Error() != "invalid response body" {
			t.Fail()
		}
	})
}
