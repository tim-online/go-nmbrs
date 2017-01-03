package hourcodes

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listAction = "HourModel_GetHourCodes"
)

// List all products
func (s *Service) List(companyID int) (*listResponse, error) {
	request, response := newListAction(companyID)
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
	listResponse, ok := response.Envelope.Body.Data.(*listResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func newListAction(companyID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newListRequest(companyID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListResponse()

	return request, response
}

type listRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
}

func newListRequest(companyID int) *listRequest {
	return &listRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listAction,
		},

		CompanyID: companyID,
	}
}

type listResponse struct {
	HourCodes []HourCode `xml:"HourModel_GetHourCodesResult>HourCode"`
}

func newListResponse() *listResponse {
	return &listResponse{}
}
