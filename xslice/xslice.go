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
	"fmt"
	"reflect"

	"github.com/likexian/gokit/assert"
)

// Version returns package version
func Version() string {
	return "0.8.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Unique returns an unique slice
func Unique(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	r := reflect.MakeSlice(reflect.TypeOf(v), 0, vv.Cap())
	for i := 0; i < vv.Len(); i++ {
		if !assert.IsContains(r.Interface(), vv.Index(i).Interface()) {
			r = reflect.Append(r, vv.Index(i))
		}
	}

	return r.Interface()
}

// IsUnique returns whether slice is unique
func IsUnique(v interface{}) bool {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return true
	}

	if vv.Len() <= 1 {
		return true
	}

	x := vv.Index(0)
	y := vv.Slice(1, vv.Len())

	if assert.IsContains(y.Interface(), x.Interface()) {
		return false
	}

	return IsUnique(y.Interface())
}

// UniqueAppend append to slice if not exists
func UniqueAppend(v interface{}, x interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	r := vv.Slice(0, vv.Len())
	if !assert.IsContains(v, x) {
		r = reflect.Append(r, reflect.ValueOf(x))
	}

	return r.Interface()
}

// Intersect returns intersection of two slice
func Intersect(x, y interface{}) interface{} {
	xx := reflect.ValueOf(x)
	if xx.Kind() != reflect.Slice {
		return nil
	}

	yy := reflect.ValueOf(y)
	if yy.Kind() != reflect.Slice {
		return nil
	}

	hash := func(x interface{}) interface{} {
		return fmt.Sprintf("%#v", x)
	}

	h := make(map[interface{}]bool)
	for i := 0; i < xx.Len(); i++ {
		h[hash(xx.Index(i).Interface())] = true
	}

	r := reflect.MakeSlice(reflect.TypeOf(x), 0, 0)
	for i := 0; i < yy.Len(); i++ {
		if _, ok := h[hash(yy.Index(i).Interface())]; ok {
			r = reflect.Append(r, yy.Index(i))
		}
	}

	return r.Interface()
}
