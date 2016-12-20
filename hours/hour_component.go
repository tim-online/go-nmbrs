package hours

type HourComponent struct {
	ID       int     `xml:"Id"`
	HourCode int     `xml:"HourCode"`
	Hours    float64 `xml:"Hours"`
}
