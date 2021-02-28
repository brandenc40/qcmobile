package entities

import "time"

const (
	_dateFmt      = "2006-01-02"
	_timestampFmt = "2006-01-02T15:04:05-0700"
)

// TimestampString -
type TimestampString string

// String -
func (t TimestampString) String() string {
	return string(t)
}

// Parse -
func (t TimestampString) Parse() (time.Time, error) {
	return time.Parse(_timestampFmt, t.String())
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
