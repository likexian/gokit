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
	"strings"
	"testing"

	"github.com/likexian/gokit/assert"
)

type a struct {
	x, y int
}

type b struct {
	x, y int
}

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestIsSlice(t *testing.T) {
	assert.False(t, IsSlice(0))
	assert.False(t, IsSlice("0"))
	assert.True(t, IsSlice([]int{0, 1, 2}))
	assert.True(t, IsSlice([]string{"0", "1", "2"}))
}

func TestUnique(t *testing.T) {
	// Not a slice
	tests := []struct {
		in  interface{}
		out interface{}
	}{
		{1, 1},
		{1.0, 1.0},
		{true, true},
	}

	for _, v := range tests {
		assert.Equal(t, Unique(v.in), v.out)
		assert.True(t, IsUnique(v.in))
	}

	// Is a slice
	tests = []struct {
		in  interface{}
		out interface{}
	}{
		{[]int{0, 0, 1, 1, 1, 2, 2, 3}, []int{0, 1, 2, 3}},
		{[]int8{0, 0, 1, 1, 1, 2, 2, 3}, []int8{0, 1, 2, 3}},
		{[]int16{0, 0, 1, 1, 1, 2, 2, 3}, []int16{0, 1, 2, 3}},
		{[]int32{0, 0, 1, 1, 1, 2, 2, 3}, []int32{0, 1, 2, 3}},
		{[]int64{0, 0, 1, 1, 1, 2, 2, 3}, []int64{0, 1, 2, 3}},
		{[]uint{0, 0, 1, 1, 1, 2, 2, 3}, []uint{0, 1, 2, 3}},
		{[]uint8{0, 0, 1, 1, 1, 2, 2, 3}, []uint8{0, 1, 2, 3}},
		{[]uint16{0, 0, 1, 1, 1, 2, 2, 3}, []uint16{0, 1, 2, 3}},
		{[]uint32{0, 0, 1, 1, 1, 2, 2, 3}, []uint32{0, 1, 2, 3}},
		{[]uint64{0, 0, 1, 1, 1, 2, 2, 3}, []uint64{0, 1, 2, 3}},
		{[]float32{0, 0, 1, 1, 1, 2, 2, 3}, []float32{0, 1, 2, 3}},
		{[]float64{0, 0, 1, 1, 1, 2, 2, 3}, []float64{0, 1, 2, 3}},
		{[]string{"a", "a", "b", "b", "b", "c"}, []string{"a", "b", "c"}},
		{[]bool{true, true, true, false}, []bool{true, false}},
		{[]interface{}{0, 1, 1, "1", 2}, []interface{}{0, 1, "1", 2}},
		{[]interface{}{[]int{0, 1}, []int{0, 1}, []int{1, 2}}, []interface{}{[]int{0, 1}, []int{1, 2}}},
		{[]interface{}{a{0, 1}, a{1, 2}, a{0, 1}, b{0, 1}}, []interface{}{a{0, 1}, a{1, 2}, b{0, 1}}},
	}

	for _, v := range tests {
		assert.Equal(t, Unique(v.in), v.out)
		assert.False(t, IsUnique(v.in))
		assert.True(t, IsUnique(v.out))
	}
}

func TestUniqueAppend(t *testing.T) {
	// Not a slice
	tests := []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{1, 1, 1},
		{1.0, 1.0, 1.0},
		{true, true, true},
	}

	for _, v := range tests {
		assert.Equal(t, UniqueAppend(v.x, v.y), v.out)
	}

	// Is a slice
	tests = []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{[]int{0, 1, 2, 3}, int(0), []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, int(4), []int{0, 1, 2, 3, 4}},
		{[]int8{0, 1, 2, 3}, int8(0), []int8{0, 1, 2, 3}},
		{[]int8{0, 1, 2, 3}, int8(4), []int8{0, 1, 2, 3, 4}},
		{[]int16{0, 1, 2, 3}, int16(0), []int16{0, 1, 2, 3}},
		{[]int16{0, 1, 2, 3}, int16(4), []int16{0, 1, 2, 3, 4}},
		{[]int32{0, 1, 2, 3}, int32(0), []int32{0, 1, 2, 3}},
		{[]int32{0, 1, 2, 3}, int32(4), []int32{0, 1, 2, 3, 4}},
		{[]int64{0, 1, 2, 3}, int64(0), []int64{0, 1, 2, 3}},
		{[]int64{0, 1, 2, 3}, int64(4), []int64{0, 1, 2, 3, 4}},
		{[]float32{0, 1, 2, 3}, float32(0), []float32{0, 1, 2, 3}},
		{[]float32{0, 1, 2, 3}, float32(4), []float32{0, 1, 2, 3, 4}},
		{[]float64{0, 1, 2, 3}, float64(0), []float64{0, 1, 2, 3}},
		{[]float64{0, 1, 2, 3}, float64(4), []float64{0, 1, 2, 3, 4}},
		{[]string{"a", "b", "c"}, "a", []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, "d", []string{"a", "b", "c", "d"}},
		{[]bool{true, false}, false, []bool{true, false}},
		{[]bool{true}, false, []bool{true, false}},
		{[]interface{}{0, 1, "1", 2}, 0, []interface{}{0, 1, "1", 2}},
		{[]interface{}{0, 1, "1", 2}, 3, []interface{}{0, 1, "1", 2, 3}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}}, []int{0, 1}, []interface{}{[]int{0, 1}, []int{1, 2}}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}}, []int{2, 3}, []interface{}{[]int{0, 1}, []int{1, 2}, []int{2, 3}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, a{0, 1}, []interface{}{a{0, 1}, a{1, 2}, b{0, 1}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, b{1, 2}, []interface{}{a{0, 1}, a{1, 2}, b{0, 1}, b{1, 2}}},
	}

	for _, v := range tests {
		assert.Equal(t, UniqueAppend(v.x, v.y), v.out)
	}

	// Append multi
	assert.Equal(t, UniqueAppend([]int{0, 1, 2, 3}, 2, 3), []int{0, 1, 2, 3})
	assert.Equal(t, UniqueAppend([]int{0, 1, 2, 3}, 2, 3, 4, 5), []int{0, 1, 2, 3, 4, 5})
}

func TestIntersect(t *testing.T) {
	// Not a slice
	tests := []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{1, 1, nil},
		{1.0, 1.0, nil},
		{true, true, nil},
		{[]int{1}, 1, nil},
		{[]float64{1.0}, 1, nil},
		{[]bool{true}, true, nil},
	}

	for _, v := range tests {
		assert.Equal(t, Intersect(v.x, v.y), v.out)
	}

	// Is a slice
	tests = []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{[]int{0, 1, 2}, []int{1, 2, 3}, []int{1, 2}},
		{[]int8{0, 1, 2}, []int8{1, 2, 3}, []int8{1, 2}},
		{[]int16{0, 1, 2}, []int16{1, 2, 3}, []int16{1, 2}},
		{[]int32{0, 1, 2}, []int32{1, 2, 3}, []int32{1, 2}},
		{[]int64{0, 1, 2}, []int64{1, 2, 3}, []int64{1, 2}},
		{[]float32{0, 1, 2}, []float32{1, 2, 3}, []float32{1, 2}},
		{[]float64{0, 1, 2}, []float64{1, 2, 3}, []float64{1, 2}},
		{[]string{"0", "1", "2"}, []string{"1", "2", "3"}, []string{"1", "2"}},
		{[]bool{true, false}, []bool{true}, []bool{true}},
		{[]interface{}{0, 1, "1", 2}, []interface{}{1, "1", 2, 3}, []interface{}{1, "1", 2}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}}, []interface{}{[]int{1, 2}, []int{2, 3}}, []interface{}{[]int{1, 2}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, []interface{}{a{1, 2}, b{2, 3}}, []interface{}{a{1, 2}}},
	}

	for _, v := range tests {
		assert.Equal(t, Intersect(v.x, v.y), v.out)
	}
}

func TestDifferent(t *testing.T) {
	// Not a slice
	tests := []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{1, 1, nil},
		{1.0, 1.0, nil},
		{true, true, nil},
		{[]int{1}, 1, nil},
		{[]float64{1.0}, 1, nil},
		{[]bool{true}, true, nil},
	}

	for _, v := range tests {
		assert.Equal(t, Different(v.x, v.y), v.out)
	}

	// Is a slice
	tests = []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{[]int{0, 1, 2}, []int{1, 2, 3}, []int{0}},
		{[]int8{0, 1, 2}, []int8{1, 2, 3}, []int8{0}},
		{[]int16{0, 1, 2}, []int16{1, 2, 3}, []int16{0}},
		{[]int32{0, 1, 2}, []int32{1, 2, 3}, []int32{0}},
		{[]int64{0, 1, 2}, []int64{1, 2, 3}, []int64{0}},
		{[]float32{0, 1, 2}, []float32{1, 2, 3}, []float32{0}},
		{[]float64{0, 1, 2}, []float64{1, 2, 3}, []float64{0}},
		{[]string{"0", "1", "2"}, []string{"1", "2", "3"}, []string{"0"}},
		{[]bool{true, false}, []bool{true}, []bool{false}},
		{[]interface{}{0, 1, "1", 2}, []interface{}{1, "1", 2, 3}, []interface{}{0}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}}, []interface{}{[]int{1, 2}, []int{2, 3}}, []interface{}{[]int{0, 1}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, []interface{}{a{1, 2}, b{2, 3}}, []interface{}{a{0, 1}, b{0, 1}}},
	}

	for _, v := range tests {
		assert.Equal(t, Different(v.x, v.y), v.out)
	}
}

func TestMerge(t *testing.T) {
	// Not a slice
	tests := []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{1, 1, 1},
		{1.0, 1.0, 1.0},
		{true, true, true},
		{[]int{1}, 1, []int{1}},
		{[]float64{1.0}, 1, []float64{1.0}},
		{[]bool{true}, true, []bool{true}},
	}

	for _, v := range tests {
		assert.Equal(t, Merge(v.x, v.y), v.out)
	}

	// Is a slice
	tests = []struct {
		x   interface{}
		y   interface{}
		out interface{}
	}{
		{[]int{0, 1, 2}, []int{1, 2, 3}, []int{0, 1, 2, 3}},
		{[]int8{0, 1, 2}, []int8{1, 2, 3}, []int8{0, 1, 2, 3}},
		{[]int16{0, 1, 2}, []int16{1, 2, 3}, []int16{0, 1, 2, 3}},
		{[]int32{0, 1, 2}, []int32{1, 2, 3}, []int32{0, 1, 2, 3}},
		{[]int64{0, 1, 2}, []int64{1, 2, 3}, []int64{0, 1, 2, 3}},
		{[]float32{0, 1, 2}, []float32{1, 2, 3}, []float32{0, 1, 2, 3}},
		{[]float64{0, 1, 2}, []float64{1, 2, 3}, []float64{0, 1, 2, 3}},
		{[]string{"0", "1", "2"}, []string{"1", "2", "3"}, []string{"0", "1", "2", "3"}},
		{[]bool{true, false}, []bool{true}, []bool{true, false}},
		{[]interface{}{0, 1, "1", 2}, []interface{}{1, "1", 2, 3}, []interface{}{0, 1, "1", 2, 3}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}}, []interface{}{[]int{1, 2}, []int{2, 3}}, []interface{}{[]int{0, 1}, []int{1, 2}, []int{2, 3}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, []interface{}{a{1, 2}, b{2, 3}}, []interface{}{a{0, 1}, a{1, 2}, b{0, 1}, b{2, 3}}},
	}

	for _, v := range tests {
		assert.Equal(t, Merge(v.x, v.y), v.out)
	}
}

func TestReverse(t *testing.T) {
	// Not a slice
	tests := []struct {
		in  interface{}
		out interface{}
	}{
		{1, 1},
		{1.0, 1.0},
		{true, true},
	}

	for _, v := range tests {
		assert.Equal(t, Reverse(v.in), v.out)
	}

	// Is a slice
	tests = []struct {
		in  interface{}
		out interface{}
	}{
		{[]int{0, 1, 2, 3, 4}, []int{4, 3, 2, 1, 0}},
		{[]int8{0, 1, 2, 3, 4}, []int8{4, 3, 2, 1, 0}},
		{[]int16{0, 1, 2, 3, 4}, []int16{4, 3, 2, 1, 0}},
		{[]int32{0, 1, 2, 3, 4}, []int32{4, 3, 2, 1, 0}},
		{[]int64{0, 1, 2, 3, 4}, []int64{4, 3, 2, 1, 0}},
		{[]float32{0, 1, 2, 3, 4}, []float32{4, 3, 2, 1, 0}},
		{[]float64{0, 1, 2, 3, 4}, []float64{4, 3, 2, 1, 0}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"e", "d", "c", "b", "a"}},
		{[]bool{true, false, true, false}, []bool{false, true, false, true}},
		{[]interface{}{0, 1, 2, "3", 3}, []interface{}{3, "3", 2, 1, 0}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}}, []interface{}{[]int{1, 2}, []int{0, 1}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, []interface{}{b{0, 1}, a{1, 2}, a{0, 1}}},
	}

	for _, v := range tests {
		assert.Equal(t, Reverse(v.in), v.out)
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		v   interface{}
		n   int
		out interface{}
	}{
		{1, -1, nil},
		{1, 0, nil},
		{1, 1, []int{1}},
		{1, 3, []int{1, 1, 1}},
		{int(1), 3, []int{1, 1, 1}},
		{int8(1), 3, []int8{1, 1, 1}},
		{int16(1), 3, []int16{1, 1, 1}},
		{int32(1), 3, []int32{1, 1, 1}},
		{int64(1), 3, []int64{1, 1, 1}},
		{float32(1), 3, []float32{1, 1, 1}},
		{float64(1), 3, []float64{1, 1, 1}},
		{"a", 3, []string{"a", "a", "a"}},
		{true, 3, []bool{true, true, true}},
		{[]int{1, 2}, 3, [][]int{{1, 2}, {1, 2}, {1, 2}}},
		{a{1, 2}, 3, []a{{1, 2}, {1, 2}, {1, 2}}},
		{[]interface{}{0, "1"}, 3, [][]interface{}{{0, "1"}, {0, "1"}, {0, "1"}}},
		{[]interface{}{[]int{0, 1}}, 3, [][]interface{}{{[]int{0, 1}}, {[]int{0, 1}}, {[]int{0, 1}}}},
		{[]interface{}{a{0, 1}}, 3, [][]interface{}{{a{x: 0, y: 1}}, {a{x: 0, y: 1}}, {a{x: 0, y: 1}}}},
	}

	for _, v := range tests {
		assert.Equal(t, Fill(v.v, v.n), v.out)
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		v   interface{}
		f   func(interface{}) bool
		out interface{}
	}{
		{1, func(v interface{}) bool { return v.(int) >= 0 }, 1},
		{[]int{-2, -1, 0, 1, 2}, func(v interface{}) bool { return v.(int) >= 0 }, []int{0, 1, 2}},
		{[]string{"a_0", "a_1", "b_1", "b_2", "a_2"}, func(v interface{}) bool { return strings.HasPrefix(v.(string), "a_") }, []string{"a_0", "a_1", "a_2"}},
	}

	for _, v := range tests {
		assert.Equal(t, Filter(v.v, v.f), v.out)
	}
}
