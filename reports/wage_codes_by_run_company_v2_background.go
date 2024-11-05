package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	wageCodesByRunCompanyV2Action = "Reports_GetWageCodesByRunCompany_v2_Background"
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

	CompanyID int `xml:"companyId"`
	RunID     int `xml:"runId"`
	Year      int `xml:"year"`
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
	Reports_GetWageCodesByRunCompany_v2_BackgroundResult string `xml:"Reports_GetWageCodesByRunCompany_v2_BackgroundResult"`
}

func newWageCodesByRunCompanyV2Response() *wageCodesByRunCompanyV2Response {
	return &wageCodesByRunCompanyV2Response{}
}
