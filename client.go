package toolkit

type Client struct {
	uri, token string

	registrationClient *RegistrationClient
	operationClient    *OperationClient
	wantedClient       *WantedClient
	alprClient         *ALPRClient
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

func (client *Client) ALPR() *ALPRClient {
	if client.alprClient == nil {
		client.alprClient = &ALPRClient{
			base: client,
		}
	}

	return client.alprClient
}
