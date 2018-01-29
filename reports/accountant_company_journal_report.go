package reports

import (
	"encoding/xml"
	"strings"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	accountantCompanyJournalReportAction = "Reports_Accountant_Company_CompanyJournalReport"
)

// Generate company journal report within specified period
func (s *Service) AccountantCompanyJournalReport(companyID int, startYear int, startPeriod int, endYear int, endPeriod int) (*accountantCompanyJournalReportResponse, error) {
	request, response := newAccountantCompanyJournalReport(companyID, startYear, startPeriod, endYear, endPeriod)

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
	reportResponse, ok := response.Envelope.Body.Data.(*accountantCompanyJournalReportResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newAccountantCompanyJournalReport(companyID int, startYear int, startPeriod int, endYear int, endPeriod int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newAccountantCompanyJournalReportRequest(companyID, startYear, startPeriod, endYear, endPeriod)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newAccountantCompanyJournalReportResponse()

	return request, response
}

type accountantCompanyJournalReportRequest struct {
	XMLName xml.Name

	CompanyID   int `xml:"CompanyId"`
	FromYear    int `xml:"FromYear"`
	FromPeriod  int `xml:"FromPeriod"`
	UntilYear   int `xml:"UntilYear"`
	UntilPeriod int `xml:"UntilPeriod"`
}

func newAccountantCompanyJournalReportRequest(companyID int, startYear int, startPeriod int, endYear int, endPeriod int) *accountantCompanyJournalReportRequest {
	return &accountantCompanyJournalReportRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: accountantCompanyJournalReportAction,
		},

		CompanyID:   companyID,
		FromYear:    startYear,
		FromPeriod:  startPeriod,
		UntilYear:   endYear,
		UntilPeriod: endPeriod,
	}
}

type accountantCompanyJournalReportResponse struct {
	CompanyJournalReport AccountantCompanyJournalReport `xml:"Reports_Accountant_Company_CompanyJournalReportResult"`
}

func newAccountantCompanyJournalReportResponse() *accountantCompanyJournalReportResponse {
	return &accountantCompanyJournalReportResponse{}
}

type AccountantCompanyJournalReport struct {
	JournalReportItems AccountantCompanyJournalReportItems `xml:"JournalReportItems>JournalReportItem"`
}

func (r *AccountantCompanyJournalReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v := []byte{}
	err := d.DecodeElement(&v, &start)
	// err := d.Decode(&v)
	if err != nil {
		return err
	}

	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
	v = []byte(s)

	// Create alias with UnmarshalXML method
	type Alias AccountantCompanyJournalReport
	a := (*Alias)(r)

	// Unmarshal alias with cdata xml
	err = xml.Unmarshal(v, a)
	if err != nil {
		return err
	}

	// Copy alias back to original
	*r = (AccountantCompanyJournalReport)(*a)
	return nil
}

type AccountantCompanyJournalReportItems []AccountantCompanyJournalReportItem

// func (r *AccountantCompanyJournalReportItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	v := []byte{}
// 	err := d.DecodeElement(&v, &start)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("---------------------")
// 	log.Printf("%+v", start.Name)
// 	log.Println(string(v))
// 	log.Println("---------------------")
// 	os.Exit(43)
// 	return nil
// }

type AccountantCompanyJournalReportItem struct {
	CompanyID            int
	CompanyName          string
	CompanyNr            string
	EmployeeID           int
	EmployeeName         string
	EmployeeNr           string
	Period               int
	Year                 int
	Run                  string
	Department           string
	CostCenter           string
	CostUnit             string
	Percentage           string
	GeneralLedgerAccount string
	ComponentName        string
	Debit                float64
	Credit               float64
}
