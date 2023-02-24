package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	TypeListAction = "EmployeeType_GetList"
)

// TypeList gets all employees that belong to a company and are active as given.
func (s *Service) TypeList() (*TypeListResponse, error) {
	// get a new request & response envelope
	request, response := newTypeListAction()

	// copy authheader to new envelope
	request.Envelope.Header.Data = s.AuthHeader

	// create a new HTTP request from the SOAP envelope
	httpReq, err := s.Client.NewRequest(s.Endpoint.String(), request)
	if err != nil {
		return nil, err
	}

	// make the HTTP request
	_, err = s.Client.Do(httpReq, response)
	if err != nil {
		return nil, err
	}

	// @TODO: check if this can be better
	listResponse, ok := response.Envelope.Body.Data.(*TypeListResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func newTypeListAction() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newTypeListRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newTypeListResponse()

	return request, response
}

type TypeListRequest struct {
	XMLName xml.Name
}

func newTypeListRequest() *TypeListRequest {
	return &TypeListRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: TypeListAction,
		},
	}
}

type TypeListResponse struct {
	EmployeeTypes EmployeeTypes `xml:"EmployeeType_GetListResult>EmployeeType"`
}

func newTypeListResponse() *TypeListResponse {
	return &TypeListResponse{}
}
