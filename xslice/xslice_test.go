/*
 * Copyright 2012-2021 Li Kexian
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

func TestIndex(t *testing.T) {
	// Not a slice
	tests := []struct {
		x interface{}
		y interface{}
		z int
	}{
		{1, 1, -1},
		{1.0, 1.0, -1},
		{true, true, -1},
	}

	for _, v := range tests {
		assert.Panic(t, func() { Index(v.x, v.y) })
	}

	// Is a slice
	tests = []struct {
		x interface{}
		y interface{}
		z int
	}{
		{[]int{0, 0, 1, 1, 1, 2, 2, 3}, int(0), 0},
		{[]int8{0, 0, 1, 1, 1, 2, 2, 3}, int8(1), 2},
		{[]int16{0, 0, 1, 1, 1, 2, 2, 3}, int16(2), 5},
		{[]int32{0, 0, 1, 1, 1, 2, 2, 3}, int32(3), 7},
		{[]int64{0, 0, 1, 1, 1, 2, 2, 3}, int64(4), -1},
		{[]uint{0, 0, 1, 1, 1, 2, 2, 3}, uint(0), 0},
		{[]uint8{0, 0, 1, 1, 1, 2, 2, 3}, uint8(1), 2},
		{[]uint16{0, 0, 1, 1, 1, 2, 2, 3}, uint16(2), 5},
		{[]uint32{0, 0, 1, 1, 1, 2, 2, 3}, uint32(3), 7},
		{[]uint64{0, 0, 1, 1, 1, 2, 2, 3}, uint64(4), -1},
		{[]float32{0, 0, 1, 1, 1, 2, 2, 3}, float32(0), 0},
		{[]float64{0, 0, 1, 1, 1, 2, 2, 3}, float64(4), -1},
		{[]string{"a", "a", "b", "b", "b", "c"}, "a", 0},
		{[]bool{true, true, true, false}, false, 3},
		{[]interface{}{0, 1, 1, "1", 2}, "1", 3},
		{[]interface{}{[]int{0, 1}, []int{0, 1}, []int{1, 2}}, []int{1, 2}, 2},
		{[]interface{}{a{0, 1}, a{1, 2}, a{0, 1}, b{0, 1}}, a{0, 1}, 0},
	}

	for _, v := range tests {
		assert.Equal(t, Index(v.x, v.y), v.z)
	}
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
		assert.Panic(t, func() { Unique(v.in) })
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
	}
}

func TestIsUnique(t *testing.T) {
	// Not a slice
	tests := []struct {
		in interface{}
	}{
		{1},
		{1.0},
		{true},
	}

	for _, v := range tests {
		assert.Panic(t, func() { IsUnique(v.in) })
	}

	// Is a slice
	tests = []struct {
		in interface{}
	}{
		{[]int{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]int8{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]int16{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]int32{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]int64{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]uint{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]uint8{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]uint16{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]uint32{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]uint64{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]float32{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]float64{0, 0, 1, 1, 1, 2, 2, 3}},
		{[]string{"a", "a", "b", "b", "b", "c"}},
		{[]bool{true, true, true, false}},
		{[]interface{}{0, 1, 1, "1", 2}},
		{[]interface{}{[]int{0, 1}, []int{0, 1}, []int{1, 2}}},
		{[]interface{}{a{0, 1}, a{1, 2}, a{0, 1}, b{0, 1}}},
	}

	for _, v := range tests {
		assert.False(t, IsUnique(v.in))
	}

	// Is a slice
	tests = []struct {
		in interface{}
	}{
		{[]int{1}},
		{[]int{0, 1, 2, 3}},
		{[]int8{0, 1, 2, 3}},
		{[]int16{0, 1, 2, 3}},
		{[]int32{0, 1, 2, 3}},
		{[]int64{0, 1, 2, 3}},
		{[]uint{0, 1, 2, 3}},
		{[]uint8{0, 1, 2, 3}},
		{[]uint16{0, 1, 2, 3}},
		{[]uint32{0, 1, 2, 3}},
		{[]uint64{0, 1, 2, 3}},
		{[]float32{0, 1, 2, 3}},
		{[]float64{0, 1, 2, 3}},
		{[]string{"a", "b", "c"}},
		{[]bool{true, false}},
		{[]interface{}{0, 1, "1", 2}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}},
	}

	for _, v := range tests {
		assert.True(t, IsUnique(v.in))
	}
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
		assert.Panic(t, func() { Intersect(v.x, v.y) })
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
		assert.Panic(t, func() { Different(v.x, v.y) })
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
		assert.Panic(t, func() { Merge(v.x, v.y) })
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
		{[]interface{}{[]int{0, 1}, []int{1, 2}}, []interface{}{[]int{1, 2},
			[]int{2, 3}}, []interface{}{[]int{0, 1}, []int{1, 2}, []int{2, 3}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, []interface{}{a{1, 2}, b{2, 3}},
			[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}, b{2, 3}}},
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
		assert.Panic(t, func() { Reverse(v.in) })
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
		Reverse(v.in)
		assert.Equal(t, v.in, v.out)
	}
}

func TestShuffle(t *testing.T) {
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
		assert.Panic(t, func() { Shuffle(v.in) })
	}

	// Is a slice
	tests = []struct {
		in  interface{}
		out interface{}
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int8{0, 1, 2, 3, 4}, []int8{0, 1, 2, 3, 4}},
		{[]int16{0, 1, 2, 3, 4}, []int16{0, 1, 2, 3, 4}},
		{[]int32{0, 1, 2, 3, 4}, []int32{0, 1, 2, 3, 4}},
		{[]int64{0, 1, 2, 3, 4}, []int64{0, 1, 2, 3, 4}},
		{[]float32{0, 1, 2, 3, 4}, []float32{0, 1, 2, 3, 4}},
		{[]float64{0, 1, 2, 3, 4}, []float64{0, 1, 2, 3, 4}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c", "d", "e"}},
		{[]bool{true, false, false, true, true}, []bool{true, false, false, true, true}},
		{[]interface{}{0, 1, 2, "3", 3}, []interface{}{0, 1, 2, "3", 3}},
		{[]interface{}{[]int{0, 1}, []int{1, 2}, []int{1, 2}}, []interface{}{[]int{0, 1}, []int{1, 2}, []int{1, 2}}},
		{[]interface{}{a{0, 1}, a{1, 2}, b{0, 1}}, []interface{}{a{0, 1}, a{1, 2}, b{0, 1}}},
	}

	for _, v := range tests {
		Shuffle(v.in)
		assert.NotEqual(t, v.in, v.out)
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
	}

	for _, v := range tests {
		assert.Panic(t, func() { Fill(v.v, v.n) })
	}

	tests = []struct {
		v   interface{}
		n   int
		out interface{}
	}{
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

func TestChunk(t *testing.T) {
	tests := []struct {
		v   interface{}
		n   int
		out interface{}
	}{
		{1, 1, 1},
		{[]int{1}, 0, nil},
	}

	for _, v := range tests {
		assert.Panic(t, func() { Chunk(v.v, v.n) })
	}

	tests = []struct {
		v   interface{}
		n   int
		out interface{}
	}{
		{[]int{0, 1, 2}, 1, [][]int{{0}, {1}, {2}}},
		{[]int{0, 1, 2, 3, 4}, 2, [][]int{{0, 1}, {2, 3}, {4}}},
		{[]int{0, 1, 2, 3, 4, 5}, 2, [][]int{{0, 1}, {2, 3}, {4, 5}}},
		{[]string{"a", "b", "c", "d", "e"}, 3, [][]string{{"a", "b", "c"}, {"d", "e"}}},
		{[]interface{}{a{0, 1}, b{2, 3}, a{4, 5}}, 2, [][]interface{}{{a{0, 1}, b{2, 3}}, {a{4, 5}}}},
	}

	for _, v := range tests {
		assert.Equal(t, Chunk(v.v, v.n), v.out)
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		in  interface{}
		out interface{}
	}{
		{1, 1},
	}

	for _, v := range tests {
		assert.Panic(t, func() { Concat(v.in) })
	}

	tests = []struct {
		in  interface{}
		out interface{}
	}{
		{[]int{}, []int{}},
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[][]int{{0, 1}, {2, 3}, {4}}, []int{0, 1, 2, 3, 4}},
		{[][]string{{"a", "b"}, {"c"}, {"d", "e"}}, []string{"a", "b", "c", "d", "e"}},
		{[][]interface{}{{a{0, 1}, b{0, 1}}, {a{1, 2}}}, []interface{}{a{0, 1}, b{0, 1}, a{1, 2}}},
	}

	for _, v := range tests {
		assert.Equal(t, Concat(v.in), v.out)
	}
}

func TestFilter(t *testing.T) {
	// Panic tests
	tests := []struct {
		v   interface{}
		f   interface{}
		out interface{}
	}{
		{1, nil, 1},
		{[]int{1}, nil, nil},
		{[]int{1}, 1, nil},
		{[]int{1}, func() {}, nil},
		{[]int{1}, func(v int) {}, nil},
		{[]int{1}, func(v int) int { return v }, nil},
	}

	for _, v := range tests {
		assert.Panic(t, func() { Filter(v.v, v.f) })
	}

	// General tests
	tests = []struct {
		v   interface{}
		f   interface{}
		out interface{}
	}{
		{[]interface{}{0, 1, nil, 2}, func(v interface{}) bool { return v != nil }, []interface{}{0, 1, 2}},
		{[]int{-2, -1, 0, 1, 2}, func(v int) bool { return v >= 0 }, []int{0, 1, 2}},
		{[]string{"a_0", "b_1", "a_1"}, func(v string) bool { return strings.HasPrefix(v, "a_") }, []string{"a_0", "a_1"}},
		{[]bool{true, false, false}, func(v bool) bool { return !v }, []bool{false, false}},
	}

	for _, v := range tests {
		assert.Equal(t, Filter(v.v, v.f), v.out)
	}
}

func TestMap(t *testing.T) {
	// Panic tests
	tests := []struct {
		v   interface{}
		f   interface{}
		out interface{}
	}{
		{1, nil, 1},
		{[]int{1}, nil, nil},
		{[]int{1}, 1, nil},
		{[]int{1}, func() {}, nil},
		{[]int{1}, func(v int) {}, nil},
	}

	for _, v := range tests {
		assert.Panic(t, func() { Map(v.v, v.f) })
	}

	// General tests
	tests = []struct {
		v   interface{}
		f   interface{}
		out interface{}
	}{
		{[]int{1, 2, 3, 4, 5}, func(v int) int { return v * v * v }, []int{1, 8, 27, 64, 125}},
		{[]int{-2, -1, 0, 1, 2}, func(v int) bool { return v > 0 }, []bool{false, false, false, true, true}},
		{[]string{"a", "b", "c"}, func(v string) string { return "x_" + v }, []string{"x_a", "x_b", "x_c"}},
		{[]bool{true, false, false}, func(v bool) bool { return !v }, []bool{false, true, true}},
		{[]interface{}{1, nil}, func(v interface{}) interface{} { return assert.If(v == nil, -1, v) }, []interface{}{1, -1}},
	}

	for _, v := range tests {
		assert.Equal(t, Map(v.v, v.f), v.out)
	}
}

func TestReduce(t *testing.T) {
	// Panic tests
	tests := []struct {
		v   interface{}
		f   interface{}
		out interface{}
	}{
		{1, nil, 1},
		{[]int{}, nil, nil},
		{[]int{0, 1}, nil, nil},
		{[]int{0, 1}, 1, nil},
		{[]int{0, 1}, func() {}, nil},
		{[]int{0, 1}, func(x int) {}, nil},
		{[]int{0, 1}, func(x, y int) {}, nil},
		{[]int{0, 1}, func(x bool, y int) int { return y }, nil},
		{[]int{0, 1}, func(x int, y bool) int { return x }, nil},
		{[]int{0, 1}, func(x int, y int) bool { return true }, nil},
	}

	for _, v := range tests {
		assert.Panic(t, func() { Reduce(v.v, v.f) })
	}

	// General tests
	tests = []struct {
		v   interface{}
		f   interface{}
		out interface{}
	}{
		{[]int{1}, func(x, y int) int { return x + y }, 1},
		{[]int{1, 2}, func(x, y int) int { return x + y }, 3},
		{[]int{1, 2, 3, 4}, func(x, y int) int { return x * y }, 24},
	}

	for _, v := range tests {
		assert.Equal(t, Reduce(v.v, v.f).(int), v.out)
	}
}
