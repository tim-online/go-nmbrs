package costcenters

type CostCenter struct {
	ID          int    `xml:"Id"`
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}
