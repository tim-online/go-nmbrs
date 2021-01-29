package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Contract_GetAll_AllEmployeesByCompanyAction = "Contract_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Contract_GetAll_AllEmployeesByCompany(companyID int) (*Contract_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewContract_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Contract_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewContract_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewContract_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewContract_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type Contract_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewContract_GetAll_AllEmployeesByCompanyRequest(companyID int) *Contract_GetAll_AllEmployeesByCompanyRequest {
	return &Contract_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Contract_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type Contract_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeeContractItems EmployeeContractItems `xml:"Contract_GetAll_AllEmployeesByCompanyResult>EmployeeContractItem"`
}

func NewContract_GetAll_AllEmployeesByCompanyResponse() *Contract_GetAll_AllEmployeesByCompanyResponse {
	return &Contract_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeeContractItems []EmployeeContractItem

type EmployeeContractItem struct {
	EmployeeID int `xml:"EmployeeId"`
	Contracts  []struct {
		ContractID              string      `xml:"ContractID"`
		CreationDate            lib.Time    `xml:"CreationDate"`
		StartDate               lib.Time    `xml:"StartDate"`
		TrialPeriod             interface{} `xml:"TrialPeriod"`
		EndDate                 lib.Time    `xml:"EndDate"`
		EmployementType         string      `xml:"EmployementType"`
		EmploymentSequenceTaxId string      `xml:"EmploymentSequenceTaxId"`
		Indefinite              string      `xml:"Indefinite"`
	} `xml:"EmployeeContracts>EmployeeContract"`
}
