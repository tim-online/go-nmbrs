package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	AccountantEmployeeNameAddressBackgroundAction = "Reports_Accountant_EmployeeNameAddress_Background"
)

// Generate company journal report within specified period
func (s *Service) AccountantEmployeeNameAddressBackground() (*AccountantEmployeeNameAddressBackgroundResponse, error) {
	request, response := newAccountantEmployeeNameAddressBackground()

	request.Envelope.Header.Data = s.AuthHeader

	httpReq, err := s.Client.NewRequest(s.Endpoint.String(), request)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, response)
	if err != nil {
		return nil, err
	}

	// @TODO: check if this can be better
	reportResponse, ok := response.Envelope.Body.Data.(*AccountantEmployeeNameAddressBackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newAccountantEmployeeNameAddressBackground() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newAccountantEmployeeNameAddressBackgroundRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newAccountantEmployeeNameAddressBackgroundResponse()

	return request, response
}

type AccountantEmployeeNameAddressBackgroundRequest struct {
	XMLName xml.Name
}

func newAccountantEmployeeNameAddressBackgroundRequest() *AccountantEmployeeNameAddressBackgroundRequest {
	return &AccountantEmployeeNameAddressBackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: AccountantEmployeeNameAddressBackgroundAction,
		},
	}
}

type AccountantEmployeeNameAddressBackgroundResponse struct {
	ReportsAccountantEmployeeNameAddressBackgroundResult string `xml:"Reports_Accountant_EmployeeNameAddress_BackgroundResult"`
}

func newAccountantEmployeeNameAddressBackgroundResponse() *AccountantEmployeeNameAddressBackgroundResponse {
	return &AccountantEmployeeNameAddressBackgroundResponse{}
}

type AccountantEmployeeNameAddresses []AccountantEmployeeNameAddress

// func (r *AccountantEmployeeNameAddresses) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	v := []byte{}
// 	err := d.DecodeElement(&v, &start)
// 	// err := d.Decode(&v)
// 	if err != nil {
// 		return err
// 	}

// 	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
// 	v = []byte(s)

// 	// Create alias with UnmarshalXML method
// 	type Alias AccountantEmployeeNameAddresses
// 	a := (*Alias)(r)

// 	// Unmarshal alias with cdata xml
// 	err = xml.Unmarshal(v, a)
// 	if err != nil {
// 		return err
// 	}

// 	// Copy alias back to original
// 	*r = (AccountantEmployeeNameAddresses)(*a)
// 	return nil
// }

type AccountantEmployeeNameAddress struct {
	EmployeeId             string `xml:"EmployeeId"`
	EmployeeNumber         string `xml:"EmployeeNumber"`
	EmployeeName           string `xml:"EmployeeName"`
	Initials               string `xml:"Initials"`
	Surname                string `xml:"Surname"`
	Firstname              string `xml:"Firstname"`
	FirstNameInFull        string `xml:"FirstNameInFull"`
	BSN                    string `xml:"BSN"`
	Birthdate              string `xml:"Birthdate"`
	PlaceOfBirth           string `xml:"PlaceOfBirth"`
	TelephonePrivate       string `xml:"TelephonePrivate"`
	TelephoneMobilePrivate string `xml:"TelephoneMobilePrivate"`
	EmailPrivate           string `xml:"EmailPrivate"`
	Gender                 string `xml:"Gender"`
	Title                  string `xml:"Title"`
	Nationality            string `xml:"Nationality"`
	State                  string `xml:"State"`
	CompanyId              string `xml:"CompanyId"`
	Company                string `xml:"Company"`
	CompanyNumber          string `xml:"CompanyNumber"`
	Debtor                 string `xml:"Debtor"`
	DebtorId               string `xml:"DebtorId"`
	DebtorNumber           string `xml:"DebtorNumber"`
	InServiceDate          string `xml:"InServiceDate"`
	OutServiceDate         string `xml:"OutServiceDate"`
	ContractStartDate      string `xml:"ContractStartDate"`
	ContractEndDate        string `xml:"ContractEndDate"`
	Address                string `xml:"Address"`
	HouseNumber            string `xml:"HouseNumber"`
	PostCode               string `xml:"PostCode"`
	City                   string `xml:"City"`
	Country                string `xml:"Country"`
	CountryOfBirth         string `xml:"CountryOfBirth"`
	EmployeeType           string `xml:"EmployeeType"`
	Prefix                 string `xml:"Prefix"`
	TelephoneOther         string `xml:"TelephoneOther"`
	TelephoneMobileWork    string `xml:"TelephoneMobileWork"`
	EmailWork              string `xml:"EmailWork"`
	TelephoneWork          string `xml:"TelephoneWork"`
}
