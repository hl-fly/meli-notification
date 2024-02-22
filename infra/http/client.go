package http

import (
	"net/http"

	"github.com/hector-leite/meli-notification/domain/contract"
)

type httpClient struct {
	client *http.Client
}

func New(client *http.Client) contract.HTTPClient {
	return &httpClient{
		client: client,
	}
}

func (s httpClient) Do(req *http.Request) (*http.Response, error) {
	response, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
