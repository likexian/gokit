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
	"math"
	"math/rand"
	"reflect"

	"github.com/likexian/gokit/assert"
)

// Version returns package version
func Version() string {
	return "0.18.0"
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
func Reverse(v interface{}) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return
	}

	swap := reflect.Swapper(v)
	for i, j := 0, vv.Len()-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// Shuffle shuffle a slice
func Shuffle(v interface{}) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return
	}

	swap := reflect.Swapper(v)
	for i := vv.Len() - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
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

// Chunk split slice into chunks
func Chunk(v interface{}, size int) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	if size < 1 {
		panic("Chunk: size less than 1")
	}

	n := int(math.Ceil(float64(vv.Len()) / float64(size)))
	r := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(v)), 0, 0)
	for i := 0; i < n; i++ {
		rr := reflect.MakeSlice(reflect.TypeOf(v), 0, 0)
		for j := 0; j < size; j++ {
			if i*size+j >= vv.Len() {
				break
			}
			rr = reflect.Append(rr, vv.Index(i*size+j))
		}
		r = reflect.Append(r, rr)
	}

	return r.Interface()
}

// Concat returns a new flatten slice of []slice
func Concat(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	if vv.Len() == 0 {
		return v
	}

	vt := reflect.TypeOf(v)
	if vt.Elem().Kind() != reflect.Slice {
		return v
	}

	r := reflect.MakeSlice(reflect.TypeOf(v).Elem(), 0, 0)
	for i := 0; i < vv.Len(); i++ {
		for j := 0; j < vv.Index(i).Len(); j++ {
			r = reflect.Append(r, vv.Index(i).Index(j))
		}
	}

	return r.Interface()
}

// Filter filter slice values using callback function fn
func Filter(v interface{}, fn interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	err := CheckIsFunc(fn, 1, 1)
	if err != nil {
		panic("Filter: " + err.Error())
	}

	fv := reflect.ValueOf(fn)
	ot := fv.Type().Out(0)
	if ot.Kind() != reflect.Bool {
		panic("Filter: fn expected to return bool but got " + ot.Kind().String())
	}

	r := reflect.MakeSlice(reflect.TypeOf(v), 0, 0)
	for i := 0; i < vv.Len(); i++ {
		if fv.Call([]reflect.Value{vv.Index(i)})[0].Interface().(bool) {
			r = reflect.Append(r, vv.Index(i))
		}
	}

	return r.Interface()
}

// Map apply callback function fn to elements of slice
func Map(v interface{}, fn interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	err := CheckIsFunc(fn, 1, 1)
	if err != nil {
		panic("Map: " + err.Error())
	}

	fv := reflect.ValueOf(fn)
	ot := fv.Type().Out(0)

	r := reflect.MakeSlice(reflect.SliceOf(ot), 0, 0)
	for i := 0; i < vv.Len(); i++ {
		r = reflect.Append(r, fv.Call([]reflect.Value{vv.Index(i)})[0])
	}

	return r.Interface()
}

// Reduce reduce the slice values using callback function fn
func Reduce(v interface{}, fn interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return v
	}

	if vv.Len() == 0 {
		panic("Reduce: slice is empty")
	}

	err := CheckIsFunc(fn, 2, 1)
	if err != nil {
		panic("Reduce: " + err.Error())
	}

	fv := reflect.ValueOf(fn)
	if vv.Type().Elem() != fv.Type().In(0) || vv.Type().Elem() != fv.Type().In(1) {
		panic(fmt.Sprintf("Reduce: fn expected to have (%s, %s) arguments but got (%s, %s)",
			vv.Type().Elem(), vv.Type().Elem(), fv.Type().In(0), fv.Type().In(1)))
	}

	if vv.Type().Elem() != fv.Type().Out(0) {
		panic(fmt.Sprintf("Reduce: fn expected to return %s but got %s",
			vv.Type().Elem(), fv.Type().Out(0).String()))
	}

	r := vv.Index(0)
	for i := 1; i < vv.Len(); i++ {
		r = fv.Call([]reflect.Value{r, vv.Index(i)})[0]
	}

	return r.Interface()
}

// CheckIsFunc check if fn is a function with n[0] arguments and n[1] returns
func CheckIsFunc(fn interface{}, n ...int) error {
	if fn == nil {
		return fmt.Errorf("fn excepted to be a function but got nil")
	}

	ft := reflect.TypeOf(fn)
	if ft.Kind() != reflect.Func {
		return fmt.Errorf("fn excepted to be a function but got " + ft.Kind().String())
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
