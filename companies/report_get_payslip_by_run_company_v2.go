package companies

import (
	"encoding/xml"
	"strings"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	reportGetPayslipByRunCompanyV2Action = "Reports_GetPayslipByRunCompany_v2"
)

// Returns the Journal Report for Company
func (s *Service) ReportGetPayslipByRunCompanyV2(companyID int, runID int, year int) (*reportGetPayslipByRunCompanyV2Response, error) {
	request, response := newReportGetPayslipByRunCompanyV2Actions(companyID, runID, year)
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
	reportGetPayslipByRunCompanyV2Response, ok := response.Envelope.Body.Data.(*reportGetPayslipByRunCompanyV2Response)
	if ok == false {
		return reportGetPayslipByRunCompanyV2Response, soap.ErrBadResponse
	}

	return reportGetPayslipByRunCompanyV2Response, err
}

func newReportGetPayslipByRunCompanyV2Actions(companyID int, runID int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newReportGetPayslipByRunCompanyV2Request(companyID, runID, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newReportGetPayslipByRunCompanyV2Response()

	return request, response
}

type reportGetPayslipByRunCompanyV2Request struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
	RunID     int `xml:"RunId"`
	Year      int `xml:"intYear"`
}

func newReportGetPayslipByRunCompanyV2Request(companyID int, runID int, year int) *ageCodesByRunCompanyV2Request {
	return &ageCodesByRunCompanyV2Request{
		XMLName: xml.Name{
			Space: xmlns,
			Local: reportGetPayslipByRunCompanyV2Action,
		},
		CompanyID: companyID,
		RunID:     runID,
		Year:      year,
	}
}

type reportGetPayslipByRunCompanyV2Response struct {
	Report ReportGetPayslipByRunCompanyV2Report `xml:"Reports_GetPayslipByRunCompany_v2Result"`
}

func newReportGetPayslipByRunCompanyV2Response() *reportGetPayslipByRunCompanyV2Response {
	return &reportGetPayslipByRunCompanyV2Response{}
}

type ReportGetPayslipByRunCompanyV2Report struct {
	Payslips Payslips `xml:"payslip"`
}

type Payslips []Payslip

type Payslip struct {
	Text         string `xml:",chardata"`
	EmployeeID   int    `xml:"employeeid"`
	Employeename string `xml:"employeename"`
	Period       int    `xml:"period"`
	Year         int    `xml:"year"`
	Run          int    `xml:"run"`
	Header       struct {
		Text       string `xml:",chardata"`
		HeaderItem []struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"headerItem"`
	} `xml:"header"`
	Lines struct {
		Text string `xml:",chardata"`
		Line []struct {
			Text        string `xml:",chardata"`
			Code        string `xml:"code"`
			Description string `xml:"description"`
			Group       string `xml:"group"`
			Amount      string `xml:"amount"`
			Value       string `xml:"value"`
			Betaling    string `xml:"betaling"`
			Tabel       string `xml:"tabel"`
			Bt          string `xml:"bt"`
			Svw         string `xml:"svw"`
			Svwbt       string `xml:"svwbt"`
			Werkgever   string `xml:"werkgever"`
			Cumulative  string `xml:"cumulative"`
		} `xml:"line"`
	} `xml:"lines"`
	Reservations struct {
		Text        string `xml:",chardata"`
		Reservation []struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
			Saldo string `xml:"saldo"`
		} `xml:"reservation"`
	} `xml:"reservations"`
	LeaveList struct {
		Text  string `xml:",chardata"`
		Leave []struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Key   string `xml:"key"`
			Value []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"value"`
		} `xml:"leave"`
	} `xml:"leave_list"`
	Footer struct {
		Text       string `xml:",chardata"`
		FooterItem []struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"footerItem"`
	} `xml:"footer"`
}

func (r *ReportGetPayslipByRunCompanyV2Report) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v := []byte{}
	err := d.DecodeElement(&v, &start)
	// err := d.Decode(&v)
	if err != nil {
		return err
	}

	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
	v = []byte(s)

	// Create alias with UnmarshalXML method
	type Alias ReportGetPayslipByRunCompanyV2Report
	a := (*Alias)(r)

	// Unmarshal alias with cdata xml
	err = xml.Unmarshal(v, a)
	if err != nil {
		return err
	}

	// Copy alias back to original
	*r = (ReportGetPayslipByRunCompanyV2Report)(*a)
	return nil
}
