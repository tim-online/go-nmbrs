package days

import (
	"log"
	"net/url"

	"github.com/tim-online/go-nmbrs/auth"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	xmlns           = "https://api.nmbrs.nl/soap/v2.1/EmployeeService"
	endpoint        = "https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx"
	sandboxEndpoint = "https://api-sandbox.nmbrs.nl/soap/v2.1/EmployeeService.asmx"
)

type Service struct {
	Client     *soap.Client
	Endpoint   *url.URL
	AuthHeader *auth.AuthHeader
}

func NewService(authHeader *auth.AuthHeader) *Service {
	endpoint, err := url.ParseRequestURI(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	authHeader.Xmlns = xmlns

	return &Service{
		Client:     soap.NewClient(nil),
		Endpoint:   endpoint,
		AuthHeader: authHeader,
	}
}
