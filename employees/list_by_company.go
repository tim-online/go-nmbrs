package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listByCompanyAction = "List_GetByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) ListByCompany(companyID int, employeeType int) (*listByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := newListByCompanyAction(companyID, employeeType)

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

func newListByCompanyAction(companyID int, employeeType int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newListByCompanyRequest(companyID, employeeType)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListByCompanyResponse()

	return request, response
}

type listByCompanyRequest struct {
	XMLName xml.Name

	CompanyID    int `xml:"CompanyId"`
	EmployeeType int `xml:"EmployeeType"`
}

func newListByCompanyRequest(companyID int, employeeType int) *listByCompanyRequest {
	return &listByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listByCompanyAction,
		},

		CompanyID:    companyID,
		EmployeeType: employeeType,
	}
}

type listByCompanyResponse struct {
	Employees []Employee `xml:"List_GetByCompanyResult>Employee"`
}

func newListByCompanyResponse() *listByCompanyResponse {
	return &listByCompanyResponse{}
}
