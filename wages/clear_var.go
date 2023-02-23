package wages

// https://api.nmbrs.nl/soap/v3/EmployeeService.asmx?op=WageComponentVar_Clear

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	clearVarAction = "WageComponentVar_Clear"
)

// Set variable days of period
func (s *Service) ClearVar(requestBody *ClearVarRequest) (*clearVarResponse, error) {
	responseBody := &clearVarResponse{}
	request, response := newClearVarAction(requestBody, responseBody)

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
	clearVarResponse, ok := response.Envelope.Body.Data.(*clearVarResponse)
	if ok == false {
		return clearVarResponse, soap.ErrBadResponse
	}

	return clearVarResponse, err
}

func newClearVarAction(requestBody *ClearVarRequest, responseBody *clearVarResponse) (*soap.Request, *soap.Response) {
	requestBody.XMLName = xml.Name{
		Space: xmlns,
		Local: clearVarAction,
	}
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

type ClearVarRequest struct {
	XMLName xml.Name

	EmployeeID      int  `xml:"EmployeeId"`
	Period          int  `xml:"Period"`
	Year            int  `xml:"Year"`
	UnprotectedMode bool `xml:"UnprotectedMode"`
}

type clearVarResponse struct {
}
