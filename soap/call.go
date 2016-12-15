package soap

import "encoding/xml"

type Request struct {
	Envelope *Envelope `xml:"Envelope"`
}

func NewRequest() *Request {
	return &Request{
		Envelope: NewEnvelope(),
	}
}

type Response struct {
	Envelope *Envelope `xml:"Envelope"`
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLNS   string   `xml:"xmlns,attr""`

	Header *Header `xml:"Header"`
	Body   *Body   `xml:"Body"`
}

func NewEnvelope() *Envelope {
	return &Envelope{
		XMLNS:  "http://schemas.xmlsoap.org/soap/envelope/",
		Header: NewHeader(),
		Body:   NewBody(),
	}
}

type Header struct {
	AuthHeader *AuthHeader `xml:"AuthHeader"`
}

func NewHeader() *Header {
	return &Header{
		AuthHeader: NewAuthHeader(),
	}
}

type AuthHeader struct {
	Username string `xml:"Username"`
	Token    string `xml:"Token"`
}

func NewAuthHeader() *AuthHeader {
	return &AuthHeader{
		Username: "leon@tim-online.nl",
		Token:    "c20cb2ee55cd454b9143be5349da3f0c",
	}
}

// <com:AuthHeader>
// 	<!--Optional:-->
// 	<com:Username>leon@tim-online.nl</com:Username>
// 	<!--Optional:-->
// 	<com:Token>c20cb2ee55cd454b9143be5349da3f0c</com:Token>
// </com:AuthHeader>

type Body struct {
	Content interface{}
}

func NewBody() *Body {
	return &Body{}
}

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
