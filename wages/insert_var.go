package wages

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	insertVarAction = "WageComponentVar_Insert"
)

// Set variable days of period
func (s *Service) InsertVar(requestBody *InsertVarRequest) (*InsertVarResponse, error) {
	responseBody := &InsertVarResponse{}
	request, response := NewInsertVarAction(requestBody, responseBody)

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
	InsertVarResponse, ok := response.Envelope.Body.Data.(*InsertVarResponse)
	if ok == false {
		return InsertVarResponse, soap.ErrBadResponse
	}

	return InsertVarResponse, err
}

func NewInsertVarAction(requestBody *InsertVarRequest, responseBody *InsertVarResponse) (*soap.Request, *soap.Response) {
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

func NewInsertVarRequest() *InsertVarRequest {
	return &InsertVarRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: getVarAction,
		},
	}
}

type InsertVarRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`

	ID           int     `xml:"Id,omitempty"`
	Code         int     `xml:"WageComponent>Code,omitempty"`
	Value        float64 `xml:"WageComponent>Value,omitempty"`
	CostCenterID int     `xml:"WageComponent>CostCenterId,omitempty"`
	CostUnitID   int     `xml:"WageComponent>CostUnitId,omitempty"`
	Comment      string  `xml:"WageComponent>Comment"`
	DocumentName string  `xml:"WageComponent>DocumentName"`
	// Attachment base64Binary `xml:"WageComponent>Attachment"`

	Period          int  `xml:"Period"`
	Year            int  `xml:"Year"`
	UnprotectedMode bool `xml:"UnprotectedMode"`
}

type InsertVarResponse struct {
}
