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
	return "0.13.2"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// IsSlice returns whether value is slice
func IsSlice(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Slice
}

// Unique returns unique values of slice
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

// IsUnique returns whether slice values is unique
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

// UniqueAppend append to slice if value x not exists in
func UniqueAppend(v interface{}, x ...interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	r := vv.Slice(0, vv.Len())
	for _, xx := range x {
		if !assert.IsContains(v, xx) {
			r = reflect.Append(r, reflect.ValueOf(xx))
		}
	}

	return r.Interface()
}

// Intersect returns values in both slices
func Intersect(x, y interface{}) interface{} {
	xx := reflect.ValueOf(x)
	if xx.Kind() != reflect.Slice {
		return nil
	}

	yy := reflect.ValueOf(y)
	if yy.Kind() != reflect.Slice {
		return nil
	}

	h := make(map[interface{}]bool)
	for i := 0; i < xx.Len(); i++ {
		h[hashValue(xx.Index(i).Interface())] = true
	}

	r := reflect.MakeSlice(reflect.TypeOf(x), 0, 0)
	for i := 0; i < yy.Len(); i++ {
		if _, ok := h[hashValue(yy.Index(i).Interface())]; ok {
			r = reflect.Append(r, yy.Index(i))
		}
	}

	return r.Interface()
}

// Different returns values in x but not in y
func Different(x, y interface{}) interface{} {
	xx := reflect.ValueOf(x)
	if xx.Kind() != reflect.Slice {
		return nil
	}

	yy := reflect.ValueOf(y)
	if yy.Kind() != reflect.Slice {
		return nil
	}

	h := make(map[interface{}]bool)
	for i := 0; i < yy.Len(); i++ {
		h[hashValue(yy.Index(i).Interface())] = true
	}

	r := reflect.MakeSlice(reflect.TypeOf(x), 0, 0)
	for i := 0; i < xx.Len(); i++ {
		if _, ok := h[hashValue(xx.Index(i).Interface())]; !ok {
			r = reflect.Append(r, xx.Index(i))
		}
	}

	return r.Interface()
}

// Merge append y values to x if not exists in
func Merge(x, y interface{}) interface{} {
	xx := reflect.ValueOf(x)
	if xx.Kind() != reflect.Slice {
		return x
	}

	yy := reflect.ValueOf(y)
	if yy.Kind() != reflect.Slice {
		return x
	}

	h := make(map[interface{}]bool)
	for i := 0; i < xx.Len(); i++ {
		h[hashValue(xx.Index(i).Interface())] = true
	}

	r := xx.Slice(0, xx.Len())
	for i := 0; i < yy.Len(); i++ {
		if _, ok := h[hashValue(yy.Index(i).Interface())]; !ok {
			r = reflect.Append(r, yy.Index(i))
		}
	}

	return r.Interface()
}

// Reverse returns a slice with elements in reverse order
func Reverse(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	swap := reflect.Swapper(v)
	for i, j := 0, vv.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}

	return v
}

// Fill returns a slice with count number of v values
func Fill(v interface{}, count int) interface{} {
	if count <= 0 {
		return nil
	}

	r := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(v)), 0, 0)
	for i := 0; i < count; i++ {
		r = reflect.Append(r, reflect.ValueOf(v))
	}

	return r.Interface()
}

// Filter filter slice values usig callback fnction fn
func Filter(v interface{}, fn interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	if fn != nil {
		err := CheckIsFunc(fn, 1, 1)
		if err != nil {
			panic("Filter: " + err.Error())
		}
	} else {
		fn = func(v interface{}) bool {
			return v != nil
		}
	}

	fv := reflect.ValueOf(fn)
	ot := fv.Type().Out(0).Kind()
	if ot != reflect.Bool {
		panic("Filter: fn expected to return a bool but got " + ot.String())
	}

	r := reflect.MakeSlice(reflect.TypeOf(v), 0, 0)
	for i := 0; i < vv.Len(); i++ {
		if fv.Call([]reflect.Value{vv.Index(i)})[0].Interface().(bool) {
			r = reflect.Append(r, vv.Index(i))
		}
	}

	return r.Interface()
}

// CheckIsFunc check if fn is a function with n[0] arguments and n[1] returns
func CheckIsFunc(fn interface{}, n ...int) error {
	ft := reflect.TypeOf(fn)
	if ft.Kind() != reflect.Func {
		return fmt.Errorf("fn is not a function")
	}

	if len(n) >= 1 && n[0] != ft.NumIn() {
		return fmt.Errorf("fn expected to have %d arguments but got %d", n[0], ft.NumIn())
	}

	if len(n) >= 2 && n[1] != ft.NumOut() {
		return fmt.Errorf("fn expected to have %d returns but got %d", n[1], ft.NumOut())
	}

	return nil
}

// hashValue returns a hashable value
func hashValue(x interface{}) interface{} {
	return fmt.Sprintf("%#v", x)
}
