package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listByCompanyAction = "https://api.nmbrs.nl/soap/v2.1/EmployeeService/List_GetByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) ListByCompany(companyID int, active ActiveFilter) (*listByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := newListByCompanyAction(companyID, active)

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
	listResponse, ok := response.Envelope.Body.Data.(*listByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func newListByCompanyAction(companyID int, active ActiveFilter) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	request.Action = url.MustParse(listByCompanyAction)
	request.Envelope.Body.Data = newListByCompanyRequest(companyID, active)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListByCompanyResponse()

	return request, response
}

type listByCompanyRequest struct {
	XMLName xml.Name `xml:"List_GetByCompany"`
	Xmlns   string   `xml:"xmlns,attr"`

	CompanyID int          `xml:"CompanyId"`
	Active    ActiveFilter `xml:"active"`
}

func newListByCompanyRequest(companyID int, active ActiveFilter) *listByCompanyRequest {
	return &listByCompanyRequest{
		Xmlns:     xmlns,
		CompanyID: companyID,
		Active:    active,
	}
}

type listByCompanyResponse struct {
	Employees []Employee `xml:"List_GetByCompanyResult>Employee"`
}

func newListByCompanyResponse() *listByCompanyResponse {
	return &listByCompanyResponse{}
}
