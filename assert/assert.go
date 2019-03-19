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
	return "0.7.0"
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

// Empty assert test value to be empty
func Empty(t *testing.T, got interface{}, args ...interface{}) {
	equal(t, IsEmpty(got), true, 1, args...)
}

// Empty assert test value to be empty
func NotEmpty(t *testing.T, got interface{}, args ...interface{}) {
	notEqual(t, IsEmpty(got), true, 1, args...)
}

// Zero assert test value to be zero
func Zero(t *testing.T, got interface{}, args ...interface{}) {
	equal(t, IsZero(got), true, 1, args...)
}

// Empty assert test value to be zero
func NotZero(t *testing.T, got interface{}, args ...interface{}) {
	notEqual(t, IsZero(got), true, 1, args...)
}

// Panic assert testing to be panic
func Panic(t *testing.T, fn func(), args ...interface{}) {
	defer func() {
		ff := func() {
			t.Error("! -", "assert expected to be panic")
			if len(args) > 0 {
				t.Error("! -", fmt.Sprint(args...))
			}
		}
		ok := recover() != nil
		assert(t, ok, ff, 2)
	}()

	fn()
}

// NotPanic assert testing to be panic
func NotPanic(t *testing.T, fn func(), args ...interface{}) {
	defer func() {
		ff := func() {
			t.Error("! -", "assert expected to be not panic")
			if len(args) > 0 {
				t.Error("! -", fmt.Sprint(args...))
			}
		}
		ok := recover() == nil
		assert(t, ok, ff, 3)
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

// IsEmpty returns value is empty
func IsEmpty(v interface{}) bool {
	vv := reflect.ValueOf(v)
	if !vv.IsValid() {
		return true
	}

	switch vv.Kind() {
	case reflect.Ptr, reflect.Interface:
		return IsEmpty(vv.Elem())
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return vv.Len() == 0
	default:
		return false
	}
}

// IsZero returns value is zero
func IsZero(v interface{}) bool {
	vv := reflect.ValueOf(v)
	if !vv.IsValid() {
		return true
	}

	switch vv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return vv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return vv.Float() == 0
	default:
		return false
	}
}
