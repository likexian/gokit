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
	"strings"
)

// IsZero returns value is zero value
func IsZero(v interface{}) bool {
	vv := reflect.ValueOf(v)
	switch vv.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Bool:
		return !vv.Bool()
	case reflect.Ptr, reflect.Interface:
		return vv.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return vv.Len() == 0
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

// IsContains returns whether value is within array
func IsContains(array interface{}, value interface{}) bool {
	vv := reflect.ValueOf(array)
	switch vv.Kind() {
	case reflect.Invalid:
		return false
	case reflect.Slice:
		for i := 0; i < vv.Len(); i++ {
			if reflect.DeepEqual(value, vv.Index(i).Interface()) {
				return true
			}
		}
		return false
	case reflect.Map:
		s := vv.MapKeys()
		for i := 0; i < len(s); i++ {
			if reflect.DeepEqual(value, s[i].Interface()) {
				return true
			}
		}
		return false
	case reflect.String:
		ss := reflect.ValueOf(value)
		switch ss.Kind() {
		case reflect.String:
			return strings.Contains(vv.String(), ss.String())
		}
		return false
	default:
		return reflect.DeepEqual(array, value)
	}
}

// VLen returns length of value
func VLen(v interface{}) int {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr || vv.Kind() == reflect.Interface {
		if vv.IsNil() {
			return 0
		} else {
			vv = vv.Elem()
		}
	}

	switch vv.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Ptr, reflect.Interface:
		return 0
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return vv.Len()
	default:
		return len(fmt.Sprintf("%#v", v))
	}
}
