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

type AccountantEmployeeContracts []AccountantEmployeeContract

type AccountantEmployeeContract struct {
	EmployeeID        string `xml:"EmployeeID"`
	EmployeeNumber    string `xml:"EmployeeNumber"`
	EmployeeName      string `xml:"EmployeeName"`
	CompanyID         string `xml:"CompanyID"`
	CompanyNumber     string `xml:"CompanyNumber"`
	CompanyName       string `xml:"CompanyName"`
	DebtorID          string `xml:"DebtorID"`
	DebtorNumber      string `xml:"DebtorNumber"`
	DebtorName        string `xml:"DebtorName"`
	ServiceStartDate  string `xml:"ServiceStartDate"`
	StartPeriod       string `xml:"StartPeriod"`
	TotalPeriod       string `xml:"TotalPeriod"`
	ContractStartDate string `xml:"ContractStartDate"`
	Function          string `xml:"Function"`
	Department        string `xml:"Department"`
	Manager           string `xml:"Manager"`
	ContractDuration  string `xml:"ContractDuration"`
	ContractHours     struct {
		Nil string `xml:"nil,attr"`
	} `xml:"ContractHours"`
	NatureOfEmployment string `xml:"NatureOfEmployment"`
	ContractType       string `xml:"ContractType"`
	HasWrittenContract string `xml:"HasWrittenContract"`
	SeniorityDate      struct {
		Nil string `xml:"nil,attr"`
	} `xml:"SeniorityDate"`
	InfluenceCode            string `xml:"InfluenceCode"`
	OutOfServiceDate         string `xml:"OutOfServiceDate"`
	ReasonEndServiceInterval string `xml:"ReasonEndServiceInterval"`
	ContractEndDate          string `xml:"ContractEndDate"`
	TrialPeriod              string `xml:"TrialPeriod"`
}
