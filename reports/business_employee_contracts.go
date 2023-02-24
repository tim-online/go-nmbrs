package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	businessEmployeeContractsBackgroundAction = "Reports_Business_EmployeeContracts_Background"
)

// Generate company journal report within specified period
func (s *Service) BusinessEmployeeContractsBackground() (*businessEmployeeContractResponse, error) {
	request, response := newBusinessEmployeeContractsBackground()

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
	reportResponse, ok := response.Envelope.Body.Data.(*businessEmployeeContractResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBusinessEmployeeContractsBackground() (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBusinessEmployeeContractsBackgroundRequest()
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBusinessEmployeeContractsBackgroundResponse()

	return request, response
}

type businessEmployeeContractRequest struct {
	XMLName xml.Name
}

func newBusinessEmployeeContractsBackgroundRequest() *businessEmployeeContractRequest {
	return &businessEmployeeContractRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: businessEmployeeContractsBackgroundAction,
		},
	}
}

type businessEmployeeContractResponse struct {
	ReportsBusinessEmployeeContractsBackgroundResult string `xml:"Reports_Business_EmployeeContracts_BackgroundResult"`
}

func newBusinessEmployeeContractsBackgroundResponse() *businessEmployeeContractResponse {
	return &businessEmployeeContractResponse{}
}

// type BusinessEmployeeContractsBackgroundReport struct {
// 	EmployeeContracts BusinessEmployeeContractsBackground `xml:"EmployeeContractReport>Employee"`
// }

// func (r *BusinessEmployeeContractsBackgroundReport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	v := []byte{}
// 	err := d.DecodeElement(&v, &start)
// 	// err := d.Decode(&v)
// 	if err != nil {
// 		return err
// 	}

// 	s := strings.Replace(string(v), "utf-16", "utf-8", -1)
// 	v = []byte(s)

// 	// Create alias with UnmarshalXML method
// 	type Alias BusinessEmployeeContractsBackgroundReport
// 	a := (*Alias)(r)

// 	// Unmarshal alias with cdata xml
// 	err = xml.Unmarshal(v, a)
// 	if err != nil {
// 		return err
// 	}

// 	// Copy alias back to original
// 	*r = (BusinessEmployeeContractsBackgroundReport)(*a)
// 	return nil
// }
