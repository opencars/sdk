package toolkit

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	operationsFixture    = "./fixtures/operations.json"
	registrationsFixture = "./fixtures/registrations.json"
)

func TestAPI_Operation(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "something went wrong", http.StatusOK)
		}),
	)
	defer server.Close()

	t.Run("server is not running", func(t *testing.T) {
		api := NewSDK("http://invalid")
		_, err := api.Operation("AX1234BT")

		if err == nil {
			t.Errorf("should throw an error")
		}
	})

	t.Run("invalid response body", func(t *testing.T) {
		_, err := NewSDK(server.URL).Operation("AX1234BT")
		if err.Error() != "invalid response body" {
			t.Fail()
		}
	})
}

func TestAPI_Operations(t *testing.T) {
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
		api := NewSDK("http://invalid")
		_, err := api.Operations("AX1234BT", 1)

		if err == nil {
			t.Errorf("should throw an error")
		}
	})

	t.Run("invalid response body", func(t *testing.T) {
		_, err := NewSDK(okServer.URL).Operations("AX1234BT", 1)
		if err.Error() != "invalid response body" {
			t.Fail()
		}
	})
}

func TestAPI_OperationsWithFixture(t *testing.T) {
	fake, err := os.Open(operationsFixture)
	if err != nil {
		panic(err)
	}
	defer fake.Close()

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.Copy(w, fake)
		}),
	)
	defer server.Close()

	operations, err := NewSDK(server.URL).Operations("АА9359РС", 1)
	if err != nil {
		t.Fail()
	}

	if len(operations) != 1 {
		t.Fail()
	}

	if operations[0].Number != "АА9359РС" {
		t.Errorf("%v != %v", operations[0].Number, "АА9359РС")
	}
}

func TestAPI_Registrations(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "something went wrong", http.StatusOK)
		}),
	)
	defer server.Close()

	t.Run("server is not running", func(t *testing.T) {
		api := NewSDK("http://invalid")
		_, err := api.Registrations("CXI012345")

		if err == nil {
			t.Errorf("should throw an error")
		}
	})

	t.Run("invalid response body", func(t *testing.T) {
		_, err := NewSDK(server.URL).Registrations("CXI012345")
		if err.Error() != "invalid response body" {
			t.Fail()
		}
	})
}

func TestAPI_RegistrationsWithFixture(t *testing.T) {
	fake, err := os.Open(registrationsFixture)
	if err != nil {
		panic(err)
	}
	defer fake.Close()

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.Copy(w, fake)
		}),
	)
	defer server.Close()

	registrations, err := NewSDK(server.URL).Registrations("СХН484154")
	if err != nil {
		t.Fail()
	}

	if len(registrations) != 1 {
		t.Fail()
	}

	if registrations[0].Code != "СХН484154" {
		t.Errorf("%v != %v", registrations[0].Code, "СХН484154")
	}
}
