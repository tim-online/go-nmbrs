package costcenter

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx?op=CostCenter_Update

const (
	updateAction = "CostCenter_Update"
)

// Update the kostenplaatsen starting from given period.
func (s *Service) Update(requestBody *UpdateRequest) (*UpdateResponse, error) {
	responseBody := &UpdateResponse{}
	request, response := newUpdateAction(requestBody, responseBody)

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
	updateResponse, ok := response.Envelope.Body.Data.(*UpdateResponse)
	if ok == false {
		return updateResponse, soap.ErrBadResponse
	}

	return updateResponse, err
}

func newUpdateAction(requestBody *UpdateRequest, responseBody *UpdateResponse) (*soap.Request, *soap.Response) {
	requestBody.init()
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

func NewUpdateRequest() *UpdateRequest {
	req := &UpdateRequest{}
	req.init()
	return req
}

type UpdateRequest struct {
	XMLName xml.Name

	EmployeeID  int                  `xml:"EmployeeId"`
	CostCenters []EmployeeCostCenter `xml:"CostCenteren>EmployeeCostCenter"`
	Period      int                  `xml:"Period"`
	Year        int                  `xml:"Year"`
}

func (r *UpdateRequest) init() {
	r.XMLName = xml.Name{
		Space: xmlns,
		Local: updateAction,
	}
}

type UpdateResponse struct {
	Result int `xml:"CostCenter_UpdateResponse"`
}

type EmployeeCostCenter struct {
	ID          int         `xml:"Id,omitempty"`
	CostCenter  CostCenter  `xml:"CostCenter,omitempty"`
	Kostensoort Kostensoort `xml:"Kostensoort,omitempty"`
	Percentage  float64     `xml:"Percentage"`
	Default     bool        `xml:"Default"`
}

type CostCenter struct {
	Code        int    `xml:"Code"`
	Description string `xml:"Description"`
	ID          int    `xml:"Id,omitempty"`
}

type Kostensoort struct {
	Code        int    `xml:"Code"`
	Description string `xml:"Description"`
}
