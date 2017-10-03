package schedules

import (
	"encoding/xml"
	"time"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx?op=Schedule_Get

const (
	getAction = "Schedule_Get"
)

// Get schedule the active schedule for given period.
func (s *Service) Get(requestBody *GetRequest) (*GetResponse, error) {
	responseBody := &GetResponse{}
	request, response := newGetAction(requestBody, responseBody)

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
	getResponse, ok := response.Envelope.Body.Data.(*GetResponse)
	if ok == false {
		return getResponse, soap.ErrBadResponse
	}

	return getResponse, err
}

func newGetAction(requestBody *GetRequest, responseBody *GetResponse) (*soap.Request, *soap.Response) {
	requestBody.init()
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

func (s *Service) NewGetRequest() *GetRequest {
	now := time.Now()
	req := &GetRequest{
		EmployeeID: 0,
		Period:     int(now.Month()),
		Year:       now.Year(),
	}
	req.init()
	return req
}

type GetRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
	Period     int `xml:"Period"`
	Year       int `xml:"Year"`
}

func (r *GetRequest) init() {
	r.XMLName = xml.Name{
		Space: xmlns,
		Local: getAction,
	}
}

type GetResponse struct {
	Schedule Schedule `xml:"Schedule_GetResult"`
}

type Schedule struct {
	HoursMonday        float64   `xml:"HoursMonday"`
	HoursTuesday       float64   `xml:"HoursTuesday"`
	HoursWednesday     float64   `xml:"HoursWednesday"`
	HoursThursday      float64   `xml:"HoursThursday"`
	HoursFriday        float64   `xml:"HoursFriday"`
	HoursSaturday      float64   `xml:"HoursSaturday"`
	HoursSunday        float64   `xml:"HoursSunday"`
	HoursMonday2       float64   `xml:"HoursMonday2"`
	HoursTuesday2      float64   `xml:"HoursTuesday2"`
	HoursWednesday2    float64   `xml:"HoursWednesday2"`
	HoursFriday2       float64   `xml:"HoursFriday2"`
	HoursThursday2     float64   `xml:"HoursThursday2"`
	HoursSaturday2     float64   `xml:"HoursSaturday2"`
	HoursSunday2       float64   `xml:"HoursSunday2"`
	ParttimePercentage float64   `xml:"ParttimePercentage"`
	StartDate          *lib.Time `xml:"StartDate"`
}
