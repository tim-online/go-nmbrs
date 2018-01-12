package companies

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	journalReportsByCompanyAction = "Reports_GetJournalsReportByCompany"
)

// Returns the Journal Report for Company
func (s *Service) JournalReportsByCompany(companyID int, startPeriod int, endPeriod int, year int) (*journalReportsByCompanyResponse, error) {
	request, response := newJournalReportsByCompanyActions(companyID, startPeriod, endPeriod, year)
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
	journalReportsByCompanyResponse, ok := response.Envelope.Body.Data.(*journalReportsByCompanyResponse)
	if ok == false {
		return journalReportsByCompanyResponse, soap.ErrBadResponse
	}

	return journalReportsByCompanyResponse, err
}

func newJournalReportsByCompanyActions(companyID int, startPeriod int, endPeriod int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newJournalReportsByCompanyRequest(companyID, startPeriod, endPeriod, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newJournalsReportByCompanyResponse()

	return request, response
}

type journalReportsByCompanyRequest struct {
	XMLName xml.Name

	CompanyID   int `xml:"CompanyId"`
	StartPeriod int `xml:"StartPeriod"`
	EndPeriod   int `xml:"EndPeriod"`
	Year        int `xml:"Year"`
}

func newJournalReportsByCompanyRequest(companyID int, startPeriod int, endPeriod int, year int) *journalReportsByCompanyRequest {
	return &journalReportsByCompanyRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: journalReportsByCompanyAction,
		},
		CompanyID:   companyID,
		StartPeriod: startPeriod,
		EndPeriod:   endPeriod,
		Year:        year,
	}
}

type journalReportsByCompanyResponse struct {
	JournalsReport JournalsReport `xml:"Reports_GetJournalsReportByCompanyResult"`
}

func newJournalsReportByCompanyResponse() *journalReportsByCompanyResponse {
	return &journalReportsByCompanyResponse{}
}

type JournalsReport struct {
	CompanyId   int                 `xml:"CompanyId"`
	StartPeriod int                 `xml:"StartPeriod"`
	EndPeriod   int                 `xml:"EndPeriod"`
	Year        int                 `xml:"Year"`
	Items       JournalsReportItems `xml:"Items>Item"`
}

// type JournalsReport []byte

func (r *JournalsReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	data := []byte{}
	err := d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}

	type alias JournalsReport
	a := alias{}
	err = xml.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	r2 := JournalsReport(a)
	*r = r2
	return nil
}

type JournalsReportItems []JournalsReportItem

type JournalsReportItem struct {
	EmployeeName         string `xml:"EmployeeName"`
	Period               int    `xml:"Period"`
	Year                 int    `xml:"Year"`
	Run                  string `xml:"Run"`
	Department           string `xml:"Department"`
	CostCenter           string `xml:"CostCenter"`
	CostUnit             string `xml:"CostUnit"`
	Percentage           string `xml:"Percentage"`
	GeneralLedgerAccount string `xml:"GeneralLedgerAccount"`
	ComponentName        string `xml:"ComponentName"`
	Debit                string `xml:"Debit"`
	Credit               string `xml:"Credit"`
}
