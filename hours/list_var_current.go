package hours

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listVarCurrentAction = "HourComponentVar_GetCurrent"
)

// List fixed hour components
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
	body := newListVarCurrentRequest(employeeID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListVarCurrentResponse()

	return request, response
}

type listVarCurrentRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
	Period     int `xml:"Period"`
	Year       int `xml:"Year"`
}

func newListVarCurrentRequest(employeeID int) *listVarCurrentRequest {
	return &listVarCurrentRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listVarCurrentAction,
		},
		EmployeeID: employeeID,
	}
}

type listVarCurrentResponse struct {
	HourComponents []HourComponent `xml:"HourComponentVar_GetCurrentResult>HourComponent"`
}

func newListVarCurrentResponse() *listVarCurrentResponse {
	return &listVarCurrentResponse{}
}
