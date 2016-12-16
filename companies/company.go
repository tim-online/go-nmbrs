package companies

type Company struct {
	// XMLName             xml.Name `xml:"Company"`
	ID                  string `xml:"ID"`
	Number              int    `xml:"Number"`
	PhoneNumber         string `xml:"PhoneNumber"`
	FaxNumber           string `xml:"FaxNumber"`
	Email               string `xml:"Email"`
	LoonaangifteTijdvak string `xml:"LoonaangifteTijdvak"`
}
