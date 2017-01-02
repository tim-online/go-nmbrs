package days

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	setVarAction = "DaysVar_Set"
)

// Set variable days of period
func (s *Service) SetVar(employeeID int, days int, year int, period int) (*setVarResponse, error) {
	// get a new request & response envelope
	requestBody := &setVarRequest{
		EmployeeID: employeeID,
		Days:       days,
		Year:       year,
		Period:     period,
	}

	responseBody := &setVarResponse{}
	request, response := newSetVarAction(requestBody, responseBody)

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
	setVarResponse, ok := response.Envelope.Body.Data.(*setVarResponse)
	if ok == false {
		return setVarResponse, soap.ErrBadResponse
	}

	return setVarResponse, err
}

// func newSetVarAction(employeeID int, days int, year int, period int) (*soap.Request, *soap.Response) {
// 	request := soap.NewRequest()
// 	body := newSetVarRequest(employeeID, days, year, period)
// 	request.Envelope.Body.Data = body
// 	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

// 	response := soap.NewResponse()
// 	response.Envelope.Body.Data = newSetVarResponse()

// 	return request, response
// }

func newSetVarAction(requestBody *setVarRequest, responseBody *setVarResponse) (*soap.Request, *soap.Response) {
	requestBody.XMLName = xml.Name{
		Space: xmlns,
		Local: setVarAction,
	}
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

type setVarRequest struct {
	XMLName xml.Name `xml:"DaysVar_Set"`

	EmployeeID int `xml:"EmployeeId"`
	Days       int `xml:"Days"`
	Year       int `xml:"Year"`
	Period     int `xml:"Period"`
}

type setVarResponse struct {
}
