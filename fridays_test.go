package main

import (
	"reflect"
	"testing"
	time2 "time"
)

func TestSpecialDay_next(t *testing.T) {
	type fields struct {
		parent int        `json:"parent,omitempty"`
		date   time2.Time `json:"date"`
	}
	tests := []struct {
		name   string
		fields fields
		want   *SpecialDay
	}{
		{
			name: "sunny day 1",
			fields: fields{
				parent: Dad,
				date:   time2.Date(2023, time2.March, 31, 0, 0, 0, 0, time2.Local),
			},
			want: &SpecialDay{
				parent: Mom,
				date:   time2.Date(2023, time2.April, 21, 0, 0, 0, 0, time2.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &SpecialDay{
				parent: tt.fields.parent,
				date:   tt.fields.date,
			}
			if got := f.next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}
