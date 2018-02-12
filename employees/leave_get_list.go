package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	leave_GetListAction = "Leave_GetList"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Leave_GetList(employeeID int, typ string, soort string, year int) (*Leave_GetListResponse, error) {
	// get a new request & response envelope
	request, response := s.NewLeave_GetListAction(employeeID, typ, soort, year)

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
	listResponse, ok := response.Envelope.Body.Data.(*Leave_GetListResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewLeave_GetListAction(employeeID int, typ string, soort string, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewLeave_GetListRequest(employeeID, typ, soort, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewLeave_GetListResponse()

	return request, response
}

type Leave_GetListRequest struct {
	XMLName xml.Name

	EmployeeID int    `xml:"EmployeeId"`
	Type       string `xml:"Type"`
	Soort      string `xml:"Soort"`
	Year       int    `xml:"Year"`
}

func NewLeave_GetListRequest(employeeID int, typ string, soort string, year int) *Leave_GetListRequest {
	return &Leave_GetListRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: leave_GetListAction,
		},

		EmployeeID: employeeID,
		Type:       typ,
		Soort:      soort,
		Year:       year,
	}
}

type Leave_GetListResponse struct {
	Leaves []Leave `xml:"Leave_GetListResult>Leave"`
}

func NewLeave_GetListResponse() *Leave_GetListResponse {
	return &Leave_GetListResponse{}
}

type Leave struct {
	Description string   `xml:"Description"`
	Hours       float64  `xml:"Hours"`
	UsageType   string   `xml:"UsageType"`
	Start       lib.Time `xml:"Start"`
	End         lib.Time `xml:"End"`
	StartHours  float64  `xml:"StartHours"`
	EndHours    float64  `xml:"EndHours"`
	Type        string   `xml:"Type"`
	Status      string   `xml:"Status"`
}
