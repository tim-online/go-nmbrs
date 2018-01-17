package companies

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listByDebtorAction = "List_GetByDebtor"
)

// List all products
func (s *Service) ListByDebtor(debtorID int) (*listByDebtorResponse, error) {
	request, response := newListByDebtorAction(debtorID)

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
	listResponse, ok := response.Envelope.Body.Data.(*listByDebtorResponse)
	if ok == false {
		return listResponse, soap.ErrBadResponse
	}

	return listResponse, err
}

func newListByDebtorAction(debtorID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newListByDebtorRequest()
	body.DebtorID = debtorID
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListByDebtorResponse()

	return request, response
}

type listByDebtorRequest struct {
	XMLName xml.Name

	DebtorID int `xml:"DebtorId"`
}

func newListByDebtorRequest() *listByDebtorRequest {
	return &listByDebtorRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listByDebtorAction,
		},
	}
}

type listByDebtorResponse struct {
	Companies []Company `xml:"List_GetByDebtorResult>Company"`
}

func newListByDebtorResponse() *listByDebtorResponse {
	return &listByDebtorResponse{}
}
