package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	getPayslipsByRunCompanyV2BackgroundAction = "Reports_GetPayslipByRunCompany_v2_Background"
)

func (s *Service) GetPayslipsByRunCompanyV2Background(companyID, runID, year int) (*getPayslipsByRunCompanyV2BackgroundResponse, error) {
	request, response := newGetPayslipsByRunCompanyV2Background(companyID, runID, year)

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
	reportResponse, ok := response.Envelope.Body.Data.(*getPayslipsByRunCompanyV2BackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newGetPayslipsByRunCompanyV2Background(companyID, runID, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newGetPayslipsByRunCompanyV2BackgroundRequest(companyID, runID, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newGetPayslipsByRunCompanyV2BackgroundResponse()

	return request, response
}

type getPayslipsByRunCompanyV2BackgroundRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"companyId"`
	RunID     int `xml:"RunId"`
	Year      int `xml:"year"`
}

func newGetPayslipsByRunCompanyV2BackgroundRequest(companyID, runID, year int) *getPayslipsByRunCompanyV2BackgroundRequest {
	return &getPayslipsByRunCompanyV2BackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: getPayslipsByRunCompanyV2BackgroundAction,
		},

		CompanyID: companyID,
		RunID:     runID,
		Year:      year,
	}
}

type getPayslipsByRunCompanyV2BackgroundResponse struct {
	GetPayslipByRunCompanyV2BackgroundResult string `xml:"Reports_GetPayslipByRunCompany_v2_BackgroundResult"`
}

func newGetPayslipsByRunCompanyV2BackgroundResponse() *getPayslipsByRunCompanyV2BackgroundResponse {
	return &getPayslipsByRunCompanyV2BackgroundResponse{}
}
