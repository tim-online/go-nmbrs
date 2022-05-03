package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	journalsReportByCompanyBackgroundAction = "Reports_GetJournalsReportByCompany_Background"
)

// Generate company journal report within specified period
func (s *Service) JournalsReportByCompanyBackground(companyID, startPeriod, endPeriod, year int) (*journalsReportByCompanyBackgroundResponse, error) {
	request, response := newJournalsReportByCompanyBackground(companyID, startPeriod, endPeriod, year)

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
	reportResponse, ok := response.Envelope.Body.Data.(*journalsReportByCompanyBackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newJournalsReportByCompanyBackground(companyID, startPeriod, endPeriod, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newJournalsReportByCompanyBackgroundRequest(companyID, startPeriod, endPeriod, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newJournalsReportByCompanyBackgroundResponse()

	return request, response
}

type journalsReportByCompanyBackgroundRequest struct {
	XMLName xml.Name

	CompanyID   int `xml:"companyId"`
	StartPeriod int `xml:"startPeriod"`
	EndPeriod   int `xml:"endPeriod"`
	Year        int `xml:"year"`
}

func newJournalsReportByCompanyBackgroundRequest(companyID, startPeriod, endPeriod, year int) *journalsReportByCompanyBackgroundRequest {
	return &journalsReportByCompanyBackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: journalsReportByCompanyBackgroundAction,
		},

		CompanyID:   companyID,
		StartPeriod: startPeriod,
		EndPeriod:   endPeriod,
		Year:        year,
	}
}

type journalsReportByCompanyBackgroundResponse struct {
	JournalsReportByCompanyBackgroundResult string `xml:"Reports_GetJournalsReportByCompany_BackgroundResult"`
}

func newJournalsReportByCompanyBackgroundResponse() *journalsReportByCompanyBackgroundResponse {
	return &journalsReportByCompanyBackgroundResponse{}
}
