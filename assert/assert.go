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
	return "0.2.0"
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
func Equal(t *testing.T, exp, got interface{}, args ...interface{}) {
	fn := func() {
		t.Errorf("! expect %#v, but got %#v", exp, got)
		if len(args) > 0 {
			t.Error("! -", fmt.Sprint(args...))
		}
	}
	ok := reflect.DeepEqual(exp, got)
	assert(t, 1, ok, fn)
}

// NotEqual assert test value to be not equal
func NotEqual(t *testing.T, exp, got interface{}, args ...interface{}) {
	fn := func() {
		t.Errorf("! Unexpected: <%#v>", exp)
		if len(args) > 0 {
			t.Error("! -", fmt.Sprint(args...))
		}
	}
	ok := !reflect.DeepEqual(exp, got)
	assert(t, 1, ok, fn)
}

// Nil assert test value to be nil
func Nil(t *testing.T, got interface{}, args ...interface{}) {
	Equal(t, nil, got, args...)
}

// NotNil assert test value to be not nil
func NotNil(t *testing.T, got interface{}, args ...interface{}) {
	NotEqual(t, nil, got, args...)
}

// True assert test value to be true
func True(t *testing.T, got interface{}, args ...interface{}) {
	Equal(t, true, got, args...)
}

// False assert test value to be false
func False(t *testing.T, got interface{}, args ...interface{}) {
	NotEqual(t, true, got, args...)
}

func assert(t *testing.T, step int, result bool, f func()) {
	if !result {
		_, file, line, _ := runtime.Caller(step + 1)
		t.Errorf("%s:%d", file, line)
		f()
		t.FailNow()
	}
}
