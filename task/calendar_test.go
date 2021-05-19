package task

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		d  string
		t  string
		wd string
	}
	tests := []struct {
		name    string
		args    args
		want    *Calendar
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.d, tt.args.t, tt.args.wd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validate(t *testing.T) {
	type args struct {
		i       string
		min     int
		allowed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: too_long, not_allowed
		{"Time_pass", args{i: "1223*", min: 3, allowed: "*123456789"}, true},
		{"Time_too_short", args{i: "23", min: 3, allowed: "*123456789"}, true},

		{"date_pass", args{i: "*12*59", min: 3, allowed: "*123456789"}, true},
		{"date_too_short", args{i: "23", min: 3, allowed: "*123456789"}, true},

		{"weekday_pass", args{i: "12", min: 1, allowed: "*123456"}, true},
		{"weeday_too_short", args{i: "", min: 1, allowed: "*123456"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validate(tt.args.i, tt.args.min, tt.args.allowed); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
