package employees

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	uploadDocumentAction = "EmployeeDocument_UploadDocument"
)

// Uploads document for company
func (s *Service) UploadDocument(employeeID int, name string, body string, documentType string) (*uploadDocumentResponse, error) {
	request, response := newUploadDocumentAction(employeeID, name, body, documentType)

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
	uploadDocumentResponse, ok := response.Envelope.Body.Data.(*uploadDocumentResponse)
	if ok == false {
		return uploadDocumentResponse, soap.ErrBadResponse
	}

	return uploadDocumentResponse, err
}

func newUploadDocumentAction(employeeID int, name string, body string, documentType string) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	requestBody := newUploadDocumentRequest(employeeID, name, body, documentType)
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newUploadDocumentResponse()

	return request, response
}

type uploadDocumentRequest struct {
	XMLName xml.Name

	EmployeeID   int    `xml:"EmployeeId"`
	Name         string `xml:"StrDocumentName"`
	Body         string `xml:"Body"`
	DocumentType string `xml:"GuidDocumentType"`
}

func newUploadDocumentRequest(employeeID int, name string, body string, documentType string) *uploadDocumentRequest {
	return &uploadDocumentRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: uploadDocumentAction,
		},
		EmployeeID:   employeeID,
		Name:         name,
		Body:         body,
		DocumentType: documentType,
	}
}

type uploadDocumentResponse struct {
}

func newUploadDocumentResponse() *uploadDocumentResponse {
	return &uploadDocumentResponse{}
}
