package toolkit

type Client struct {
	uri, token string

	registrationClient *RegistrationClient
	operationClient    *OperationClient
	wantedClient       *WantedClient
}

// New initializes new instance of client structure.
func New(uri string, token string) *Client {
	client := new(Client)

	client.uri = uri
	client.token = token

	return client
}

func (client *Client) Registration() *RegistrationClient {
	if client.registrationClient == nil {
		client.registrationClient = &RegistrationClient{
			base: client,
		}
	}

	return client.registrationClient
}

func (client *Client) Operation() *OperationClient {
	if client.operationClient == nil {
		client.operationClient = &OperationClient{
			base: client,
		}
	}

	return client.operationClient
}

func (client *Client) Wanted() *WantedClient {
	if client.wantedClient == nil {
		client.wantedClient = &WantedClient{
			base: client,
		}
	}

	return client.wantedClient
}

//
//func (client *Client) searchRegistrations(code string) ([]Registration, error) {
//	if code == "" {
//		return nil, errors.New("code is empty")
//	}
//
//	query := client.uri + registrationsPath + "?code=" + code
//
//	req, err := http.NewRequest(http.MethodGet, query, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Api-Key", client.token)
//
//	response, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	models := make([]Registration, 0)
//	if err := json.NewDecoder(response.Body).Decode(&models); err != nil {
//		return nil, errors.New("invalid response body")
//	}
//
//	return models, nil
//}
//
//// Operation makes API request to get operations details by government number.
//// Returns first operation from OpenCars operations tables.
//func (client *Client) Operation(number string) (*Operation, error) {
//	operation, err := client.searchOperations(number, 1)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &operation[0], nil
//}
//
//// Registrations makes API request to get registration info by unique transport cerificate code.
//// Returns all operations from OpenCars operations tables.
//func (client *Client) Registrations(code string) ([]Registration, error) {
//	registrations, err := client.searchRegistrations(code)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return registrations, nil
//}
//
//// Operations makes API request to get operations details by government number.
//func (client *Client) Operations(number string, limit int) ([]Operation, error) {
//	return client.searchOperations(number, limit)
//}
