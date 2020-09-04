package toolkit

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	wantedPath = "/api/v1/wanted/vehicles"
)

type WantedClient struct {
	base *Client
}

func (client *WantedClient) FindByVIN(vin string) ([]WantedVehicle, error) {
	query := client.base.uri + wantedPath + "?vin=" + vin

	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(APIKeyHeader, client.base.token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	vehicles := make([]WantedVehicle, 0)
	if err := json.NewDecoder(response.Body).Decode(&vehicles); err != nil {
		return nil, errors.New("invalid response body")
	}

	return vehicles, nil
}

func (client *WantedClient) FindByNumber(number string) ([]WantedVehicle, error) {
	query := client.base.uri + wantedPath + "?number=" + number

	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(APIKeyHeader, client.base.token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	vehicles := make([]WantedVehicle, 0)
	if err := json.NewDecoder(response.Body).Decode(&vehicles); err != nil {
		return nil, errors.New("invalid response body")
	}

	return vehicles, nil
}
