package costcenter

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx?op=CostCenter_Get

const (
	getAction = "CostCenter_Get"
)

// Get the kostenplaatsen for an employee
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
	req := &GetRequest{}
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
	CostCenters []EmployeeCostCenter `xml:"CostCenter_GetResult>EmployeeCostCenter"`
}
