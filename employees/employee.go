package employees

// Employee represents the base information of an employee
type Employee struct {
	ID          int `xml:"Id"`
	Number      int    `xml:"Number"`
	DisplayName string `xml:"DisplayName"`
}
