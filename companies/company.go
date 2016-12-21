package companies

const (
	None   Tijdvak = "None"
	Month  Tijdvak = "Month"
	Week4  Tijdvak = "Week4"
	Month6 Tijdvak = "Month6"
	Year   Tijdvak = "Year"
)

type Company struct {
	// XMLName             xml.Name `xml:"Company"`
	ID                  int     `xml:"ID"`
	Number              int     `xml:"Number"`
	PhoneNumber         string  `xml:"PhoneNumber"`
	FaxNumber           string  `xml:"FaxNumber"`
	Email               string  `xml:"Email"`
	Website             string  `xml:"Website"`
	LoonaangifteTijdvak Tijdvak `xml:"LoonaangifteTijdvak"`
	KvkNr               string  `xml:"KvkNr"`
}

type Tijdvak string
