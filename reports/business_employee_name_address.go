package reports

import (
	"encoding/xml"
	"strings"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	businessEmployeeNameAddressAction = "Reports_Business_EmployeeNameAddress"
)

// Generate company journal report within specified period
func (s *Service) BusinessEmployeeNameAddress() (*businessEmployeeNameAddressResponse, error) {
	request, response := newBusinessEmployeeNameAddress()

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
	reportResponse, ok := response.Envelope.Body.Data.(*businessEmployeeNameAddressResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBusinessEmployeeNameAddress() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBusinessEmployeeNameAddressRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBusinessEmployeeNameAddressResponse()

	return request, response
}

type businessEmployeeNameAddressRequest struct {
	XMLName xml.Name
}

func newBusinessEmployeeNameAddressRequest() *businessEmployeeNameAddressRequest {
	return &businessEmployeeNameAddressRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: businessEmployeeNameAddressAction,
		},
	}
}

type businessEmployeeNameAddressResponse struct {
	Report BusinessEmployeeNameAddressReport `xml:"Reports_Business_EmployeeNameAddressResult"`
}

func newBusinessEmployeeNameAddressResponse() *businessEmployeeNameAddressResponse {
	return &businessEmployeeNameAddressResponse{}
}

type BusinessEmployeeNameAddressReport struct {
	EmployeeNameAddresses BusinessEmployeeNameAddresses `xml:"EmployeeList>Employee"`
}

func (r *BusinessEmployeeNameAddressReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v := []byte{}
	err := d.DecodeElement(&v, &start)
	// err := d.Decode(&v)
	if err != nil {
		return err
	}

	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
	v = []byte(s)

	// Create alias with UnmarshalXML method
	type Alias BusinessEmployeeNameAddressReport
	a := (*Alias)(r)

	// Unmarshal alias with cdata xml
	err = xml.Unmarshal(v, a)
	if err != nil {
		return err
	}

	// Copy alias back to original
	*r = (BusinessEmployeeNameAddressReport)(*a)
	return nil
}

type BusinessEmployeeNameAddresses []BusinessEmployeeNameAddress

type BusinessEmployeeNameAddress struct {
	EmployeeID             int `xml:"EmployeeId"`
	EmployeeNumber         int
	EmployeeName           string
	Initials               string
	Surname                string
	FirstName              string
	FirstNameInFull        string
	BSN                    int
	Birthdate              lib.Time
	PlaceOfBirth           string
	TelephonePrivate       string
	TelephoneMobilePrivate string
	EmailPrivate           string
	Gender                 string
	Title                  string
	Nationality            string
	State                  string
	CompanyID              int `xml:"CompanyId"`
	Company                string
	CompanyNumber          string
	Debtor                 string
	DebtorID               int `xml:"DebtorId"`
	DebtorNumber           int
	InServiceDate          lib.Time
	OutServiceDate         lib.Time
	ContractStartDate      lib.Time
	ContractEndDate        lib.Time
	Address                string
	HouseNumber            string
	PostCode               string
	City                   string
	Country                string
	CountryOfBirth         string
}
