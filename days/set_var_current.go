package days

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	setVarCurrentAction = "https://api.nmbrs.nl/soap/v2.1/EmployeeService/DaysVar_SetCurrent"
)

// Set variable days of current period
func (s *Service) SetVarCurrent(employeeID int, days int) (*setVarCurrentResponse, error) {
	// get a new request & response envelope
	request, response := newSetVarCurrentAction(employeeID, days)

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
	setVarCurrentResponse, ok := response.Envelope.Body.Data.(*setVarCurrentResponse)
	if ok == false {
		return setVarCurrentResponse, soap.ErrBadResponse
	}

	return setVarCurrentResponse, err
}

func newSetVarCurrentAction(employeeID int, days int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	request.Action = url.MustParse(setVarCurrentAction)
	request.Envelope.Body.Data = newSetVarCurrentRequest(employeeID, days)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newSetVarCurrentResponse()

	return request, response
}

type setVarCurrentRequest struct {
	XMLName xml.Name `xml:"DaysVar_SetCurrent"`
	Xmlns   string   `xml:"xmlns,attr"`

	EmployeeID int `xml:"EmployeeId"`
	Days       int `xml:"Days"`
}

func newSetVarCurrentRequest(employeeID int, days int) *setVarCurrentRequest {
	return &setVarCurrentRequest{
		Xmlns:      xmlns,
		EmployeeID: employeeID,
		Days:       days,
	}
}

type setVarCurrentResponse struct {
}

func newSetVarCurrentResponse() *setVarCurrentResponse {
	return &setVarCurrentResponse{}
}
