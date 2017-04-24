package wages

// https://api.nmbrs.nl/soap/v2.1/EmployeeService.asmx?op=WageComponentVar_Insert_Batch

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	insertVarBatchAction = "WageComponentVar_Insert_Batch"
)

// Set variable days of period
func (s *Service) InsertVarBatch(requestBody *InsertVarBatchRequest) (*insertVarBatchResponse, error) {
	responseBody := &insertVarBatchResponse{}
	request, response := newInsertVarBatchAction(requestBody, responseBody)

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
	insertVarBatchResponse, ok := response.Envelope.Body.Data.(*insertVarBatchResponse)
	if ok == false {
		return insertVarBatchResponse, soap.ErrBadResponse
	}

	return insertVarBatchResponse, err
}

func (s *Service) NewWageComponent() *EmployeeWageComponent {
	return &EmployeeWageComponent{}
}

func (s *Service) NewInsertVarBatchRequest() *InsertVarBatchRequest {
	return &InsertVarBatchRequest{}
}

func newInsertVarBatchAction(requestBody *InsertVarBatchRequest, responseBody *insertVarBatchResponse) (*soap.Request, *soap.Response) {
	requestBody.XMLName = xml.Name{
		Space: xmlns,
		Local: insertVarBatchAction,
	}
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

type EmployeeWageComponent struct {
	EmployeeID   int     `xml:"EmployeeWageComponent>EmployeeId"`
	ID           int     `xml:"EmployeeWageComponent>Id,omitempty"`
	Code         int     `xml:"EmployeeWageComponent>Code,omitempty"`
	Value        float64 `xml:"EmployeeWageComponent>Value,omitempty"`
	CostCenterID int     `xml:"EmployeeWageComponent>CostCenterId,omitempty"`
	CostUnitID   int     `xml:"EmployeeWageComponent>CostUnitId,omitempty"`
	Comment      string  `xml:"EmployeeWageComponent>Comment"`
	DocumentName string  `xml:"EmployeeWageComponent>DocumentName"`
	// Attachment base64Binary `xml:"EmployeeWageComponent>Hourcomponent>Attachment"`
}

type InsertVarBatchRequest struct {
	XMLName xml.Name

	WageComponents  []EmployeeWageComponent `xml:"WageComponents"`
	Period          int                     `xml:"Period"`
	Year            int                     `xml:"Year"`
	UnprotectedMode bool                    `xml:"UnprotectedMode"`
}

type insertVarBatchResponse struct {
}
