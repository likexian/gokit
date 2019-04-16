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

package assert

import (
	"regexp"
	"testing"
)

func TestIsZero(t *testing.T) {
	var i interface{}
	tests := []interface{}{
		i,
		"",
		false,
		[]byte{},
		[]int{},
		[]string{},
		map[string]int{},
		map[string]string{},
		map[string]interface{}{},
		0,
		int(0),
		int8(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint32(0),
		uint64(0),
		float32(0),
		float64(0),
	}

	for _, v := range tests {
		True(t, IsZero(v))
	}

	i = "a"
	tests = []interface{}{
		i,
		&i,
		"a",
		true,
		[]byte{0},
		[]int{0},
		[]string{"a"},
		map[string]int{"a": 1},
		map[string]string{"a": ""},
		map[string]interface{}{"a": "b"},
		1,
		int(1),
		int8(1),
		int32(1),
		int64(1),
		uint(1),
		uint8(1),
		uint32(1),
		uint64(1),
		float32(0.1),
		float64(0.1),
		struct{ x int }{1},
	}

	for _, v := range tests {
		False(t, IsZero(v))
	}
}

func TestIsContains(t *testing.T) {
	var s *string
	var i interface{} = s
	tests := [][]interface{}{
		[]interface{}{nil, nil, false},
		[]interface{}{s, s, false},
		[]interface{}{&i, &i, true},

		[]interface{}{[]int{0, 1, 2}, 1, true},
		[]interface{}{[]int{0, 1, 2}, 3, false},
		[]interface{}{[]int{0, 1, 2}, int64(1), false},
		[]interface{}{[]int{0, 1, 2}, "1", false},
		[]interface{}{[]int{0, 1, 2}, true, false},

		[]interface{}{[]int64{0, 1, 2}, int64(1), true},
		[]interface{}{[]int64{0, 1, 2}, int64(3), false},
		[]interface{}{[]int64{0, 1, 2}, 1, false},
		[]interface{}{[]int64{0, 1, 2}, "1", false},
		[]interface{}{[]int64{0, 1, 2}, true, false},

		[]interface{}{[]float64{0.0, 1.0, 2.0}, 1.0, true},
		[]interface{}{[]float64{0.0, 1.0, 2.0}, float64(1), true},
		[]interface{}{[]float64{0.0, 1.0, 2.0}, 3.0, false},
		[]interface{}{[]float64{0.0, 1.0, 2.0}, 1, false},

		[]interface{}{[]string{"a", "b", "c"}, "a", true},
		[]interface{}{[]string{"a", "b", "c"}, "d", false},
		[]interface{}{[]string{"a", "b", "c"}, 1, false},
		[]interface{}{[]string{"a", "b", "c"}, true, false},

		[]interface{}{[]interface{}{0, "1", 2}, "1", true},
		[]interface{}{[]interface{}{0, "1", 2}, 1, false},
		[]interface{}{[]interface{}{0, 1, 2}, true, false},
		[]interface{}{[]interface{}{0, true, 2}, true, true},
		[]interface{}{[]interface{}{0, false, 2}, true, false},

		[]interface{}{[]interface{}{[]int{0, 1}, []int{1, 2}}, []int{1, 2}, true},
		[]interface{}{[]interface{}{[]int{0, 1}, []int{1, 2, 3}}, []int{1, 2}, false},

		[]interface{}{[]A{A{0, 1}, A{1, 2}, A{1, 3}}, A{1, 2}, true},
		[]interface{}{[]interface{}{A{0, 1}, B{1, 2}, A{1, 3}}, B{1, 2}, true},
		[]interface{}{[]interface{}{A{0, 1}, B{1, 2}, A{1, 3}}, A{1, 2}, false},

		[]interface{}{map[string]int{"a": 1}, "a", true},
		[]interface{}{map[string]int{"a": 1}, "d", false},
		[]interface{}{map[string]int{"a": 1}, 1, false},
		[]interface{}{map[string]int{"a": 1}, true, false},

		[]interface{}{"abc", "a", true},
		[]interface{}{"abc", "d", false},
		[]interface{}{"abc", 1, false},
		[]interface{}{"abc", true, false},

		[]interface{}{"a", "a", true},
		[]interface{}{1, 1, true},
		[]interface{}{-1, -1, true},
		[]interface{}{1.0, 1.0, true},
		[]interface{}{true, true, true},
		[]interface{}{false, false, true},
	}

	for _, v := range tests {
		Equal(t, IsContains(v[0], v[1]), v[2])
	}
}

func TestIsMatch(t *testing.T) {
	var i interface{}
	tests := []struct {
		r   interface{}
		v   interface{}
		out bool
	}{
		{regexp.MustCompile("v\\d+"), "v100", true},
		{"v\\d+", "v100", true},
		{"\\d+\\.\\d+", 100.1, true},
		{regexp.MustCompile("v\\d+"), "x100", false},
		{"v\\d+", "x100", false},
		{"\\d+\\.\\d+", "x100", false},
		{"v\\d+", i, false},
		{i, 100.1, false},
	}

	for _, v := range tests {
		vv := IsMatch(v.r, v.v)
		Equal(t, vv, v.out, v)
	}
}

func TestLength(t *testing.T) {
	var s *string
	var i interface{} = s
	tests := [][]interface{}{
		[]interface{}{nil, 0},
		[]interface{}{s, 0},
		[]interface{}{&i, 0},
		[]interface{}{"", 0},
		[]interface{}{"1", 1},
		[]interface{}{true, 4},
		[]interface{}{false, 5},
		[]interface{}{[]byte{}, 0},
		[]interface{}{[]byte{0}, 1},
		[]interface{}{[]int{}, 0},
		[]interface{}{[]int{0}, 1},
		[]interface{}{[]string{}, 0},
		[]interface{}{[]string{"a"}, 1},
		[]interface{}{map[string]int{}, 0},
		[]interface{}{map[string]int{"a": 1}, 1},
		[]interface{}{map[string]string{}, 0},
		[]interface{}{map[string]string{"a": "b"}, 1},
		[]interface{}{map[string]interface{}{}, 0},
		[]interface{}{map[string]interface{}{"a": i}, 1},
		[]interface{}{0, 1},
		[]interface{}{1, 1},
		[]interface{}{int(0), 1},
		[]interface{}{int(1), 1},
		[]interface{}{int8(0), 1},
		[]interface{}{int8(1), 1},
		[]interface{}{int32(0), 1},
		[]interface{}{int32(1), 1},
		[]interface{}{int64(0), 1},
		[]interface{}{int64(1), 1},
		[]interface{}{uint(0), 3},
		[]interface{}{uint(1), 3},
		[]interface{}{uint8(0), 3},
		[]interface{}{uint8(1), 3},
		[]interface{}{uint32(0), 3},
		[]interface{}{uint32(1), 3},
		[]interface{}{uint64(0), 3},
		[]interface{}{uint64(1), 3},
		[]interface{}{float32(0), 1},
		[]interface{}{float32(1), 1},
		[]interface{}{float32(0.1), 3},
		[]interface{}{float64(0), 1},
		[]interface{}{float64(1), 1},
		[]interface{}{float64(0.1), 3},
	}

	for _, v := range tests {
		Equal(t, Length(v[0]), v[1])
	}
}

func TestCompare(t *testing.T) {
	var s *string
	var i interface{} = s
	tests := []struct {
		x   interface{}
		y   interface{}
		op  string
		err error
	}{
		{nil, nil, "", ErrInvalid},
		{nil, nil, CMP.LT, ErrInvalid},
		{s, s, CMP.LT, ErrInvalid},
		{&i, &i, CMP.LT, ErrInvalid},

		{"a", "b", CMP.LT, nil},
		{"b", "a", CMP.LT, ErrGreater},
		{"a", "a", CMP.LT, ErrGreater},
		{"a", 1, CMP.LT, ErrInvalid},
		{"a", "b", CMP.LE, nil},
		{"b", "a", CMP.LE, ErrGreater},
		{"a", "a", CMP.LE, nil},
		{"a", 1, CMP.LE, ErrInvalid},
		{"b", "a", CMP.GT, nil},
		{"a", "b", CMP.GT, ErrLess},
		{"a", "a", CMP.GT, ErrLess},
		{"a", 1, CMP.GT, ErrInvalid},
		{"b", "a", CMP.GE, nil},
		{"a", "b", CMP.GE, ErrLess},
		{"a", "a", CMP.GE, nil},
		{"a", 1, CMP.GE, ErrInvalid},

		{int(1), int(2), CMP.LT, nil},
		{int(2), int(1), CMP.LT, ErrGreater},
		{int(1), int(1), CMP.LT, ErrGreater},
		{int(1), "1", CMP.LT, ErrGreater},
		{int(1), "a", CMP.LT, ErrInvalid},
		{int(1), int(2), CMP.LE, nil},
		{int(2), int(1), CMP.LE, ErrGreater},
		{int(1), int(1), CMP.LE, nil},
		{int(1), "1", CMP.LE, nil},
		{int(1), "a", CMP.LE, ErrInvalid},
		{int(2), int(1), CMP.GT, nil},
		{int(1), int(2), CMP.GT, ErrLess},
		{int(1), int(1), CMP.GT, ErrLess},
		{int(1), "1", CMP.GT, ErrLess},
		{int(1), "a", CMP.GT, ErrInvalid},
		{int(2), int(1), CMP.GE, nil},
		{int(1), int(2), CMP.GE, ErrLess},
		{int(1), int(1), CMP.GE, nil},
		{int(1), "1", CMP.GE, nil},
		{int(1), "a", CMP.GE, ErrInvalid},

		{uint(1), uint(2), CMP.LT, nil},
		{uint(2), uint(1), CMP.LT, ErrGreater},
		{uint(1), uint(1), CMP.LT, ErrGreater},
		{uint(1), "1", CMP.LT, ErrGreater},
		{uint(1), "a", CMP.LT, ErrInvalid},
		{uint(1), uint(2), CMP.LE, nil},
		{uint(2), uint(1), CMP.LE, ErrGreater},
		{uint(1), uint(1), CMP.LE, nil},
		{uint(1), "1", CMP.LE, nil},
		{uint(1), "a", CMP.LE, ErrInvalid},
		{uint(2), uint(1), CMP.GT, nil},
		{uint(1), uint(2), CMP.GT, ErrLess},
		{uint(1), uint(1), CMP.GT, ErrLess},
		{uint(1), "1", CMP.GT, ErrLess},
		{uint(1), "a", CMP.GT, ErrInvalid},
		{uint(2), uint(1), CMP.GE, nil},
		{uint(1), uint(2), CMP.GE, ErrLess},
		{uint(1), uint(1), CMP.GE, nil},
		{uint(1), "1", CMP.GE, nil},
		{uint(1), "a", CMP.GE, ErrInvalid},

		{float64(1), float64(2), CMP.LT, nil},
		{float64(2), float64(1), CMP.LT, ErrGreater},
		{float64(1), float64(1), CMP.LT, ErrGreater},
		{float64(1), "1", CMP.LT, ErrGreater},
		{float64(1), "a", CMP.LT, ErrInvalid},
		{float64(1), float64(2), CMP.LE, nil},
		{float64(2), float64(1), CMP.LE, ErrGreater},
		{float64(1), float64(1), CMP.LE, nil},
		{float64(1), "1", CMP.LE, nil},
		{float64(1), "a", CMP.LE, ErrInvalid},
		{float64(2), float64(1), CMP.GT, nil},
		{float64(1), float64(2), CMP.GT, ErrLess},
		{float64(1), float64(1), CMP.GT, ErrLess},
		{float64(1), "1", CMP.GT, ErrLess},
		{float64(1), "a", CMP.GT, ErrInvalid},
		{float64(2), float64(1), CMP.GE, nil},
		{float64(1), float64(2), CMP.GE, ErrLess},
		{float64(1), float64(1), CMP.GE, nil},
		{float64(1), "1", CMP.GE, nil},
		{float64(1), "a", CMP.GE, ErrInvalid},

		{[]int{1}, []int{1, 2}, CMP.LT, nil},
		{[]int{1, 2}, []int{1}, CMP.LT, ErrGreater},
		{[]int{1}, []int{1}, CMP.LT, ErrGreater},
		{[]int{1}, "1", CMP.LT, ErrInvalid},
		{[]int{1}, []int{1, 2}, CMP.LE, nil},
		{[]int{1, 2}, []int{1}, CMP.LE, ErrGreater},
		{[]int{1}, []int{1}, CMP.LE, nil},
		{[]int{1}, "1", CMP.LE, ErrInvalid},
		{[]int{1, 2}, []int{1}, CMP.GT, nil},
		{[]int{1}, []int{1, 2}, CMP.GT, ErrLess},
		{[]int{1}, []int{1}, CMP.GT, ErrLess},
		{[]int{1}, "1", CMP.GT, ErrInvalid},
		{[]int{1, 2}, []int{1}, CMP.GE, nil},
		{[]int{1}, []int{1, 2}, CMP.GE, ErrLess},
		{[]int{1}, []int{1}, CMP.GE, nil},
		{[]int{1}, "1", CMP.GE, ErrInvalid},

		{map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}, CMP.LT, nil},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1}, CMP.LT, ErrGreater},
		{map[string]int{"a": 1}, map[string]int{"a": 1}, CMP.LT, ErrGreater},
		{map[string]int{"a": 1}, "1", CMP.LT, ErrInvalid},
		{map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}, CMP.LE, nil},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1}, CMP.LE, ErrGreater},
		{map[string]int{"a": 1}, map[string]int{"a": 1}, CMP.LE, nil},
		{map[string]int{"a": 1}, "1", CMP.LE, ErrInvalid},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1}, CMP.GT, nil},
		{map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}, CMP.GT, ErrLess},
		{map[string]int{"a": 1}, map[string]int{"a": 1}, CMP.GT, ErrLess},
		{map[string]int{"a": 1}, "1", CMP.GT, ErrInvalid},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1}, CMP.GE, nil},
		{map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2}, CMP.GE, ErrLess},
		{map[string]int{"a": 1}, map[string]int{"a": 1}, CMP.GE, nil},
		{map[string]int{"a": 1}, "1", CMP.GE, ErrInvalid},
	}

	for _, v := range tests {
		vv := Compare(v.x, v.y, v.op)
		Equal(t, vv, v.err)
		if v.op == CMP.LT {
			Equal(t, IsLt(v.x, v.y), v.err == nil)
		}
		if v.op == CMP.LE {
			Equal(t, IsLe(v.x, v.y), v.err == nil)
		}
		if v.op == CMP.GT {
			Equal(t, IsGt(v.x, v.y), v.err == nil)
		}
		if v.op == CMP.GE {
			Equal(t, IsGe(v.x, v.y), v.err == nil)
		}
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		in  interface{}
		out interface{}
		err error
	}{
		{int64(1), int64(1), nil},
		{uint64(1), int64(1), nil},
		{float64(1), int64(1), nil},
		{"1", int64(1), nil},
		{"1a", int64(0), ErrInvalid},
		{"aa", int64(0), ErrInvalid},
		{true, int64(0), ErrInvalid},
		{[]int{1}, int64(0), ErrInvalid},
		{map[string]int{"a": 1}, int64(0), ErrInvalid},
	}

	for _, v := range tests {
		vv, err := ToInt64(v.in)
		Equal(t, err, v.err)
		Equal(t, vv, v.out)
	}
}

func TestToUint64(t *testing.T) {
	tests := []struct {
		in  interface{}
		out interface{}
		err error
	}{
		{int64(1), uint64(1), nil},
		{uint64(1), uint64(1), nil},
		{float64(1), uint64(1), nil},
		{"1", uint64(1), nil},
		{"1a", uint64(0), ErrInvalid},
		{"aa", uint64(0), ErrInvalid},
		{true, uint64(0), ErrInvalid},
		{[]int{1}, uint64(0), ErrInvalid},
		{map[string]int{"a": 1}, uint64(0), ErrInvalid},
	}

	for _, v := range tests {
		vv, err := ToUint64(v.in)
		Equal(t, err, v.err)
		Equal(t, vv, v.out)
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		in  interface{}
		out interface{}
		err error
	}{
		{int64(1), float64(1), nil},
		{uint64(1), float64(1), nil},
		{float64(1), float64(1), nil},
		{"1", float64(1), nil},
		{"1a", float64(0), ErrInvalid},
		{"aa", float64(0), ErrInvalid},
		{true, float64(0), ErrInvalid},
		{[]int{1}, float64(0), ErrInvalid},
		{map[string]int{"a": 1}, float64(0), ErrInvalid},
	}

	for _, v := range tests {
		vv, err := ToFloat64(v.in)
		Equal(t, err, v.err)
		Equal(t, vv, v.out)
	}
}
