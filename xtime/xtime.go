/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xtime

import (
	"time"
)

// Version returns package version
func Version() string {
	return "0.1.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// Now returns time.Now
func Now() time.Time {
	return time.Now()
}

// String returns string of now
func String() string {
	return TimeToStr(S())
}

// S returns unix timestamp in Second
func S() int64 {
	return Now().Unix()
}

// Ns returns unix timestamp in Nanosecond
func Ns() int64 {
	return Now().UnixNano()
}

// Us returns unix timestamp in Microsecond
func Us() int64 {
	return Now().UnixNano() / int64(time.Microsecond)
}

// Ms returns unix timestamp in Millisecond
func Ms() int64 {
	return Now().UnixNano() / int64(time.Millisecond)
}

// Sleep n Second
func Sleep(n int64) {
	time.Sleep(time.Duration(n) * time.Second)
}

// Usleep n Microsecond
func Usleep(n int64) {
	time.Sleep(time.Duration(n) * time.Microsecond)
}

// StrToTime returns unix timestamp in Second of string
func StrToTime(s string, layout ...string) (int64, error) {
	format := "2006-01-02 15:04:05"
	if len(layout) > 0 {
		format = layout[0]
	} else {
		if len(s) == 10 {
			format = "2006-01-02"
		}
	}

	t, err := time.ParseInLocation(format, s, time.Local)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}

// TimeToStr returns time string of unix timestamp in Second
func TimeToStr(n int64, layout ...string) string {
	format := "2006-01-02 15:04:05"
	if len(layout) > 0 {
		format = layout[0]
	}

	return time.Unix(n, 0).Format(format)
}
