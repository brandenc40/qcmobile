package qcmobile

import "time"

const (
	_dateFmt      = "2006-01-02"
	_timestampFmt = "2006-01-02T15:04:05-0700"
)

// Timestamp -
type Timestamp string

// String -
func (t Timestamp) String() string {
	return string(t)
}

// Parse -
func (t Timestamp) Parse() (time.Time, error) {
	return time.Parse(_timestampFmt, t.String())
}

// Date -
type Date string

// String -
func (d Date) String() string {
	return string(d)
}

// Parse -
func (d Date) Parse() (time.Time, error) {
	return time.Parse(_dateFmt, d.String())
}
