package hours

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listVarAction = "HourComponentVar_Get"
)

// List variable hour components
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
	body := newListVarRequest(employeeID, period, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListVarResponse()

	return request, response
}

type listVarRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
	Period     int `xml:"Period"`
	Year       int `xml:"Year"`
}

func newListVarRequest(employeeID int, period int, year int) *listVarRequest {
	return &listVarRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listVarAction,
		},
		EmployeeID: employeeID,
		Period:     period,
		Year:       year,
	}
}

type listVarResponse struct {
	HourComponents []HourComponent `xml:"HourComponentVar_GetResult>HourComponent"`
}

func newListVarResponse() *listVarResponse {
	return &listVarResponse{}
}
