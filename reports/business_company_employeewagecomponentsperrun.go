package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	businessCompanyEmployeeWageComponentsPerRunAction = "Reports_Business_Company_EmployeeWageComponentsPerRun"
)

// Generate company journal report within specified period
func (s *Service) BusinessCompanyEmployeeWageComponentsPerRun(companyID int, year int) (*businessCompanyEmployeeWageComponentsPerRunResponse, error) {
	request, response := newBusinessCompanyEmployeeWageComponentsPerRun(companyID, year)

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
	reportResponse, ok := response.Envelope.Body.Data.(*businessCompanyEmployeeWageComponentsPerRunResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBusinessCompanyEmployeeWageComponentsPerRun(companyID int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBusinessCompanyEmployeeWageComponentsPerRunRequest(companyID, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBusinessCompanyEmployeeWageComponentsPerRunResponse()

	return request, response
}

type businessCompanyEmployeeWageComponentsPerRunRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
	Year      int `xml:"Year"`
}

func newBusinessCompanyEmployeeWageComponentsPerRunRequest(companyID int, year int) *businessCompanyEmployeeWageComponentsPerRunRequest {
	return &businessCompanyEmployeeWageComponentsPerRunRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: businessCompanyEmployeeWageComponentsPerRunAction,
		},

		CompanyID: companyID,
		Year:      year,
	}
}

type businessCompanyEmployeeWageComponentsPerRunResponse struct {
	Report BusinessCompanyEmployeeWageComponentsPerRunReport `xml:"Reports_Business_Company_EmployeeWageComponentsPerRunResult"`
}

func newBusinessCompanyEmployeeWageComponentsPerRunResponse() *businessCompanyEmployeeWageComponentsPerRunResponse {
	return &businessCompanyEmployeeWageComponentsPerRunResponse{}
}

type BusinessCompanyEmployeeWageComponentsPerRunReport struct {
	// JournalReportItems BusinessCompanyEmployeeWageComponentsPerRunItems `xml:"JournalReportItems>JournalReportItem"`
}

// func (r *BusinessCompanyEmployeeWageComponentsPerRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	v := []byte{}
// 	err := d.DecodeElement(&v, &start)
// 	// err := d.Decode(&v)
// 	if err != nil {
// 		return err
// 	}

// 	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
// 	v = []byte(s)

// 	// Create alias with UnmarshalXML method
// 	type Alias BusinessCompanyEmployeeWageComponentsPerRun
// 	a := (*Alias)(r)

// 	// Unmarshal alias with cdata xml
// 	err = xml.Unmarshal(v, a)
// 	if err != nil {
// 		return err
// 	}

// 	// Copy alias back to original
// 	*r = (BusinessCompanyEmployeeWageComponentsPerRun)(*a)
// 	return nil
// }

// type BusinessCompanyEmployeeWageComponentsPerRunItems []BusinessCompanyEmployeeWageComponentsPerRunItem

// type BusinessCompanyEmployeeWageComponentsPerRunItem struct {
// 	EmployeeID         int
// 	EmployeeNumber     string
// 	EmployeeName       string
// 	CompanyID          int
// 	CompanyNumber      string
// 	CompanyName        string
// 	DebtorID           int
// 	DebtorNumber       string
// 	DebtorName         string
// 	Initials           string
// 	Firstname          string
// 	FirstnameInFull    string
// 	Surname            string
// 	ParttimePercentage float64
// 	FTE                float64
// 	TotalHoursWeek     float64
// 	TotalHoursPeriod   float64
// 	SalaryType         string
// 	SalaryValue        float64
// 	Function           string
// 	Department         string
// }
