package employees

type Employee struct {
	ID          string `xml:"Id"`
	Number      int    `xml:"Number"`
	DisplayName string `xml:"DisplayName"`
}
