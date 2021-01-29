package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Function_GetAll_AllEmployeesByCompanyAction = "Function_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Function_GetAll_AllEmployeesByCompany(companyID int) (*Function_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewFunction_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Function_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewFunction_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewFunction_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewFunction_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type Function_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewFunction_GetAll_AllEmployeesByCompanyRequest(companyID int) *Function_GetAll_AllEmployeesByCompanyRequest {
	return &Function_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Function_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type Function_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeeFunctionItems EmployeeFunctionItems `xml:"Function_GetAll_AllEmployeesByCompanyResult>EmployeeFunctionItem"`
}

func NewFunction_GetAll_AllEmployeesByCompanyResponse() *Function_GetAll_AllEmployeesByCompanyResponse {
	return &Function_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeeFunctionItems []EmployeeFunctionItem

type EmployeeFunctionItem struct {
	EmployeeID        int `xml:"EmployeeId"`
	EmployeeFunctions []struct {
		Text         string   `xml:",chardata"`
		ID           string   `xml:"Id"`
		Code         string   `xml:"Code"`
		Description  string   `xml:"Description"`
		CreationDate lib.Time `xml:"CreationDate"`
		StartPeriod  string   `xml:"StartPeriod"`
		StartYear    string   `xml:"StartYear"`
	} `xml:"EmployeeFunctions>Function_V2"`
}
