package costcenter

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx?op=CostCenter_GetCurrent

const (
	getCurrentAction = "CostCenter_GetCurrent"
)

// Get the kostenplaatsen for an employee for the current period
func (s *Service) GetCurrent(requestBody *GetCurrentRequest) (*GetCurrentResponse, error) {
	responseBody := &GetCurrentResponse{}
	request, response := newGetCurrentAction(requestBody, responseBody)

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
	getCurrentResponse, ok := response.Envelope.Body.Data.(*GetCurrentResponse)
	if ok == false {
		return getCurrentResponse, soap.ErrBadResponse
	}

	return getCurrentResponse, err
}

func newGetCurrentAction(requestBody *GetCurrentRequest, responseBody *GetCurrentResponse) (*soap.Request, *soap.Response) {
	requestBody.init()
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

func (s *Service) NewGetCurrentRequest() *GetCurrentRequest {
	req := &GetCurrentRequest{}
	req.init()
	return req
}

type GetCurrentRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
}

func (r *GetCurrentRequest) init() {
	r.XMLName = xml.Name{
		Space: xmlns,
		Local: getCurrentAction,
	}
}

type GetCurrentResponse struct {
	CostCenters []EmployeeCostCenter `xml:"CostCenter_GetCurrentResult>EmployeeCostCenter"`
}
