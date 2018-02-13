package companies

import (
	"encoding/xml"
	"strings"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	wageCodesByRunCompanyV2Action = "Reports_GetWageCodesByRunCompany_v2"
)

// Returns the Journal Report for Company
func (s *Service) WageCodesByRunCompanyV2(companyID int, runID int, year int) (*wageCodesByRunCompanyV2Response, error) {
	request, response := newWageCodesByRunCompanyV2Actions(companyID, runID, year)
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
	wageCodesByRunCompanyV2Response, ok := response.Envelope.Body.Data.(*wageCodesByRunCompanyV2Response)
	if ok == false {
		return wageCodesByRunCompanyV2Response, soap.ErrBadResponse
	}

	return wageCodesByRunCompanyV2Response, err
}

func newWageCodesByRunCompanyV2Actions(companyID int, runID int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newWageCodesByRunCompanyV2Request(companyID, runID, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newWageCodesByRunCompanyV2Response()

	return request, response
}

type ageCodesByRunCompanyV2Request struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
	RunID     int `xml:"RunId"`
	Year      int `xml:"intYear"`
}

func newWageCodesByRunCompanyV2Request(companyID int, runID int, year int) *ageCodesByRunCompanyV2Request {
	return &ageCodesByRunCompanyV2Request{
		XMLName: xml.Name{
			Space: xmlns,
			Local: wageCodesByRunCompanyV2Action,
		},
		CompanyID: companyID,
		RunID:     runID,
		Year:      year,
	}
}

type wageCodesByRunCompanyV2Response struct {
	Report WageCodesByRunCompanyV2Report `xml:"Reports_GetWageCodesByRunCompany_v2Result"`
}

func newWageCodesByRunCompanyV2Response() *wageCodesByRunCompanyV2Response {
	return &wageCodesByRunCompanyV2Response{}
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

func (r *WageCodesByRunCompanyV2Report) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	v := []byte{}
	err := d.DecodeElement(&v, &start)
	// err := d.Decode(&v)
	if err != nil {
		return err
	}

	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
	v = []byte(s)

	// Create alias with UnmarshalXML method
	type Alias WageCodesByRunCompanyV2Report
	a := (*Alias)(r)

	// Unmarshal alias with cdata xml
	err = xml.Unmarshal(v, a)
	if err != nil {
		return err
	}

	// Copy alias back to original
	*r = (WageCodesByRunCompanyV2Report)(*a)
	return nil
}
