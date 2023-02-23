package auth

import "encoding/xml"

type AuthHeader struct {
	XMLName xml.Name

	Username string `xml:"Username"`
	Token    string `xml:"Token"`
}

func NewAuthHeader() *AuthHeader {
	return &AuthHeader{
		XMLName: xml.Name{
			Space: "",
			Local: "AuthHeader",
		},
		Username: "",
		Token:    "",
	}
}

type AuthHeaderWithDomain struct {
	XMLName xml.Name

	Username string `xml:"Username"`
	Token    string `xml:"Token"`
	Domain   string `xml:"Domain"`
}

func NewAuthHeaderWithDomain() *AuthHeaderWithDomain {
	return &AuthHeaderWithDomain{
		XMLName: xml.Name{
			Space: "",
			Local: "AuthHeaderWithDomain",
		},
		Username: "",
		Token:    "",
		Domain:   "",
	}
}
