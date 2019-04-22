package sdk

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Transport represents JSON object for opencars API.
type Transport struct {
	Registration string `json:"registration"`
	Date         string `json:"date"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	Color        string `json:"color"`
	Kind         string `json:"kind"`
	Body         string `json:"body"`
	Purpose      string `json:"purpose"`
	Fuel         string `json:"fuel"`
	Capacity     int    `json:"capacity"`
	Weight       int    `json:"own_weight"`
	Number       string `json:"number"`
}

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

func (client *Client) search(number string, limit int) ([]Transport, error) {
	if number == "" {
		return nil, errors.New("number is empty")
	}

	query := client.uri + "/transport/?number=" + number
	response, err := http.Get(query)

	if err != nil {
		return nil, err
	}

	models := make([]Transport, 0)
	if err := json.NewDecoder(response.Body).Decode(&models); err != nil {
		return nil, errors.New("invalid response body")
	}

	return models, nil
}

// Search makes API request to get transport info by government number.
// Returns first transport from opencars transports tables.
func (client *Client) Search(number string) (*Transport, error) {
	transport, err := client.search(number, 1)

	if err != nil {
		return nil, err
	}

	return &transport[0], nil
}

// SearchLimit makes API request to opencars with limit parameter.
func (client *Client) SearchLimit(number string, limit int) ([]Transport, error) {
	return client.search(number, limit)
}
