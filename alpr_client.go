package toolkit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	alprPath = "/api/v1/alpr"
)

type ALPRClient struct {
	base *Client
}

func (client *ALPRClient) Recognize(url string) ([]ResultALPR, error) {
	query := client.base.uri + alprPath + "/private/recognize?image_url=" + url

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
		results := make([]ResultALPR, 0)

		if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
			return nil, errors.New("invalid response body")
		}

		return results, nil
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
