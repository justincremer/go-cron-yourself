package task

import (
	"errors"
	"strconv"
	"strings"
)

var Weekday = [7]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// TODO: add methods for intializing calendar cells

type Calendar struct {
	D  int
	T  int
	WD int
}

func Create(d string, t string, wd string) (*Calendar, error) {
	var resError error = nil
	allowedWD := "*0123456"
	allowedDT := allowedWD + "789"

	catch := func(err error) {
		if err != nil {
			resError = errors.New("Malformed time during calendar creation")
		}
	}

	if ok := validate(d, 6, allowedDT) && validate(t, 6, allowedDT) && validate(wd, 2, allowedWD); !ok {
		return nil, errors.New("Malformed time during calendar creation")
	}

	D, err := strconv.Atoi(d)
	catch(err)
	T, err := strconv.Atoi(t)
	catch(err)
	WD, err := strconv.Atoi(wd)
	catch(err)

	return &Calendar{
		D:  D,
		T:  T,
		WD: WD,
	}, resError
}

func validate(i string, min int, allowed string) bool {
	for _, c := range i {
		if !strings.ContainsRune(allowed, c) {
			return false
		}
	}

	iLen := len(i)
	return iLen >= min && iLen <= min*2
}

// {"Time_pass", args{i: "123459", min: 3, allowed: "*123456789"}, true},
// {"Time_too_short", args{i: "23", min: 3, allowed: "*123456789"}, true},

// {"date_pass", args{i: "*123459", min: 3, allowed: "*123456789"}, true},
// {"date_too_short", args{i: "23", min: 3, allowed: "*123456789"}, true},

// {"weekday_pass", args{i: "2", min: 1, allowed: "*123456"}, true},
// {"weeday_too_short", args{i: "23", min: 3, allowed: "*123456"}, true},
