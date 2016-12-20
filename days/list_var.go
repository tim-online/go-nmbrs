package days

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listVarAction = "https://api.nmbrs.nl/soap/v2.1/EmployeeService/DaysVar_Get"
)

// List variable days
func (s *Service) ListVar(employeeID int, period int, year int) (*listVarResponse, error) {
	// get a new request & response envelope
	request, response := newListVarAction(employeeID, period, year)

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
	listVarResponse, ok := response.Envelope.Body.Data.(*listVarResponse)
	if ok == false {
		return listVarResponse, soap.ErrBadResponse
	}

	return listVarResponse, err
}

func newListVarAction(employeeID int, period int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	request.Action = url.MustParse(listVarAction)
	request.Envelope.Body.Data = newListVarRequest(employeeID, period, year)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListVarResponse()

	return request, response
}

type listVarRequest struct {
	XMLName xml.Name `xml:"DaysVar_Get"`
	Xmlns   string   `xml:"xmlns,attr"`

	EmployeeID int `xml:"EmployeeId"`
	Period     int `xml:"Period"`
	Year       int `xml:"Year"`
}

func newListVarRequest(employeeID int, period int, year int) *listVarRequest {
	return &listVarRequest{
		Xmlns:      xmlns,
		EmployeeID: employeeID,
		Period:     period,
		Year:       year,
	}
}

type listVarResponse struct {
	Days int `xml:"DaysVar_GetResult"`
}

func newListVarResponse() *listVarResponse {
	return &listVarResponse{}
}
