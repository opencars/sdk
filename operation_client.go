package toolkit

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	operationsPath = "/api/v1/operations"
)

type OperationClient struct {
	base *Client
}

func (client *OperationClient) FindByNumber(number string) ([]Operation, error) {
	query := client.base.uri + operationsPath + "?number=" + number

	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Api-Key", client.base.token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	operations := make([]Operation, 0)
	if err := json.NewDecoder(response.Body).Decode(&operations); err != nil {
		return nil, errors.New("invalid response body")
	}

	return operations, nil
}
