package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	BusinessEmployeeHoursSalaryBackgroundAction = "Reports_Business_EmployeeHoursSalary_Background"
)

// Generate company journal report within specified period
func (s *Service) BusinessEmployeeHoursSalaryBackground() (*BusinessEmployeeHoursSalaryBackgroundResponse, error) {
	request, response := newBusinessEmployeeHoursSalaryBackground()

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
	reportResponse, ok := response.Envelope.Body.Data.(*BusinessEmployeeHoursSalaryBackgroundResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBusinessEmployeeHoursSalaryBackground() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBusinessEmployeeHoursSalaryBackgroundRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBusinessEmployeeHoursSalaryBackgroundResponse()

	return request, response
}

type BusinessEmployeeHoursSalaryBackgroundRequest struct {
	XMLName xml.Name
}

func newBusinessEmployeeHoursSalaryBackgroundRequest() *BusinessEmployeeHoursSalaryBackgroundRequest {
	return &BusinessEmployeeHoursSalaryBackgroundRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: BusinessEmployeeHoursSalaryBackgroundAction,
		},
	}
}

type BusinessEmployeeHoursSalaryBackgroundResponse struct {
	ReportsBusinessEmployeeHoursSalaryBackgroundResult string `xml:"Reports_Business_EmployeeHoursSalary_BackgroundResult"`
}

func newBusinessEmployeeHoursSalaryBackgroundResponse() *BusinessEmployeeHoursSalaryBackgroundResponse {
	return &BusinessEmployeeHoursSalaryBackgroundResponse{}
}
