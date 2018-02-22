package fileexplorer

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	uploadFileAction = "FileExplorer_UploadFile"
)

// Uploads document for company
func (s *Service) UploadFile(companyID int, name string, subFolder string, body []byte) (*uploadFileResponse, error) {
	request, response := newUploadFileAction(companyID, name, subFolder, body)

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
	uploadFileResponse, ok := response.Envelope.Body.Data.(*uploadFileResponse)
	if ok == false {
		return uploadFileResponse, soap.ErrBadResponse
	}

	return uploadFileResponse, err
}

func newUploadFileAction(companyID int, name string, subFolder string, body []byte) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	requestBody := newUploadFileRequest(companyID, name, subFolder, body)
	request.Envelope.Body.Data = requestBody
	request.Action = url.MustParse(requestBody.XMLName.Space + "/" + requestBody.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newUploadFileResponse()

	return request, response
}

type uploadFileRequest struct {
	XMLName xml.Name

	CompanyID int    `xml:"CompanyId"`
	Name      string `xml:"StrDocumentName"`
	SubFolder string `xml:"StrDocumentSubFolder"`
	Body      []byte `xml:"Body"`
}

func newUploadFileRequest(companyID int, name string, subFolder string, body []byte) *uploadFileRequest {
	return &uploadFileRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: uploadFileAction,
		},
		CompanyID: companyID,
		Name:      name,
		SubFolder: subFolder,
		Body:      body,
	}
}

type uploadFileResponse struct {
}

func newUploadFileResponse() *uploadFileResponse {
	return &uploadFileResponse{}
}
