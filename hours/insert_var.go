package hours

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	insertVarAction = "HourComponentVar_Insert"
)

// Set variable days of period
func (s *Service) InsertVar(requestBody *InsertVarRequest) (*insertVarResponse, error) {
	responseBody := &insertVarResponse{}
	request, response := newInsertVarAction(requestBody, responseBody)

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
	insertVarResponse, ok := response.Envelope.Body.Data.(*insertVarResponse)
	if ok == false {
		return insertVarResponse, soap.ErrBadResponse
	}

	return insertVarResponse, err
}

func newInsertVarAction(requestBody *InsertVarRequest, responseBody *insertVarResponse) (*soap.Request, *soap.Response) {
	requestBody.XMLName = xml.Name{
		Space: xmlns,
		Local: insertVarAction,
	}
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

type InsertVarRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`

	ID           int     `xml:"Id,omitempty"`
	HourCode     int     `xml:"Hourcomponent>HourCode,omitempty"`
	Hours        float64 `xml:"Hourcomponent>Hours,omitempty"`
	CostCenterID int     `xml:"Hourcomponent>CostCenterId,omitempty"`
	CostUnitID   int     `xml:"Hourcomponent>CostUnitId,omitempty"`
	Comment      string  `xml:"Hourcomponent>Comment"`
	DocumentName string  `xml:"Hourcomponent>DocumentName"`
	// Attachment base64Binary `xml:"Hourcomponent>Attachment"`

	Period          int  `xml:"Period"`
	Year            int  `xml:"Year"`
	UnprotectedMode bool `xml:"UnprotectedMode"`
}

type insertVarResponse struct {
}
