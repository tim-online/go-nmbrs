package soap

import (
	"encoding/xml"
	"net/url"
)

type Request struct {
	Envelope *Envelope `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Action   *url.URL  `xml:""`
}

func NewRequest() *Request {
	return &Request{
		Envelope: NewEnvelope(),
	}
}

type Response struct {
	Envelope *Envelope `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
}

func NewResponse() *Response {
	return &Response{
		Envelope: NewEnvelope(),
	}
}

// http://stackoverflow.com/questions/16202170/marshalling-xml-go-xmlname-xmlns
type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Header *Header `xml:"Header"`
	Body   *Body   `xml:"Body"`
}

func NewEnvelope() *Envelope {
	return &Envelope{
		// Xsi:           "http://www.w3.org/2001/XMLSchema-instance",
		// Soapenc:       "http://schemas.xmlsoap.org/soap/encoding/",
		// Xsd:           "http://www.w3.org/2001/XMLSchema",
		// EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		// Soap:          "http://schemas.xmlsoap.org/soap/envelope/",

		Header: NewHeader(),
		Body:   NewBody(),
	}
}

type Header struct {
	Data interface{}
}

func NewHeader() *Header {
	return &Header{
		Data: nil,
	}
}

type Body struct {
	// If the XML element contains a sub-element that hasn't matched any
	// of the above rules and the struct has a field with tag ",any",
	// unmarshal maps the sub-element to that struct field.
	Data interface{} `xml:",any"`
}

func NewBody() *Body {
	return &Body{}
}

// func (b *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	fmt.Printf("%+v\n", start)
// 	return d.DecodeElement(b.Data, &start)
// }

// @TODO: check if this would be handy
// type Call struct {
// 	Request  *Request
// 	Response *Response
// }

// see twinfield/api/soap/twinfield.go:134
// type ListRequest struct {
// 	Call

// 	action   string
// 	response *Response
// }

// type SOAPEnvelope struct {
// 	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
// 	Header  *SOAPHeader
// 	Body    SOAPBody
// }

// type SOAPHeader struct {
// 	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

// 	Header interface{}
// }

// type SOAPBody struct {
// 	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

// 	Fault   *SOAPFault  `xml:",omitempty"`
// 	Content interface{} `xml:",omitempty"`
// }

// type SOAPFault struct {
// 	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

// 	Code   string `xml:"faultcode,omitempty"`
// 	String string `xml:"faultstring,omitempty"`
// 	Actor  string `xml:"faultactor,omitempty"`
// 	Detail string `xml:"detail,omitempty"`
// }
