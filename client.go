package nmbrs

import (
	"net/http"

	"github.com/tim-online/go-nmbrs/companies"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-nmbrs/" + libraryVersion
)

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client: &soap.Client{
			Client:    httpClient,
			UserAgent: userAgent,
		},
	}

	c.Companies = companies.NewService()
	c.Companies.Client = c.client

	return c
}

type Client struct {
	// SOAP client used to communicate with the API.
	client *soap.Client

	// Services used for communicating with the API
	Companies *companies.Service
}
