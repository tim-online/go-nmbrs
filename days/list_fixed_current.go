package days

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listFixedCurrentAction = "DaysFixed_GetCurrent"
)

// List fixed days
func (s *Service) ListFixedCurrent(employeeID int) (*listFixedCurrentResponse, error) {
	// get a new request & response envelope
	request, response := newListFixedCurrentAction(employeeID)

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
	listFixedCurrentResponse, ok := response.Envelope.Body.Data.(*listFixedCurrentResponse)
	if ok == false {
		return listFixedCurrentResponse, soap.ErrBadResponse
	}

	return listFixedCurrentResponse, err
}

func newListFixedCurrentAction(employeeID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newListFixedCurrentRequest(employeeID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListFixedCurrentResponse()

	return request, response
}

type listFixedCurrentRequest struct {
	XMLName xml.Name `xml:"DaysFixed_GetCurrent"`

	EmployeeID int `xml:"EmployeeId"`
}

func newListFixedCurrentRequest(employeeID int) *listFixedCurrentRequest {
	return &listFixedCurrentRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listVarAction,
		},

		EmployeeID: employeeID,
	}
}

type listFixedCurrentResponse struct {
	Days int `xml:"DaysFixed_GetCurrentResult"`
}

func newListFixedCurrentResponse() *listFixedCurrentResponse {
	return &listFixedCurrentResponse{}
}
