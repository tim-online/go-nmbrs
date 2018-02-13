package companies

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib"
	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	runGetListAction = "Run_GetList"
)

// List all products
func (s *Service) RunGetList(companyID int, year int) (*runGetListResponse, error) {
	request, response := newRunGetListAction(companyID, year)

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
	runGetListResponse, ok := response.Envelope.Body.Data.(*runGetListResponse)
	if ok == false {
		return runGetListResponse, soap.ErrBadResponse
	}

	return runGetListResponse, err
}

func newRunGetListAction(companyID int, year int) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newRunGetListRequest(companyID, year)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newRunGetListResponse()

	return request, response
}

type runGetListRequest struct {
	XMLName xml.Name

	CompanyID int `xml:"CompanyId"`
	Year      int `xml:"Year"`
}

func newRunGetListRequest(companyID int, year int) *runGetListRequest {
	return &runGetListRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: runGetListAction,
		},
		CompanyID: companyID,
		Year:      year,
	}
}

type runGetListResponse struct {
	Runs []Run `xml:"Run_GetListResult>RunInfo"`
}

func newRunGetListResponse() *runGetListResponse {
	return &runGetListResponse{}
}

type Run struct {
	ID          int      `xml:"ID"`
	Number      int      `xml:"Number"`
	Year        int      `xml:"Year"`
	PeriodStart int      `xml:"PeriodStart"`
	PeriodEnd   int      `xml:"PeriodEnd"`
	Description string   `xml:"Description"`
	RunAt       lib.Time `xml:"RunAt"`
	IsLocked    bool     `xml:"IsLocked"`
}
