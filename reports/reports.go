package reports

import (
	"log"
	"net/url"

	"github.com/tim-online/go-nmbrs/auth"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	xmlns           = "https://api.nmbrs.nl/soap/v3/ReportService"
	Endpoint        = "https://api.nmbrs.nl/soap/v3/ReportService.asmx"
	SandboxEndpoint = "https://api-sandbox.nmbrs.nl/soap/v3/ReportService.asmx"
)

type Service struct {
	Client     *soap.Client
	Endpoint   *url.URL
	AuthHeader *auth.AuthHeaderWithDomain
}

func NewService(authHeader *auth.AuthHeaderWithDomain) *Service {
	endpoint, err := url.ParseRequestURI(Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// set correct xmlns
	authHeaderCopy := &auth.AuthHeaderWithDomain{}
	*authHeaderCopy = *authHeader
	authHeaderCopy.XMLName.Space = xmlns

	return &Service{
		Client:     soap.NewClient(nil),
		Endpoint:   endpoint,
		AuthHeader: authHeaderCopy,
	}
}
