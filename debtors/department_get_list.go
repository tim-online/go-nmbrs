package debtors

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Department_GetListAction = "Department_GetList"
)

// Returns departments that belong to a debtor.
func (s *Service) Department_GetList(debtorID int) (*Department_GetListResponse, error) {
	// get a new request & response envelope
	request, response := s.NewDepartment_GetListAction(debtorID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Department_GetListResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewDepartment_GetListAction(debtorID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewDepartment_GetListRequest(debtorID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewDepartment_GetListResponse()

	return request, response
}

type Department_GetListRequest struct {
	XMLName xml.Name

	DebtorID int `xml:"DebtorId"`
}

func NewDepartment_GetListRequest(debtorID int) *Department_GetListRequest {
	return &Department_GetListRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Department_GetListAction,
		},

		DebtorID: debtorID,
	}
}

type Department_GetListResponse struct {
	Departments []Department `xml:"Department_GetListResult>Department"`
}

func NewDepartment_GetListResponse() *Department_GetListResponse {
	return &Department_GetListResponse{}
}

type Department struct {
	ID          int
	Code        string
	Description string
}
