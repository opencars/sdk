package sdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Operation represents public registrations of transport.
type Operation struct {
	Person      string `json:"person"`
	Address     string `json:"address"`
	Code        int    `json:"operation"`
	Description string `json:"description"`
	Date        string `json:"date"`
	OfficeID    int    `json:"office_id"`
	OfficeName  string `json:"office_name"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Year        int    `json:"year"`
	Color       string `json:"color"`
	Kind        string `json:"kind"`
	Body        string `json:"body"`
	Purpose     string `json:"purpose"`
	Fuel        string `json:"fuel"`
	Capacity    int    `json:"capacity"`
	Weight      int    `json:"weight"`
	Number      string `json:"number"`
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

func (client *Client) search(number string, limit int) ([]Operation, error) {
	if number == "" {
		return nil, errors.New("number is empty")
	}

	query := client.uri + "/transport?number=" + number + "&limit=" + strconv.Itoa(limit)
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

// Search makes API request to get transport info by government number.
// Returns first operation from opencars operation tables.
func (client *Client) Search(number string) (*Operation, error) {
	operation, err := client.search(number, 1)

	if err != nil {
		return nil, err
	}

	return &operation[0], nil
}

// SearchLimit makes API request to opencars with limit parameter.
func (client *Client) SearchLimit(number string, limit int) ([]Operation, error) {
	return client.search(number, limit)
}
