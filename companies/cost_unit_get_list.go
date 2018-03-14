package companies

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	costUnit_GetListAction = "CostUnit_GetList"
)

// Returns kostensoorten that belong to a company
func (s *Service) CostUnit_GetList(companyID int) (*CostUnit_GetListResponse, error) {
	// get a new request & response envelope
	request, response := s.NewCostUnit_GetListAction(companyID)

	// copy authheader to new envelope
	request.Envelope.Header.Data = s.AuthHeader

	// create a new HTTP request from the SOAP envelope
	httpReq, err := s.Client.NewRequest(s.Endpoint.String(), request)
	if err != nil {
		return nil, err
	}

	// make the HTTP request
	_, err = s.Client.Do(httpReq, response)
	if err != nil {
		return nil, err
	}

	// @TODO: check if this can be better
	listResponse, ok := response.Envelope.Body.Data.(*CostUnit_GetListResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func (s *Service) NewCostUnit_GetListAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := NewCostUnit_GetListRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = NewCostUnit_GetListResponse()

	return request, response
}

type CostUnit_GetListRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
}

func NewCostUnit_GetListRequest(companyID int) *CostUnit_GetListRequest {
	return &CostUnit_GetListRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: costUnit_GetListAction,
		},

		CompanyID: companyID,
	}
}

type CostUnit_GetListResponse struct {
	CostUnits CostUnits `xml:"CostUnit_GetListResult>CostUnit"`
}

func NewCostUnit_GetListResponse() *CostUnit_GetListResponse {
	return &CostUnit_GetListResponse{}
}

type CostUnits []CostUnit

type CostUnit struct {
	ID          int    `xml:"Id"`
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}
