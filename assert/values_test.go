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
	}

	for _, v := range tests {
		False(t, IsZero(v))
	}
}

func TestIsContains(t *testing.T) {
	tests := [][]interface{}{
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

func TestVLen(t *testing.T) {
	var i interface{}
	tests := [][]interface{}{
		[]interface{}{i, 0},
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
		Equal(t, VLen(v[0]), v[1])
	}
}
