package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	PersonalInfo_GetAll_AllEmployeesByCompanyAction = "PersonalInfo_GetAll_AllEmployeesByCompany"
)

// ListByCompany gets all employees that belong to a company and are active as given.
func (s *Service) PersonalInfo_GetAll_AllEmployeesByCompany(companyID int) (*PersonalInfo_GetAll_AllEmployeesByCompanyResponse, error) {
	// get a new request & response envelope
	request, response := s.NewPersonalInfo_GetAll_AllEmployeesByCompanyAction(companyID)

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
	listResponse, ok := response.Envelope.Body.Data.(*PersonalInfo_GetAll_AllEmployeesByCompanyResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewPersonalInfo_GetAll_AllEmployeesByCompanyAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewPersonalInfo_GetAll_AllEmployeesByCompanyRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewPersonalInfo_GetAll_AllEmployeesByCompanyResponse()

	return request, response
}

type PersonalInfo_GetAll_AllEmployeesByCompanyRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyID"`
}

func NewPersonalInfo_GetAll_AllEmployeesByCompanyRequest(companyID int) *PersonalInfo_GetAll_AllEmployeesByCompanyRequest {
	return &PersonalInfo_GetAll_AllEmployeesByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: PersonalInfo_GetAll_AllEmployeesByCompanyAction,
		},

		CompanyID: companyID,
	}
}

type PersonalInfo_GetAll_AllEmployeesByCompanyResponse struct {
	EmployeePersonalInfoItems EmployeePersonalInfoItems `xml:"PersonalInfo_GetAll_AllEmployeesByCompanyResult>PersonalInfoItem"`
}

func NewPersonalInfo_GetAll_AllEmployeesByCompanyResponse() *PersonalInfo_GetAll_AllEmployeesByCompanyResponse {
	return &PersonalInfo_GetAll_AllEmployeesByCompanyResponse{}
}

type EmployeePersonalInfoItems []EmployeePersonalInfoItem

type EmployeePersonalInfoItem struct {
	EmployeeID    int `xml:"EmployeeId"`
	PersonalInfos []struct {
		Text                   string   `xml:",chardata"`
		ID                     string   `xml:"Id"`
		DisplayName            string   `xml:"DisplayName"`
		EmployeeNumber         string   `xml:"EmployeeNumber"`
		BSN                    string   `xml:"BSN"`
		Initials               string   `xml:"Initials"`
		Prefix                 string   `xml:"Prefix"`
		LastName               string   `xml:"LastName"`
		Nickname               string   `xml:"Nickname"`
		Gender                 string   `xml:"Gender"`
		NationalityCode        string   `xml:"NationalityCode"`
		PlaceOfBirth           string   `xml:"PlaceOfBirth"`
		CountryOfBirthISOCode  string   `xml:"CountryOfBirthISOCode"`
		IdentificationType     string   `xml:"IdentificationType"`
		TelephoneWork          string   `xml:"TelephoneWork"`
		EmailWork              string   `xml:"EmailWork"`
		BurgerlijkeStaat       string   `xml:"BurgerlijkeStaat"`
		Naamstelling           string   `xml:"Naamstelling"`
		Birthday               lib.Time `xml:"Birthday"`
		CreationDate           lib.Time `xml:"CreationDate"`
		StartPeriod            string   `xml:"StartPeriod"`
		StartYear              string   `xml:"StartYear"`
		FirstName              string   `xml:"FirstName"`
		TelephoneMobileWork    string   `xml:"TelephoneMobileWork"`
		PartnerPrefix          string   `xml:"PartnerPrefix"`
		PartnerLastName        string   `xml:"PartnerLastName"`
		TelephonePrivate       string   `xml:"TelephonePrivate"`
		InCaseOfEmergency      string   `xml:"InCaseOfEmergency"`
		TelephoneMobilePrivate string   `xml:"TelephoneMobilePrivate"`
		EmailPrivate           string   `xml:"EmailPrivate"`
		Title                  string   `xml:"Title"`
	} `xml:"EmployeePersonalInfos>PersonalInfo_V2"`
}
