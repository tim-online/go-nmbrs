package reports

import (
	"encoding/xml"

	"github.com/tim-online/go-nmbrs/lib/url"
	"github.com/tim-online/go-nmbrs/soap"
)

const (
	backgroundTaskResultAction = "Reports_BackgroundTask_Result"
)

// Generate company journal report within specified period
func (s *Service) BackgroundTaskResult(taskID string) (*backgroundTaskResultResponse, error) {
	request, response := newBackgroundTaskResult(taskID)

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
	reportResponse, ok := response.Envelope.Body.Data.(*backgroundTaskResultResponse)
	if ok == false {
		return reportResponse, soap.ErrBadResponse
	}

	return reportResponse, err
}

func newBackgroundTaskResult(taskID string) (*soap.Request, *soap.Response) {
	request := soap.NewRequest()
	body := newBackgroundTaskResultRequest(taskID)
	request.Envelope.Body.Data = body
	request.Action = url.MustParse(body.XMLName.Space + "/" + body.XMLName.Local)

	response := soap.NewResponse()
	response.Envelope.Body.Data = newBackgroundTaskResultResponse()

	return request, response
}

type backgroundTaskResultRequest struct {
	XMLName xml.Name

	TaskID string `xml:"TaskId"`
}

func newBackgroundTaskResultRequest(taskID string) *backgroundTaskResultRequest {
	return &backgroundTaskResultRequest{
		XMLName: xml.Name{
			Space: xmlns,
			Local: backgroundTaskResultAction,
		},

		TaskID: taskID,
	}
}

type backgroundTaskResultResponse struct {
	BackgroundTaskResultResult BackgroundTaskResultResult `xml:"Reports_BackgroundTask_ResultResult"`
}

func newBackgroundTaskResultResponse() *backgroundTaskResultResponse {
	return &backgroundTaskResultResponse{}
}

type BackgroundTaskResultResult struct {
	TaskID  string `xml:"TaskId"`
	Status  string `xml:"Status"`
	Content string `xml:"Content"`
}
