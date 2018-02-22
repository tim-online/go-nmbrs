package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	department_GetCurrentAction = "Department_GetCurrent"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Department_GetCurrent(employeeID int) (*Department_GetCurrentResponse, error) {
	// get a new request & response envelope
	request, response := s.NewDepartment_GetCurrentAction(employeeID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Department_GetCurrentResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewDepartment_GetCurrentAction(employeeID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewDepartment_GetCurrentRequest(employeeID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewDepartment_GetCurrentResponse()

	return request, response
}

type Department_GetCurrentRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
}

func NewDepartment_GetCurrentRequest(employeeID int) *Department_GetCurrentRequest {
	return &Department_GetCurrentRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: department_GetCurrentAction,
		},

		EmployeeID: employeeID,
	}
}

type Department_GetCurrentResponse struct {
	Department Department `xml:"Department_GetCurrentResult"`
}

func NewDepartment_GetCurrentResponse() *Department_GetCurrentResponse {
	return &Department_GetCurrentResponse{}
}

type Department struct {
	ID          int    `xml:"Id"`
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}
