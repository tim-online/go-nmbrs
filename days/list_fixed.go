package days

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listFixedAction = "DaysFixed_Get"
)

// List fixed days
func (s *Service) ListFixed(employeeID int, period int, year int) (*listFixedResponse, error) {
	// get a new request & response envelope
	request, response := newListFixedAction(employeeID, period, year)

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
	listFixedResponse, ok := response.Envelope.Body.Data.(*listFixedResponse)
	if ok == false {
		return listFixedResponse, soap.ErrBadResponse
	}

	return listFixedResponse, err
}

func newListFixedAction(employeeID int, period int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newListFixedRequest(employeeID, period, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListFixedResponse()

	return request, response
}

type listFixedRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
	Period     int `xml:"Period"`
	Year       int `xml:"Year"`
}

func newListFixedRequest(employeeID int, period int, year int) *listFixedRequest {
	return &listFixedRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listFixedAction,
		},
		EmployeeID: employeeID,
		Period:     period,
		Year:       year,
	}
}

type listFixedResponse struct {
	Days int `xml:"DaysFixed_GetResult"`
}

func newListFixedResponse() *listFixedResponse {
	return &listFixedResponse{}
}
