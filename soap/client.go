package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
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
)

type Client struct {
	// HTTP client used to communicate with the DO API.
	Client *http.Client

	// User agent for client
	UserAgent string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func NewClient() *Client {
	c := &Client{
		Client:    http.DefaultClient,
		UserAgent: defaultUserAgent,
	}
	return c
}

// Do sends an API request and returns the API response. The API response is XML decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	dump, err := httputil.DumpRequestOut(req, true)
	fmt.Println(string(dump))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	dump, err = httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))

	// check if the response isn't an error
	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	// check the provided interface parameter
	if v == nil {
		return resp, err
	}

	// interface implements io.Writer: write Body to it
	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return resp, err
	}

	// try to decode body into interface parameter
	err = xml.NewDecoder(resp.Body).Decode(v)
	fmt.Println("--------------")
	fmt.Printf("%+v\n", err)
	return resp, err
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is XML encoded and included in as the request body.
func (c *Client) NewRequest(urlStr string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	soapRequest := NewRequest()
	fmt.Printf("%+v\n", soapRequest)
	soapRequest.Envelope.Body.Content = body

	buf := new(bytes.Buffer)
	if body != nil {
		buf.Write([]byte(xml.Header))
		err := xml.NewEncoder(buf).Encode(soapRequest.Envelope)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", mediaType, charset))
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("SOAPAction", "https://api.nmbrs.nl/soap/v2.1/CompanyService/List_GetAll")
	return req, nil
}

// OnRequestCompleted sets the DO API request completion callback
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}
