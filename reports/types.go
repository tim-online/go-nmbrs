package reports

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/tim-online/go-nmbrs/lib"
)

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

type EmployeeHoursSalaryReport struct {
	BusinessEmployeeHoursSalaryItems `xml:"EmployeeHoursSalaryReport>Employee"`
}

type BusinessEmployeeContractReport struct {
	BusinessEmployeeContracts `xml:"EmployeeContractReport>Employee"`
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
