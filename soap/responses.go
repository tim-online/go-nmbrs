package soap

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a XML response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// @TODO: figure out nmbrs errors
	errorResponse := &ErrorResponse{Response: r}

	// check content-type (text/xml; charset=utf-8)
	header := r.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != "text/xml" {
		errorResponse.Message = fmt.Sprintf("Expected Content-Type \"text/xml\", got \"%s\"", contentType)
		return errorResponse
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert xml to struct
	err = xml.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Message = fmt.Sprintf("Malformed json response")
		return errorResponse
	}

	return errorResponse
}

// An ErrorResponse reports the error caused by an API request
// <soap:Body>
// 	<soap:Fault>
// 		<faultcode>soap:Client</faultcode>
// 		<faultstring>Unable to handle request without a valid action parameter. Please supply a valid soap action.</faultstring>
// 		<detail />
// 		</soap:Fault>
// 	</soap:Body>
// </soap:Envelope>
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Fault code
	Code string `xml:"Body>Fault>faultcode"`

	// Fault message
	Message string `xml:"Body>Fault>faultstring"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
