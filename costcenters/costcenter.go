package costcenters

type CostCenter struct {
	ID          int    `xml:"Id"`
	Code        int    `xml:"Code"`
	Description string `xml:"Description"`
}
