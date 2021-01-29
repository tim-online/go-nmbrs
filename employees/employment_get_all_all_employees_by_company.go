package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Employment_GetAll_AllEmployeesByCompanyAction = "Employment_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Employment_GetAll_AllEmployeesByCompany(companyID int) (*Employment_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewEmployment_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Employment_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewEmployment_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewEmployment_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewEmployment_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type Employment_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewEmployment_GetAll_AllEmployeesByCompanyRequest(companyID int) *Employment_GetAll_AllEmployeesByCompanyRequest {
	return &Employment_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Employment_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type Employment_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeeEmploymentItems EmployeeEmploymentItems `xml:"Employment_GetAll_AllEmployeesByCompanyResult>EmployeeEmploymentItem"`
}

func NewEmployment_GetAll_AllEmployeesByCompanyResponse() *Employment_GetAll_AllEmployeesByCompanyResponse {
	return &Employment_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeeEmploymentItems []EmployeeEmploymentItem

type EmployeeEmploymentItem struct {
	EmployeeID  int `xml:"EmployeeId"`
	Employments []struct {
		EmploymentID            string      `xml:"EmploymentID"`
		CreationDate            lib.Time    `xml:"CreationDate"`
		StartDate               lib.Time    `xml:"StartDate"`
		TrialPeriod             interface{} `xml:"TrialPeriod"`
		EndDate                 lib.Time    `xml:"EndDate"`
		EmployementType         string      `xml:"EmployementType"`
		EmploymentSequenceTaxId string      `xml:"EmploymentSequenceTaxId"`
		Indefinite              string      `xml:"Indefinite"`
	} `xml:"EmployeeEmployments>Employment"`
}
