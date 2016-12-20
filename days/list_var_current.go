package days

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listVarCurrentAction = "https://api.nmbrs.nl/soap/v2.1/EmployeeService/DaysVar_GetCurrent"
)

// List variable days
func (s *Service) ListVarCurrent(employeeID int) (*listVarCurrentResponse, error) {
	// get a new request & response envelope
	request, response := newListVarCurrentAction(employeeID)

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
	listVarCurrentResponse, ok := response.Envelope.Body.Data.(*listVarCurrentResponse)
	if ok == false {
		return listVarCurrentResponse, soap.ErrBadResponse
	}

	return listVarCurrentResponse, err
}

func newListVarCurrentAction(employeeID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	request.Action = url.MustParse(listVarCurrentAction)
	request.Envelope.Body.Data = newListVarCurrentRequest(employeeID)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListVarCurrentResponse()

	return request, response
}

type listVarCurrentRequest struct {
	XMLName xml.Name `xml:"DaysVar_GetCurrent"`
	Xmlns   string   `xml:"xmlns,attr"`

	EmployeeID int `xml:"EmployeeId"`
}

func newListVarCurrentRequest(employeeID int) *listVarCurrentRequest {
	return &listVarCurrentRequest{
		Xmlns:      xmlns,
		EmployeeID: employeeID,
	}
}

type listVarCurrentResponse struct {
	Days int `xml:"DaysVar_GetCurrentResult"`
}

func newListVarCurrentResponse() *listVarCurrentResponse {
	return &listVarCurrentResponse{}
}

