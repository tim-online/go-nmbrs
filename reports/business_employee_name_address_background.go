package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	BusinessEmployeeNameAddressBackgroundAction = "Reports_Business_EmployeeNameAddress_Background"
)

// Generate company journal report within specified period
func (s *Service) BusinessEmployeeNameAddressBackground() (*BusinessEmployeeNameAddressBackgroundResponse, error) {
	request, response := newBusinessEmployeeNameAddressBackground()

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
	reportResponse, ok := response.Envelope.Body.Data.(*BusinessEmployeeNameAddressBackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBusinessEmployeeNameAddressBackground() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBusinessEmployeeNameAddressBackgroundRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBusinessEmployeeNameAddressBackgroundResponse()

	return request, response
}

type BusinessEmployeeNameAddressBackgroundRequest struct {
	XMLName xml.Name
}

func newBusinessEmployeeNameAddressBackgroundRequest() *BusinessEmployeeNameAddressBackgroundRequest {
	return &BusinessEmployeeNameAddressBackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: BusinessEmployeeNameAddressBackgroundAction,
		},
	}
}

type BusinessEmployeeNameAddressBackgroundResponse struct {
	ReportsBusinessEmployeeNameAddressBackgroundResult string `xml:"Reports_GetJournalsReportByCompany_BackgroundResult"`
}

func newBusinessEmployeeNameAddressBackgroundResponse() *BusinessEmployeeNameAddressBackgroundResponse {
	return &BusinessEmployeeNameAddressBackgroundResponse{}
}
