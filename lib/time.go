package lib

import (
	"encoding/xml"
	gotime "time"
)

// type Time gotime.Time
type Time struct {
	gotime.Time
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s := ""
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	layout := "2006-01-02T15:04:05"
	time, err := gotime.Parse(layout, s)
	*t = Time{Time: time}
	return err
}

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "2006-01-02T15:04:05"
	s := t.Format(layout)
	return e.EncodeElement(s, start)
}

// type Time gotime.Time
type Date struct {
	gotime.Time
}

func (d *Date) UnmarshalText(text []byte) (err error) {
	value := string(text)
	if value == "" {
		return nil
	}

	layout := "2006-01-02"
	time, err := gotime.Parse(layout, string(text))
	// newTime := Time(time)
	// t = &newTime
	*d = Date{Time: time}
	return err
}
