package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Department_GetAll_AllEmployeesByCompanyAction = "Department_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Department_GetAll_AllEmployeesByCompany(companyID int) (*Department_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewDepartment_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Department_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewDepartment_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewDepartment_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewDepartment_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type Department_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewDepartment_GetAll_AllEmployeesByCompanyRequest(companyID int) *Department_GetAll_AllEmployeesByCompanyRequest {
	return &Department_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Department_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type Department_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeeDepartmentItems EmployeeDepartmentItems `xml:"Department_GetAll_AllEmployeesByCompanyResult>EmployeeDepartmentItem"`
}

func NewDepartment_GetAll_AllEmployeesByCompanyResponse() *Department_GetAll_AllEmployeesByCompanyResponse {
	return &Department_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeeDepartmentItems []EmployeeDepartmentItem

type EmployeeDepartmentItem struct {
	EmployeeID  int `xml:"EmployeeId"`
	Departments []struct {
		ID           int      `xml:"Id"`
		Code         int      `xml:"Code"`
		Description  string   `xml:"Description"`
		CreationDate lib.Time `xml:"CreationDate"`
		StartPeriod  int      `xml:"StartPeriod"`
		StartYear    int      `xml:"StartYear"`
	} `xml:"EmployeeDepartments>Department_V2"`
}
