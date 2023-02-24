package employees

type EmployeeTypes []EmployeeType

type EmployeeType struct {
	ID          int    `xml:"Id"`
	Description string `xml:"Description"`
}
