package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Schedule_GetAll_AllEmployeesByCompanyAction = "Schedule_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Schedule_GetAll_AllEmployeesByCompany(companyID int) (*Schedule_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewSchedule_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Schedule_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewSchedule_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewSchedule_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewSchedule_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type Schedule_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewSchedule_GetAll_AllEmployeesByCompanyRequest(companyID int) *Schedule_GetAll_AllEmployeesByCompanyRequest {
	return &Schedule_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Schedule_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type Schedule_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeeScheduleItems EmployeeScheduleItems `xml:"Schedule_GetAll_AllEmployeesByCompanyResult>EmployeeScheduleItem"`
}

func NewSchedule_GetAll_AllEmployeesByCompanyResponse() *Schedule_GetAll_AllEmployeesByCompanyResponse {
	return &Schedule_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeeScheduleItems []EmployeeScheduleItem

type EmployeeScheduleItem struct {
	EmployeeID        int `xml:"EmployeeId"`
	EmployeeSchedules []struct {
		ID                 string   `xml:"Id"`
		HoursMonday        float64  `xml:"HoursMonday"`
		HoursTuesday       float64  `xml:"HoursTuesday"`
		HoursWednesday     float64  `xml:"HoursWednesday"`
		HoursThursday      float64  `xml:"HoursThursday"`
		HoursFriday        float64  `xml:"HoursFriday"`
		HoursSaturday      float64  `xml:"HoursSaturday"`
		HoursSunday        float64  `xml:"HoursSunday"`
		HoursMonday2       float64  `xml:"HoursMonday2"`
		HoursTuesday2      float64  `xml:"HoursTuesday2"`
		HoursWednesday2    float64  `xml:"HoursWednesday2"`
		HoursThursday2     float64  `xml:"HoursThursday2"`
		HoursFriday2       float64  `xml:"HoursFriday2"`
		HoursSaturday2     float64  `xml:"HoursSaturday2"`
		HoursSunday2       float64  `xml:"HoursSunday2"`
		ParttimePercentage float64  `xml:"ParttimePercentage"`
		StartDate          lib.Time `xml:"StartDate"`
		CreationDate       lib.Time `xml:"CreationDate"`
	} `xml:"EmployeeSchedules>Schedule_V2"`
}
