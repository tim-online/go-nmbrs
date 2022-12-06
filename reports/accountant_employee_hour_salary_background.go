package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	AccountantEmployeeHoursSalaryBackgroundAction = "Reports_Accountant_EmployeeHoursSalary_Background"
)

// Generate company journal report within specified period
func (s *Service) AccountantEmployeeHoursSalaryBackground() (*AccountantEmployeeHoursSalaryBackgroundResponse, error) {
	request, response := newAccountantEmployeeHoursSalaryBackground()

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
	reportResponse, ok := response.Envelope.Body.Data.(*AccountantEmployeeHoursSalaryBackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newAccountantEmployeeHoursSalaryBackground() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newAccountantEmployeeHoursSalaryBackgroundRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newAccountantEmployeeHoursSalaryBackgroundResponse()

	return request, response
}

type AccountantEmployeeHoursSalaryBackgroundRequest struct {
	XMLName xml.Name
}

func newAccountantEmployeeHoursSalaryBackgroundRequest() *AccountantEmployeeHoursSalaryBackgroundRequest {
	return &AccountantEmployeeHoursSalaryBackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: AccountantEmployeeHoursSalaryBackgroundAction,
		},
	}
}

type AccountantEmployeeHoursSalaryBackgroundResponse struct {
	ReportsAccountantEmployeeHoursSalaryBackgroundResult string `xml:"Reports_Accountant_EmployeeHoursSalary_BackgroundResult"`
}

func newAccountantEmployeeHoursSalaryBackgroundResponse() *AccountantEmployeeHoursSalaryBackgroundResponse {
	return &AccountantEmployeeHoursSalaryBackgroundResponse{}
}
