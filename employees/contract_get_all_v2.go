package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v3/EmployeeService.asmx?op=Contract_GetAll_v2

const (
	Contract_GetAll_v2Action = "Contract_GetAll_v2"
)

// Contract_GetAll_v2 gets all contracts for the specified employee.
func (s *Service) Contract_GetAll_v2(employeeID int) (*Contract_GetAll_v2Response, error) {
	// get a new request & response envelope
	request, response := s.NewContract_GetAll_v2Action(employeeID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Contract_GetAll_v2Response)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewContract_GetAll_v2Action(employeeID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewContract_GetAll_v2Request(employeeID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewContract_GetAll_v2Response()

	return request, response
}

type Contract_GetAll_v2Request struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
}

func NewContract_GetAll_v2Request(employeeID int) *Contract_GetAll_v2Request {
	return &Contract_GetAll_v2Request{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Contract_GetAll_v2Action,
		},

		EmployeeID: employeeID,
	}
}

type Contract_GetAll_v2Response struct {
	Contracts []Contract_v2 `xml:"Contract_GetAll_v2Result>Contract_v2"`
}

func NewContract_GetAll_v2Response() *Contract_GetAll_v2Response {
	return &Contract_GetAll_v2Response{}
}

type Contract_v2 struct {
	ContractID              int       `xml:"ContractID"`
	StartDate               lib.Time  `xml:"StartDate"`
	TrialPeriod             *lib.Time `xml:"TrialPeriod"`
	EndDate                 *lib.Time `xml:"EndDate"`
	EmploymentType          int       `xml:"EmploymentType"`
	ContractType            int       `xml:"ContractType"`
	EmploymentSequenceTaxID int       `xml:"EmploymentSequenceTaxId"`
	Indefinite              bool      `xml:"Indefinite"`
	PhaseClassification     int       `xml:"PhaseClassification"`
}
