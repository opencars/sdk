package toolkit_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/opencars/toolkit"
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
		api := toolkit.New("http://invalid", "test")
		_, err := api.Operation().FindByNumber("AX1234BT")

		if err == nil {
			t.Errorf("should throw an error")
		}
	})

	t.Run("invalid response body", func(t *testing.T) {
		_, err := toolkit.New(server.URL, "test").Operation().FindByNumber("AX1234BT")
		if err.Error() != "invalid response body" {
			t.Fail()
		}
	})
}

func TestAPI_Operations(t *testing.T) {
	fake := toolkit.Operation{}

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
		api := toolkit.New("http://invalid", "test")
		_, err := api.Operation().FindByNumber("AX1234BT")

		if err == nil {
			t.Errorf("should throw an error")
		}
	})

	t.Run("invalid response body", func(t *testing.T) {
		_, err := toolkit.New(okServer.URL, "test").Operation().FindByNumber("AX1234BT")
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
			_, _ = io.Copy(w, fake)
		}),
	)
	defer server.Close()

	operations, err := toolkit.New(server.URL, "test").Operation().FindByNumber("АА9359РС")
	if err != nil {
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
		api := toolkit.New("http://invalid", "test")
		_, err := api.Registration().FindByCode("CXH484154")

		if err == nil {
			t.Errorf("should throw an error")
		}
	})

	t.Run("invalid response body", func(t *testing.T) {
		_, err := toolkit.New(server.URL, "test").Registration().FindByCode("CXH484154")
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

	registrations, err := toolkit.New(server.URL, "test").Registration().FindByCode("CXH484154")
	if err != nil {
		t.Fail()
	}

	if len(registrations) != 1 {
		t.Fail()
	}

	if registrations[0].SDoc+registrations[0].NDoc != "CXH484154" {
		t.Errorf("%v != %v", registrations[0].SDoc+registrations[0].NDoc, "CXH484154")
	}
}
