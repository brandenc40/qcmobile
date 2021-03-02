package entities

import (
	"reflect"
	"testing"
	"time"
)

func TestTimestampString_Parse(t *testing.T) {
	tests := []struct {
		name    string
		t       Timestamp
		want    time.Time
		wantErr bool
	}{
		{
			"test",
			"2021-02-28T07:25:05.638+0000",
			time.Date(2021, 2, 28, 7, 25, 5, 638000000, time.FixedZone("", 0)),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateString_Parse(t *testing.T) {
	tests := []struct {
		name    string
		d       Date
		want    time.Time
		wantErr bool
	}{
		{
			"test",
			"2021-02-28",
			time.Date(2021, 2, 28, 0, 0, 0, 0, time.UTC),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
