package address

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	listAction = "Address_GetListCurrent"
)

// List all products
func (s *Service) ListCurrent(employeeID int) (*listCurrentResponse, error) {
	request, response := newListCurrentAction(employeeID)

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
	listCurrentResponse, ok := response.Envelope.Body.Data.(*listCurrentResponse)
	if ok == false {
		return listCurrentResponse, soap.ErrBadResponse
	}

	return listCurrentResponse, err
}

func newListCurrentAction(employeeID int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newListCurrentRequest()
	body.EmployeeID = employeeID
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newListResponse()

	return request, response
}

type listCurrentRequest struct {
	XMLName xml.Name

	EmployeeID int `xml:"EmployeeId"`
}

func newListCurrentRequest() *listCurrentRequest {
	return &listCurrentRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: listAction,
		},
	}
}

type listCurrentResponse struct {
	Addresses []Address `xml:"Address_GetListCurrentResult>EmployeeAddress"`
}

func newListResponse() *listCurrentResponse {
	return &listCurrentResponse{}
}
