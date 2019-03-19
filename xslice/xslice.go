/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xslice

import (
	"reflect"
	"strings"
)

// Version returns package version
func Version() string {
	return "0.3.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// Contains returns whether value is within array
func Contains(array interface{}, value interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		s := reflect.ValueOf(array).MapKeys()
		for i := 0; i < len(s); i++ {
			if reflect.DeepEqual(value, s[i].Interface()) {
				return true
			}
		}
	case reflect.String:
		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			return strings.Contains(array.(string), value.(string))
		default:
			return false
		}
	default:
		panic("not support data type")
	}

	return false
}

// Unique returns unique processed array
func Unique(array interface{}) interface{} {
	switch array.(type) {
	case []int:
		result := []int{}
		for _, v := range array.([]int) {
			if !Contains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []int64:
		result := []int64{}
		for _, v := range array.([]int64) {
			if !Contains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []uint64:
		result := []uint64{}
		for _, v := range array.([]uint64) {
			if !Contains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []float64:
		result := []float64{}
		for _, v := range array.([]float64) {
			if !Contains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []string:
		result := []string{}
		for _, v := range array.([]string) {
			if !Contains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []bool:
		result := []bool{}
		for _, v := range array.([]bool) {
			if !Contains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []interface{}:
		result := []interface{}{}
		for _, v := range array.([]interface{}) {
			if !Contains(result, v) {
				result = append(result, v)
			}
		}
		return result
	default:
		panic("not support data type")
	}
}
