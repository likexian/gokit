/*
 * Copyright 2012-2019 Li Kexian
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
	return "Licensed under the Apache License 2.0"
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
