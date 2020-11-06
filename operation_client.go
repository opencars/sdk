package toolkit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	operationsPath = "/api/v1/operations"
)

type OperationClient struct {
	base *Client
}

func (client *OperationClient) FindByNumber(number string) ([]Operation, error) {
	query := client.base.uri + operationsPath + "?number=" + url.QueryEscape(number)

	req, err := http.NewRequest(http.MethodGet, query, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(APIKeyHeader, client.base.token)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		operations := make([]Operation, 0)
		if err := json.NewDecoder(response.Body).Decode(&operations); err != nil {
			return nil, errors.New("invalid response body")
		}

		return operations, nil
	case http.StatusUnauthorized:
		var e Error

		if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
			return nil, ErrUnauthorized
		}

		return nil, fmt.Errorf("unauthorized: %s", e.Message)
	default:
		var e Error

		if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("failed with status %d", response.StatusCode)
		}

		return nil, fmt.Errorf("failed: %s", e.Message)
	}
}
