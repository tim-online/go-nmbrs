package soap

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Number float64

func (n *Number) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	s = strings.Replace(s, ",", ".", -1)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*n = Number(f)
	return nil
}
