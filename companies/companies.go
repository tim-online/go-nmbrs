package companies

import (
	"errors"
	"log"
	"net/url"

	"github.com/tim-online/go-nmbrs/auth"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	xmlns           = "https://api.nmbrs.nl/soap/v2.1/CompanyService"
	endpoint        = "https://api.nmbrs.nl/soap/v2.1/CompanyService.asmx"
	sandboxEndpoint = "https://api-sandbox.nmbrs.nl/soap/v2.1/CompanyService.asmx"
)

var (
	ErrBadResponse = errors.New("Bad response type")
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

	return &Service{
		Client:     soap.NewClient(),
		Endpoint:   endpoint,
		AuthHeader: authHeader,
	}
}
