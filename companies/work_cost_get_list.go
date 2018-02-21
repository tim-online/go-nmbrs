package companies

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	workCostGetListAction = "WorkCost_GetList"
)

// List all products
func (s *Service) WorkCostGetList(companyID int, year int) (*workCostGetListResponse, error) {
	request, response := newWorkCostGetListAction(companyID, year)

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
	workCostGetListResponse, ok := response.Envelope.Body.Data.(*workCostGetListResponse)
	if ok == false {
		return workCostGetListResponse, soap.ErrBadResponse
	}

	return workCostGetListResponse, err
}

func newWorkCostGetListAction(companyID int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newWorkCostGetListRequest(companyID, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newWorkCostGetListResponse()

	return request, response
}

type workCostGetListRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
	Year      int `xml:"Year"`
}

func newWorkCostGetListRequest(companyID int, year int) *workCostGetListRequest {
	return &workCostGetListRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: workCostGetListAction,
		},
		CompanyID: companyID,
		Year:      year,
	}
}

type workCostGetListResponse struct {
	WorkCosts []WorkCost `xml:"WorkCost_GetListResult>WorkCost"`
}

func newWorkCostGetListResponse() *workCostGetListResponse {
	return &workCostGetListResponse{}
}

type WorkCost struct {
	Period                 int
	Year                   int
	WorkCostPayroll        float64
	WorkCostFinancial      float64
	FiscalWage             float64
	WorkCostAvailableSpace float64
	WorkCostBase           float64
	WorkCostToPay          float64
	WorkCostEstimated      float64
	WorkCostPaid           float64
}
