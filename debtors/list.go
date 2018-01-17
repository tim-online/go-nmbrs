package debtors

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listAction = "List_GetAll"
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
	body := newListRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListResponse()

	return request, response
}

type listRequest struct {
	XMLName xml.Name
}

func newListRequest() *listRequest {
	return &listRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listAction,
		},
	}
}

type listResponse struct {
	Debtors []Debtor `xml:"List_GetAllResult>Debtor"`
}

func newListResponse() *listResponse {
	return &listResponse{}
}

type Debtor struct {
	ID     int    `xml:"Id"`
	Number string `xml:"Number"`
	Name   string `xml:"Name"`
}
