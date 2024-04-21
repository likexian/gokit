/*
 * Copyright 2012-2024 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * A toolkit for Golang development
 * https://www.likexian.com/
 */

package xtime

import (
	"errors"
	"time"
)

var (
	// ErrCanceled is canceled error
	ErrCanceled = errors.New("xtime: canceled")
	// ErrTimeouted is timeouted error
	ErrTimeouted = errors.New("xtime: timeouted")
)

// TimeCallback is a callback with one return value
type TimeCallback func() interface{}

// Version returns package version
func Version() string {
	return "0.4.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
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

// StrToTime returns unix timestamp of time string
func StrToTime(s string, layout ...string) (int64, error) {
	format := "2006-01-02 15:04:05"
	if len(layout) > 0 && layout[0] != "" {
		format = layout[0]
	} else {
		if len(s) == 10 {
			format = format[:10]
		}
	}

	t, err := time.ParseInLocation(format, s, time.Local)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}

// TimeToStr returns time string of unix timestamp, format in time.Local
func TimeToStr(n int64, layout ...string) string {
	format := "2006-01-02 15:04:05"
	if len(layout) > 0 && layout[0] != "" {
		format = layout[0]
	}

	return time.Unix(n, 0).Format(format)
}

// WithTimeout execute the callback with timeout return a chan and cancel func
func WithTimeout(fn TimeCallback, timeout time.Duration) (chan interface{}, func()) {
	q := make(chan bool)
	r := make(chan interface{})

	go func() {
		r <- fn()
	}()

	go func() {
		t := time.After(timeout)
		select {
		case <-t:
			r <- ErrTimeouted
			return
		case <-q:
			r <- ErrCanceled
			return
		}
	}()

	return r, func() { close(q) }
}

// SetTimeout execute the callback after timeout return a chan and cancel func
func SetTimeout(fn TimeCallback, timeout time.Duration) (chan interface{}, func()) {
	q := make(chan bool)
	r := make(chan interface{})

	go func() {
		t := time.After(timeout)
		select {
		case <-t:
			r <- fn()
		case <-q:
			r <- ErrCanceled
			return
		}
	}()

	return r, func() { close(q) }
}

// SetInterval execute the callback every timeout return a chan and cancel func
func SetInterval(fn TimeCallback, timeout time.Duration) (chan interface{}, func()) {
	q := make(chan bool)
	r := make(chan interface{})

	go func() {
		t := time.NewTicker(timeout)
		for {
			select {
			case <-t.C:
				go func() { r <- fn() }()
			case <-q:
				t.Stop()
				r <- ErrCanceled
				return
			}
		}
	}()

	return r, func() { close(q) }
}
