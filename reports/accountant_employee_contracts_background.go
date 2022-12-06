package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	AccountantEmployeeContractsBackgroundAction = "Reports_Accountant_EmployeeContracts_Background"
)

// Generate company journal report within specified period
func (s *Service) AccountantEmployeeContractsBackground() (*AccountantEmployeeContractsBackgroundResponse, error) {
	request, response := newAccountantEmployeeContractsBackground()

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
	reportResponse, ok := response.Envelope.Body.Data.(*AccountantEmployeeContractsBackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newAccountantEmployeeContractsBackground() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newAccountantEmployeeContractsBackgroundRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newAccountantEmployeeContractsBackgroundResponse()

	return request, response
}

type AccountantEmployeeContractsBackgroundRequest struct {
	XMLName xml.Name
}

func newAccountantEmployeeContractsBackgroundRequest() *AccountantEmployeeContractsBackgroundRequest {
	return &AccountantEmployeeContractsBackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: AccountantEmployeeContractsBackgroundAction,
		},
	}
}

type AccountantEmployeeContractsBackgroundResponse struct {
	ReportsAccountantEmployeeContractsBackgroundResult string `xml:"Reports_Accountant_EmployeeContracts_BackgroundResult"`
}

func newAccountantEmployeeContractsBackgroundResponse() *AccountantEmployeeContractsBackgroundResponse {
	return &AccountantEmployeeContractsBackgroundResponse{}
}
