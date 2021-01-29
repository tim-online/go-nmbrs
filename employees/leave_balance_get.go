package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	leaveBalanceGetAction = "LeaveBalance_Get"
)

// LeaveBalanceGet gets all employees that belong to a company and are active as given.
func (s *Service) LeaveBalanceGet(employeeID int) (*leaveBalanceGetResponse, error) {
	// get a new request & response envelope
	request, response := newLeaveBalanceGetAction(employeeID)

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
	listResponse, ok := response.Envelope.Body.Data.(*leaveBalanceGetResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func newLeaveBalanceGetAction(employeeID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newLeaveBalanceGetRequest(employeeID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newLeaveBalanceGetResponse()

	return request, response
}

type leaveBalanceGetRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
}

func newLeaveBalanceGetRequest(employeeID int) *leaveBalanceGetRequest {
	return &leaveBalanceGetRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: leaveBalanceGetAction,
		},

		EmployeeID: employeeID,
	}
}

type leaveBalanceGetResponse struct {
	LeaveBalance LeaveBalance `xml:"LeaveBalance_GetResult"`
}

func newLeaveBalanceGetResponse() *leaveBalanceGetResponse {
	return &leaveBalanceGetResponse{}
}

type LeaveBalance struct {
	DecMedewerkerVerlofStartSaldo float64 `xml:"DecMedewerkerVerlofStartSaldo"`
	DecMedewerkerVerlofCurrSaldo  float64 `xml:"DecMedewerkerVerlofCurrSaldo"`
}
