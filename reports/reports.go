package reports

import (
	"log"
	"net/url"

	"github.com/tim-online/go-nmbrs/auth"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	xmlns           = "https://api.nmbrs.nl/soap/v2.1/ReportService"
	Endpoint        = "https://api.nmbrs.nl/soap/v2.1/ReportService.asmx"
	SandboxEndpoint = "https://api-sandbox.nmbrs.nl/soap/v2.1/ReportService.asmx"
)

type Service struct {
	Client     *soap.Client
	Endpoint   *url.URL
	AuthHeader *auth.AuthHeader
}

func NewService(authHeader *auth.AuthHeader) *Service {
	endpoint, err := url.ParseRequestURI(Endpoint)
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
