package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Salary_GetAll_AllEmployeesByCompanyAction = "Salary_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Salary_GetAll_AllEmployeesByCompany(companyID int) (*Salary_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewSalary_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Salary_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewSalary_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewSalary_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewSalary_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type Salary_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewSalary_GetAll_AllEmployeesByCompanyRequest(companyID int) *Salary_GetAll_AllEmployeesByCompanyRequest {
	return &Salary_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Salary_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type Salary_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeeSalaryItems EmployeeSalaryItems `xml:"Salary_GetAll_AllEmployeesByCompanyResult>EmployeeSalaryItem"`
}

func NewSalary_GetAll_AllEmployeesByCompanyResponse() *Salary_GetAll_AllEmployeesByCompanyResponse {
	return &Salary_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeeSalaryItems []EmployeeSalaryItem

type EmployeeSalaryItem struct {
	EmployeeID       int `xml:"EmployeeId"`
	EmployeeSalaries []struct {
		ID           string   `xml:"ID"`
		Value        string   `xml:"Value"`
		Type         string   `xml:"Type"`
		StartDate    lib.Time `xml:"StartDate"`
		CreationDate lib.Time `xml:"CreationDate"`
		SalaryTable  struct {
			Code        string `xml:"Code"`
			Description string `xml:"Description"`
			Schaal      struct {
				Scale              string `xml:"Scale"`
				SchaalDescription  string `xml:"SchaalDescription"`
				ScaleValue         string `xml:"ScaleValue"`
				ScalePercentageMax string `xml:"ScalePercentageMax"`
				ScalePercentageMin string `xml:"ScalePercentageMin"`
			} `xml:"Schaal"`
			Trede struct {
				Step            string `xml:"Step"`
				StepDescription string `xml:"StepDescription"`
				StepValue       string `xml:"StepValue"`
			} `xml:"Trede"`
		} `xml:"SalaryTable"`
	} `xml:"EmployeeSalaries>Salary_V2"`
}
