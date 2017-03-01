package lib

import gotime "time"

// type Time gotime.Time
type Time struct {
	gotime.Time
}

func (t *Time) UnmarshalText(text []byte) (err error) {
	value := string(text)
	if value == "" {
		return nil
	}

	layout := "2006-01-02T15:04:05"
	time, err := gotime.Parse(layout, string(text))
	// newTime := Time(time)
	// t = &newTime
	*t = Time{Time: time}
	return err
}
