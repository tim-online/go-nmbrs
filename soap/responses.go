package soap

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		// @TODO: make error response for soap
		log.Fatal("Make error response for soap")
		err := xml.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string `xml:"message"`

	// RequestID returned from the API, useful to contact support.
	RequestID string `xml:"request_id"`
}

// <soap:Body>
// 	<soap:Fault>
// 		<faultcode>soap:Client</faultcode>
// 		<faultstring>Unable to handle request without a valid action parameter. Please supply a valid soap action.</faultstring>
// 		<detail />
// 		</soap:Fault>
// 	</soap:Body>
// </soap:Envelope>

func (r *ErrorResponse) Error() string {
	if r.RequestID != "" {
		return fmt.Sprintf("%v %v: %d (request %q) %v",
			r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.RequestID, r.Message)
	}
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
