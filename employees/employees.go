package employees

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

// Service handles communication with the employee related methods of the Nmbrs
// API
type Service struct {
	Client     *soap.Client
	Endpoint   *url.URL
	AuthHeader *auth.AuthHeader
}

// NewService returns a new employee service with a default SOAP client and a
// copy of the auth header
func NewService(authHeader *auth.AuthHeader) *Service {
	endpoint, err := url.ParseRequestURI(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// set correct xmlns
	authHeaderCopy := &auth.AuthHeader{}
	*authHeaderCopy = *authHeader
	authHeaderCopy.XMLName.Space = xmlns

	return &Service{
		Client:     soap.NewClient(nil),
		Endpoint:   endpoint,
		AuthHeader: authHeaderCopy,
	}
}
