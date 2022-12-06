package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	AccountantEmployeeNameAddressBackgroundAction = "Reports_Accountant_EmployeeNameAddress_Background"
)

// Generate company journal report within specified period
func (s *Service) AccountantEmployeeNameAddressBackground() (*AccountantEmployeeNameAddressBackgroundResponse, error) {
	request, response := newAccountantEmployeeNameAddressBackground()

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
	reportResponse, ok := response.Envelope.Body.Data.(*AccountantEmployeeNameAddressBackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newAccountantEmployeeNameAddressBackground() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newAccountantEmployeeNameAddressBackgroundRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newAccountantEmployeeNameAddressBackgroundResponse()

	return request, response
}

type AccountantEmployeeNameAddressBackgroundRequest struct {
	XMLName xml.Name
}

func newAccountantEmployeeNameAddressBackgroundRequest() *AccountantEmployeeNameAddressBackgroundRequest {
	return &AccountantEmployeeNameAddressBackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: AccountantEmployeeNameAddressBackgroundAction,
		},
	}
}

type AccountantEmployeeNameAddressBackgroundResponse struct {
	ReportsAccountantEmployeeNameAddressBackgroundResult string `xml:"Reports_Accountant_EmployeeNameAddress_BackgroundResult"`
}

func newAccountantEmployeeNameAddressBackgroundResponse() *AccountantEmployeeNameAddressBackgroundResponse {
	return &AccountantEmployeeNameAddressBackgroundResponse{}
}
