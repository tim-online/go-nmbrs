package reports

import (
	"encoding/xml"
	"strings"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	businessEmployeeContractsAction = "Reports_Business_EmployeeContracts"
)

// Generate company journal report within specified period
func (s *Service) BusinessEmployeeContracts() (*businessEmployeeContractResponse, error) {
	request, response := newBusinessEmployeeContracts()

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
	reportResponse, ok := response.Envelope.Body.Data.(*businessEmployeeContractResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBusinessEmployeeContracts() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBusinessEmployeeContractsRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBusinessEmployeeContractsResponse()

	return request, response
}

type businessEmployeeContractRequest struct {
	XMLName xml.Name
}

func newBusinessEmployeeContractsRequest() *businessEmployeeContractRequest {
	return &businessEmployeeContractRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: businessEmployeeContractsAction,
		},
	}
}

type businessEmployeeContractResponse struct {
	Report BusinessEmployeeContractsReport `xml:"Reports_Business_EmployeeContractsResult"`
}

func newBusinessEmployeeContractsResponse() *businessEmployeeContractResponse {
	return &businessEmployeeContractResponse{}
}

type BusinessEmployeeContractsReport struct {
	EmployeeContracts BusinessEmployeeContracts `xml:"EmployeeContractReport>Employee"`
}

func (r *BusinessEmployeeContractsReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v := []byte{}
	err := d.DecodeElement(&v, &start)
	// err := d.Decode(&v)
	if err != nil {
		return err
	}

	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
	v = []byte(s)

	// Create alias with UnmarshalXML method
	type Alias BusinessEmployeeContractsReport
	a := (*Alias)(r)

	// Unmarshal alias with cdata xml
	err = xml.Unmarshal(v, a)
	if err != nil {
		return err
	}

	// Copy alias back to original
	*r = (BusinessEmployeeContractsReport)(*a)
	return nil
}

type BusinessEmployeeContracts []BusinessEmployeeContract

type BusinessEmployeeContract struct {
	EmployeeID        int
	EmployeeNumber    int
	EmployeeName      string
	CompanyID         int
	CompanyNumber     int
	CompanyName       string
	DebtorID          int
	DebtorNumber      int
	DebtorName        string
	ServiceStartDate  lib.Date
	OutOfServiceDate  *lib.Date
	StartPeriod       int
	TotalPeriod       string
	ContractStartDate lib.Date
	ContractEndDate   *lib.Date
	Department        string
	Manager           string
}

// func (r *BusinessEmployeeContracts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	type Alias BusinessEmployeeContracts
// 	in := struct {
// 		ParttimePercentage string
// 		FTE                string
// 		TotalHoursWeek     string
// 		TotalHoursPeriod   string
// 		SalaryValue        string

// 		*Alias
// 	}{
// 		Alias: (*Alias)(r),
// 	}

// 	err := d.DecodeElement(&in, &start)
// 	if err != nil {
// 		return err
// 	}

// 	in.ParttimePercentage = strings.Replace(in.ParttimePercentage, ",", ".", -1)
// 	if in.ParttimePercentage != "-" {
// 		in.Alias.ParttimePercentage, err = strconv.ParseFloat(in.ParttimePercentage, 64)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	in.FTE = strings.Replace(in.FTE, ",", ".", -1)
// 	if in.FTE != "-" {
// 		in.Alias.FTE, err = strconv.ParseFloat(in.FTE, 64)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	in.TotalHoursWeek = strings.Replace(in.TotalHoursWeek, ",", ".", -1)
// 	if in.TotalHoursWeek != "-" {
// 		in.Alias.TotalHoursWeek, err = strconv.ParseFloat(in.TotalHoursWeek, 64)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	in.TotalHoursPeriod = strings.Replace(in.TotalHoursPeriod, ",", ".", -1)
// 	if in.TotalHoursPeriod != "-" {
// 		in.Alias.TotalHoursPeriod, err = strconv.ParseFloat(in.TotalHoursPeriod, 64)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	in.SalaryValue = strings.Replace(in.SalaryValue, ",", ".", -1)
// 	if in.SalaryValue != "-" {
// 		in.Alias.SalaryValue, err = strconv.ParseFloat(in.SalaryValue, 64)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
