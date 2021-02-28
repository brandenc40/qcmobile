package entities

import "time"

const _dateFmt = "2006-01-02"

// TimestampString -
type TimestampString string

// String -
func (t TimestampString) String() string {
	return string(t)
}

// Parse -
func (t TimestampString) Parse() (time.Time, error) {
	return time.Parse(time.RFC3339, t.String())
}

// DateString -
type DateString string

// String -
func (d DateString) String() string {
	return string(d)
}

// Parse -
func (d DateString) Parse() (time.Time, error) {
	return time.Parse(_dateFmt, d.String())
}
