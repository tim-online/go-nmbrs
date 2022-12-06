package reports

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	businessEmployeeHoursSalaryAction = "Reports_Business_EmployeeHoursSalary"
)

// Generate company journal report within specified period
func (s *Service) BusinessEmployeeHoursSalary(companyID int, year int) (*businessEmployeeHoursSalaryResponse, error) {
	request, response := newBusinessEmployeeHoursSalary(companyID, year)

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
	reportResponse, ok := response.Envelope.Body.Data.(*businessEmployeeHoursSalaryResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBusinessEmployeeHoursSalary(companyID int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBusinessEmployeeHoursSalaryRequest(companyID, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBusinessEmployeeHoursSalaryResponse()

	return request, response
}

type businessEmployeeHoursSalaryRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
	Year      int `xml:"Year"`
}

func newBusinessEmployeeHoursSalaryRequest(companyID int, year int) *businessEmployeeHoursSalaryRequest {
	return &businessEmployeeHoursSalaryRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: businessEmployeeHoursSalaryAction,
		},

		CompanyID: companyID,
		Year:      year,
	}
}

type businessEmployeeHoursSalaryResponse struct {
	Report BusinessEmployeeHoursSalaryReport `xml:"Reports_Business_EmployeeHoursSalaryResult"`
}

func newBusinessEmployeeHoursSalaryResponse() *businessEmployeeHoursSalaryResponse {
	return &businessEmployeeHoursSalaryResponse{}
}

type BusinessEmployeeHoursSalaryReport struct {
	EmployeeHoursSalaryItems BusinessEmployeeHoursSalaryItems `xml:"EmployeeHoursSalaryReport>Employee"`
}

func (r *BusinessEmployeeHoursSalaryReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v := []byte{}
	err := d.DecodeElement(&v, &start)
	// err := d.Decode(&v)
	if err != nil {
		return err
	}

	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
	v = []byte(s)

	// Create alias with UnmarshalXML method
	type Alias BusinessEmployeeHoursSalaryReport
	a := (*Alias)(r)

	// Unmarshal alias with cdata xml
	err = xml.Unmarshal(v, a)
	if err != nil {
		return err
	}

	// Copy alias back to original
	*r = (BusinessEmployeeHoursSalaryReport)(*a)
	return nil
}

type BusinessEmployeeHoursSalaryItems []BusinessEmployeeHoursSalaryItem

type BusinessEmployeeHoursSalaryItem struct {
	EmployeeID         int     `xml:"EmployeeID"`
	EmployeeNumber     string  `xml:"EmployeeNumber"`
	EmployeeName       string  `xml:"EmployeeName"`
	CompanyID          int     `xml:"CompanyID"`
	CompanyNumber      string  `xml:"CompanyNumber"`
	CompanyName        string  `xml:"CompanyName"`
	DebtorID           int     `xml:"DebtorID"`
	DebtorNumber       string  `xml:"DebtorNumber"`
	DebtorName         string  `xml:"DebtorName"`
	Initials           string  `xml:"Initials"`
	Firstname          string  `xml:"Firstname"`
	FirstnameInFull    string  `xml:"FirstnameInFull"`
	Surname            string  `xml:"Surname"`
	ParttimePercentage float64 `xml:"ParttimePercentage"`
	FTE                float64 `xml:"FTE"`
	TotalHoursWeek     float64 `xml:"TotalHoursWeek"`
	TotalHoursPeriod   float64 `xml:"TotalHoursPeriod"`
	SalaryType         string  `xml:"SalaryType"`
	SalaryValue        float64 `xml:"SalaryValue"`
	Function           string  `xml:"Function"`
	Department         string  `xml:"Department"`
	CostCenter         string  `xml:"CostCenter"`
	InServiceDate      string  `xml:"InServiceDate"`
	Prefix             string  `xml:"Prefix"`
	OutServiceDate     string  `xml:"OutServiceDate"`
}

func (r *BusinessEmployeeHoursSalaryItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias BusinessEmployeeHoursSalaryItem
	in := struct {
		ParttimePercentage string
		FTE                string
		TotalHoursWeek     string
		TotalHoursPeriod   string
		SalaryValue        string

		*Alias
	}{
		Alias: (*Alias)(r),
	}

	err := d.DecodeElement(&in, &start)
	if err != nil {
		return err
	}

	in.ParttimePercentage = strings.Replace(in.ParttimePercentage, ",", ".", -1)
	if in.ParttimePercentage != "-" && in.ParttimePercentage != "" {
		in.Alias.ParttimePercentage, err = strconv.ParseFloat(in.ParttimePercentage, 64)
		if err != nil {
			return err
		}
	}

	in.FTE = strings.Replace(in.FTE, ",", ".", -1)
	if in.FTE != "-" && in.FTE != "" {
		in.Alias.FTE, err = strconv.ParseFloat(in.FTE, 64)
		if err != nil {
			return err
		}
	}

	in.TotalHoursWeek = strings.Replace(in.TotalHoursWeek, ",", ".", -1)
	if in.TotalHoursWeek != "-" && in.TotalHoursWeek != "" {
		in.Alias.TotalHoursWeek, err = strconv.ParseFloat(in.TotalHoursWeek, 64)
		if err != nil {
			return err
		}
	}

	in.TotalHoursPeriod = strings.Replace(in.TotalHoursPeriod, ",", ".", -1)
	if in.TotalHoursPeriod != "-" && in.TotalHoursPeriod != "" {
		in.Alias.TotalHoursPeriod, err = strconv.ParseFloat(in.TotalHoursPeriod, 64)
		if err != nil {
			return err
		}
	}

	in.SalaryValue = strings.Replace(in.SalaryValue, ",", ".", -1)
	if in.SalaryValue != "-" && in.SalaryValue != "" {
		in.Alias.SalaryValue, err = strconv.ParseFloat(in.SalaryValue, 64)
		if err != nil {
			return err
		}
	}

	return nil
}
