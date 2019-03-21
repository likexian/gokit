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
	"github.com/likexian/gokit/assert"
)

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
	return "Apache License, Version 2.0"
}

// Unique returns unique processed array
func Unique(array interface{}) interface{} {
	switch array.(type) {
	case []int:
		result := []int{}
		for _, v := range array.([]int) {
			if !assert.IsContains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []int64:
		result := []int64{}
		for _, v := range array.([]int64) {
			if !assert.IsContains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []uint64:
		result := []uint64{}
		for _, v := range array.([]uint64) {
			if !assert.IsContains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []float64:
		result := []float64{}
		for _, v := range array.([]float64) {
			if !assert.IsContains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []string:
		result := []string{}
		for _, v := range array.([]string) {
			if !assert.IsContains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []bool:
		result := []bool{}
		for _, v := range array.([]bool) {
			if !assert.IsContains(result, v) {
				result = append(result, v)
			}
		}
		return result
	case []interface{}:
		result := []interface{}{}
		for _, v := range array.([]interface{}) {
			if !assert.IsContains(result, v) {
				result = append(result, v)
			}
		}
		return result
	default:
		return array
	}
}
