package sdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const (
	operationsPath    = "/vehicle/operations"
	registrationsPath = "/vehicle/registrations"
)

// Client is simple representation of opencars SDK.
type Client struct {
	uri string
}

// New initializes new instance of client structure.
func New(uri string) *Client {
	client := new(Client)

	client.uri = uri

	return client
}

func (client *Client) searchOperations(number string, limit int) ([]Operation, error) {
	if number == "" {
		return nil, errors.New("number is empty")
	}

	query := client.uri + operationsPath + "?number=" + number + "&limit=" + strconv.Itoa(limit)
	response, err := http.Get(query)

	if err != nil {
		return nil, err
	}

	models := make([]Operation, 0)
	if err := json.NewDecoder(response.Body).Decode(&models); err != nil {
		return nil, errors.New("invalid response body")
	}

	return models, nil
}

func (client *Client) searchRegistrations(code string) ([]Registration, error) {
	if code == "" {
		return nil, errors.New("code is empty")
	}

	query := client.uri + registrationsPath + "?code=" + code
	response, err := http.Get(query)

	if err != nil {
		return nil, err
	}

	models := make([]Registration, 0)
	if err := json.NewDecoder(response.Body).Decode(&models); err != nil {
		return nil, errors.New("invalid response body")
	}

	return models, nil
}

// Search makes API request to get transport info by government number.
// Returns first operation from OpenCars operation tables.
func (client *Client) Operation(number string) (*Operation, error) {
	operation, err := client.searchOperations(number, 1)

	if err != nil {
		return nil, err
	}

	return &operation[0], nil
}

// Search makes API request to get transport info by government number.
// Returns first operation from OpenCars operation tables.
func (client *Client) Registrations(code string) ([]Registration, error) {
	registrations, err := client.searchRegistrations(code)

	if err != nil {
		return nil, err
	}

	return registrations, nil
}

// SearchLimit makes API request to opencars with limit parameter.
func (client *Client) Operations(number string, limit int) ([]Operation, error) {
	return client.searchOperations(number, limit)
}
