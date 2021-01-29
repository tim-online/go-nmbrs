package address

type Address struct {
	ID                  int    `xml:"Id"`
	Default             bool   `xml:"Default"`
	Street              string `xml:"Street"`
	HouseNumber         string `xml:"HouseNumber"`
	HouseNumberAddition string `xml:"HouseNumberAddition"`
	PostalCode          string `xml:"PostalCode"`
	City                string `xml:"City"`
	StateProvince       string `xml:"StateProvince"`
	CountryISOCode      string `xml:"CountryISOCode"`
	Type                string `xml:"Type"` // HomeAddress or PostAddress or AbsenceAddress or HolidaysAddress or WeekendAddress
}
