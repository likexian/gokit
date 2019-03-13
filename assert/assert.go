/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

// Version returns package version
func Version() string {
	return "0.6.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// Equal assert test value to be equal
func Equal(t *testing.T, got, exp interface{}, args ...interface{}) {
	equal(t, got, exp, 1, args...)
}

// NotEqual assert test value to be not equal
func NotEqual(t *testing.T, got, exp interface{}, args ...interface{}) {
	notEqual(t, got, exp, 1, args...)
}

// Nil assert test value to be nil
func Nil(t *testing.T, got interface{}, args ...interface{}) {
	equal(t, got, nil, 1, args...)
}

// NotNil assert test value to be not nil
func NotNil(t *testing.T, got interface{}, args ...interface{}) {
	notEqual(t, got, nil, 1, args...)
}

// True assert test value to be true
func True(t *testing.T, got interface{}, args ...interface{}) {
	equal(t, got, true, 1, args...)
}

// False assert test value to be false
func False(t *testing.T, got interface{}, args ...interface{}) {
	notEqual(t, got, true, 1, args...)
}

// Panic assert testing to be panic
func Panic(t *testing.T, fn func(), args ...interface{}) {
	defer func() {
		err := recover()
		if err == nil {
			_, file, line, ok := runtime.Caller(2)
			if ok {
				t.Errorf("%s:%d", file, line)
			}
			if len(args) > 0 {
				t.Error("! -", fmt.Sprint(args...))
			} else {
				t.Error("! -", "assert expected to be panic")
			}
		}
	}()

	fn()
}

// NotPanic assert testing to be panic
func NotPanic(t *testing.T, fn func(), args ...interface{}) {
	defer func() {
		err := recover()
		if err != nil {
			_, file, line, ok := runtime.Caller(2)
			if ok {
				t.Errorf("%s:%d", file, line)
			}
			if len(args) > 0 {
				t.Error("! -", fmt.Sprint(args...))
			} else {
				t.Error("! -", "assert expected to be not panic")
			}
		}
	}()

	fn()
}

func equal(t *testing.T, got, exp interface{}, step int, args ...interface{}) {
	fn := func() {
		switch got.(type) {
		case error:
			t.Errorf("! unexpected error: \"%s\"", got)
		default:
			t.Errorf("! expected %#v, but got %#v", exp, got)
		}
		if len(args) > 0 {
			t.Error("! -", fmt.Sprint(args...))
		}
	}
	ok := reflect.DeepEqual(exp, got)
	assert(t, ok, fn, step+1)
}

func notEqual(t *testing.T, got, exp interface{}, step int, args ...interface{}) {
	fn := func() {
		t.Errorf("! unexpected: %#v", got)
		if len(args) > 0 {
			t.Error("! -", fmt.Sprint(args...))
		}
	}
	ok := !reflect.DeepEqual(exp, got)
	assert(t, ok, fn, step+1)
}

func assert(t *testing.T, pass bool, fn func(), step int) {
	if !pass {
		_, file, line, ok := runtime.Caller(step + 1)
		if ok {
			t.Errorf("%s:%d", file, line)
		}
		fn()
		t.FailNow()
	}
}
