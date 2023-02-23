package wages

// https://api.nmbrs.nl/soap/v3/EmployeeService.asmx?op=WageComponentVar_Get

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	getVarAction = "WageComponentVar_Get"
)

// Set variable days of period
func (s *Service) GetVar(requestBody *GetVarRequest) (*GetVarResponse, error) {
	responseBody := &GetVarResponse{}
	request, response := NewGetVarAction(requestBody, responseBody)

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
		// faultcode 9999 (TWK exception) = period is protected, use
		// UnprotectedMode
		return nil, err
	}

	// @TODO: check if this can be better
	GetVarResponse, ok := response.Envelope.Body.Data.(*GetVarResponse)
	if ok == false {
		return GetVarResponse, soap.ErrBadResponse
	}

	return GetVarResponse, err
}

func NewGetVarAction(requestBody *GetVarRequest, responseBody *GetVarResponse) (*soap.Request, *soap.Response) {
	requestBody.XMLName = xml.Name{
		Space: xmlns,
		Local: getVarAction,
	}
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

func NewGetVarRequest(employeeID int, year int, period int) *GetVarRequest {
	return &GetVarRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: getVarAction,
		},
		EmployeeID: employeeID,
		Period:     period,
		Year:       year,
	}
}

type GetVarRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
	Period     int `xml:"Period"`
	Year       int `xml:"Year"`
}

type GetVarResponse struct {
	WageComponents []WageComponent `xml:"WageComponentVar_GetResult>WageComponent"`
}

type WageComponent struct {
	ID    int     `xml:"Id"`
	Code  int     `xml:"Code"`
	Value float64 `xml:"Value"`
}
