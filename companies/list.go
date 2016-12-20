package companies

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listAction = "https://api.nmbrs.nl/soap/v2.1/CompanyService/List_GetAll"
)

// List all products
func (s *Service) List() (*listResponse, error) {
	request, response := newListAction()

	request.Envelope.Header.Data = s.AuthHeader

	httpReq, err := s.Client.NewRequest(s.Endpoint.String(), request)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, response)
	if err != nil {
		return nil, err
	}

	// @TODO: check if this can be better
	listResponse, ok := response.Envelope.Body.Data.(*listResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func newListAction() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	request.Action = url.MustParse(listAction)
	request.Envelope.Body.Data = newListRequest()

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListResponse()

	return request, response
}

type listRequest struct {
	XMLName xml.Name `xml:"List_GetAll"`
	Xmlns   string   `xml:"xmlns,attr"`
}

func newListRequest() *listRequest {
	return &listRequest{
		Xmlns: xmlns,
	}
}

type listResponse struct {
	Companies []Company `xml:"List_GetAllResult>Company"`
}

func newListResponse() *listResponse {
	return &listResponse{}
}
