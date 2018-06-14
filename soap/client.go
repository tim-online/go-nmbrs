package soap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	mediaType = "text/xml"
	charset   = "utf-8"
)

var (
	defaultUserAgent = "go-soap"
	ErrBadResponse   = errors.New("Bad response type")
)

type Client struct {
	// HTTP client used to communicate with the DO API.
	Client *http.Client

	// Debugging flag
	Debug bool

	// User agent for client
	UserAgent string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func NewClient(httpClient *http.Client) *Client {
	c := &Client{
		Client:    http.DefaultClient,
		UserAgent: defaultUserAgent,
	}

	if httpClient != nil {
		c.Client = httpClient
	}

	return c
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, response *Response) (*http.Response, error) {
	if c.Debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.Debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if response == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	err = xml.NewDecoder(httpResp.Body).Decode(response.Envelope)
	return httpResp, err
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is XML encoded and included in as the request body.
func (c *Client) NewRequest(urlStr string, request *Request) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if request != nil {
		buf.Write([]byte(xml.Header))
		err := xml.NewEncoder(buf).Encode(request.Envelope)
		if err != nil {
			return nil, err
		}
	}

	httpReq, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", mediaType, charset))
	httpReq.Header.Add("Accept", mediaType)
	httpReq.Header.Add("User-Agent", c.UserAgent)
	httpReq.Header.Add("SOAPAction", request.Action.String())
	return httpReq, nil
}

// OnRequestCompleted sets the DO API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}
