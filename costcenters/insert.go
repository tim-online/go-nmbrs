package costcenters

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

// https://api.nmbrs.nl/soap/v2.1/CompanyService.asmx?op=CostCenter_Insert

const (
	insertAction = "CostCenter_Insert"
)

// Update or insert a kostenplaats into a company
func (s *Service) Insert(requestBody *InsertRequest) (*InsertResponse, error) {
	responseBody := &InsertResponse{}
	request, response := newInsertAction(requestBody, responseBody)

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
	insertResponse, ok := response.Envelope.Body.Data.(*InsertResponse)
	if ok == false {
		return insertResponse, soap.ErrBadResponse
	}

	return insertResponse, err
}

func newInsertAction(requestBody *InsertRequest, responseBody *InsertResponse) (*soap.Request, *soap.Response) {
	requestBody.init()
	request := soap.NewRequest()
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = responseBody

	return request, response
}

func NewInsertRequest() *InsertRequest {
	req := &InsertRequest{}
	req.init()
	return req
}

type InsertRequest struct {
	XMLName xml.Name

	CompanyID    int          `xml:"CompanyId"`
	Kostenplaats Kostenplaats `xml:"kostenplaats"`
}

func (r *InsertRequest) init() {
	r.XMLName = xml.Name{
		Space: xmlns,
		Local: insertAction,
	}
}

type InsertResponse struct {
	Result int `xml:"CostCenter_InsertResult"`
}

type Kostenplaats struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
	Id          int    `xml:"Id"`
}
