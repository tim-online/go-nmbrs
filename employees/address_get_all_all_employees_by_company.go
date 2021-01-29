package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	Address_GetAll_AllEmployeesByCompanyAction = "Address_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) Address_GetAll_AllEmployeesByCompany(companyID int) (*Address_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewAddress_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*Address_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewAddress_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewAddress_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewAddress_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type Address_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewAddress_GetAll_AllEmployeesByCompanyRequest(companyID int) *Address_GetAll_AllEmployeesByCompanyRequest {
	return &Address_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: Address_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type Address_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeeAddressItems EmployeeAddressItems `xml:"Address_GetAll_AllEmployeesByCompanyResult>EmployeeAddressItem"`
}

func NewAddress_GetAll_AllEmployeesByCompanyResponse() *Address_GetAll_AllEmployeesByCompanyResponse {
	return &Address_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeeAddressItems []EmployeeAddressItem

type EmployeeAddressItem struct {
	EmployeeID int `xml:"EmployeeId"`
	Addresses  []struct {
		ID                  string   `xml:"Id"`
		StartPeriod         string   `xml:"StartPeriod"`
		StartYear           string   `xml:"StartYear"`
		EndPeriod           int      `xml:"EndPeriod"`
		EndYear             int      `xml:"EndYear"`
		Street              string   `xml:"Street"`
		HouseNumber         string   `xml:"HouseNumber"`
		PostalCode          string   `xml:"PostalCode"`
		City                string   `xml:"City"`
		StateProvince       string   `xml:"StateProvince"`
		CreationDate        lib.Time `xml:"CreationDate"`
		CountryISOCode      string   `xml:"CountryISOCode"`
		Type                string   `xml:"Type"`
		HouseNumberAddition string   `xml:"HouseNumberAddition"`
	} `xml:"EmployeeAddresses>EmployeeAddress_V2"`
}
