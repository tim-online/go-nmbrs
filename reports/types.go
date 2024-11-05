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

type WageCodesByRunCompanyV2Report struct {
    Reports []struct {
        EmployeeID int `xml:"employeeid"`
        Period     int `xml:"period"`
        Year       int `xml:"year"`
        RunNumber  int `xml:"runnumber"`
        Lines      []struct {
            Code        int     `xml:"code"`
            Description string  `xml:"description"`
            Value       float64 `xml:"value"`
        } `xml:"lines>line"`
    } `xml:"report"`
}
