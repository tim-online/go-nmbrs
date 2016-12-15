package companies

import (
	"encoding/xml"
	"log"
	"net/url"

	"github.com/tim-online/go-nmbrs/soap"
)

const (
	xmlns           = "https://api.nmbrs.nl/soap/v2.1/CompanyService"
	endpoint        = "https://api.nmbrs.nl/soap/v2.1/CompanyService.asmx"
	sandboxEndpoint = "https://api-sandbox.nmbrs.nl/soap/v2.1/CompanyService.asmx"
)

type Service struct {
	Client   *soap.Client
	Endpoint *url.URL
}

func NewService() *Service {
	endpoint, err := url.ParseRequestURI(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		Client:   soap.NewClient(),
		Endpoint: endpoint,
	}
}

// List all products
func (s *Service) List() (*listResponse, error) {
	listRequest := listRequest{}
	req, err := s.Client.NewRequest(s.Endpoint.String(), listRequest)
	if err != nil {
		return nil, err
	}

	listResponse := newListResponse()
	_, err = s.Client.Do(req, listResponse)
	if err != nil {
		return nil, err
	}

	return listResponse, err
}

type listRequest struct {
	XMLName xml.Name `xml:"List_GetAll"`
}

type listResponse struct {
}

func newListResponse() *listResponse {
	return &listResponse{}
}
