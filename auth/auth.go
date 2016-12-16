package auth

import "encoding/xml"

type AuthHeader struct {
	XMLName  xml.Name `xml:"https://api.nmbrs.nl/soap/v2.1/CompanyService AuthHeader"`
	Username string   `xml:"Username"`
	Token    string   `xml:"Token"`
}

func NewAuthHeader() *AuthHeader {
	return &AuthHeader{
		Username: "",
		Token:    "",
	}
}
