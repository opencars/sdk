package toolkit

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	registrationPath = "/api/v1/edrmvs/"
)

type RegistrationClient struct {
	base *Client
}

func (client *RegistrationClient) FindByVIN(vin string) ([]Registration, error) {
	query := client.base.uri + registrationPath + "?vin=" + vin

	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Api-Key", client.base.token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	registrations := make([]Registration, 0)
	if err := json.NewDecoder(response.Body).Decode(&registrations); err != nil {
		return nil, errors.New("invalid response body")
	}

	return registrations, nil
}

func (client *RegistrationClient) FindByNumber(number string) ([]Registration, error) {
	query := client.base.uri + registrationPath + "?number=" + number

	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Api-Key", client.base.token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	registrations := make([]Registration, 0)
	if err := json.NewDecoder(response.Body).Decode(&registrations); err != nil {
		return nil, errors.New("invalid response body")
	}

	return registrations, nil
}

func (client *RegistrationClient) FindByCode(code string) ([]Registration, error) {
	query := client.base.uri + registrationPath + "?code=" + code

	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Api-Key", client.base.token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	registrations := make([]Registration, 0)
	if err := json.NewDecoder(response.Body).Decode(&registrations); err != nil {
		return nil, errors.New("invalid response body")
	}

	return registrations, nil
}
