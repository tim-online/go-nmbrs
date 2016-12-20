package auth

import "encoding/xml"

type AuthHeader struct {
	XMLName xml.Name `xml:"AuthHeader"`
	Xmlns   string   `xml:"xmlns,attr"`

	Username string `xml:"Username"`
	Token    string `xml:"Token"`
}

func NewAuthHeader() *AuthHeader {
	return &AuthHeader{
		Xmlns:    "",
		Username: "",
		Token:    "",
	}
}
