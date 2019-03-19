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
	"testing"
)

type A struct {
	x, y int
}

type B struct {
	x, y int
}

func TestVersion(t *testing.T) {
	assert.NotEqual(t, Version(), "")
	assert.NotEqual(t, Author(), "")
	assert.NotEqual(t, License(), "")
}

func TestContains(t *testing.T) {
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

		[]interface{}{"abc", "a", true},
		[]interface{}{"abc", "d", false},
		[]interface{}{"abc", 1, false},
		[]interface{}{"abc", true, false},

		[]interface{}{map[string]int{"a": 1}, "a", true},
		[]interface{}{map[string]int{"a": 1}, "d", false},
		[]interface{}{map[string]int{"a": 1}, 1, false},
		[]interface{}{map[string]int{"a": 1}, true, false},
	}

	for _, v := range tests {
		ok := Contains(v[0], v[1])
		assert.Equal(t, ok, v[2])
	}

	test := map[string]interface{}{"test": 1}
	assert.Panic(t, func() { Contains(test["test"], test["test"]) })
}

func TestUnique(t *testing.T) {
	tests := [][]interface{}{
		[]interface{}{[]int{0, 0, 1, 1, 1, 2, 2, 3}, []int{0, 1, 2, 3}},
		[]interface{}{[]int64{0, 0, 1, 1, 1, 2, 2, 3}, []int64{0, 1, 2, 3}},
		[]interface{}{[]uint64{0, 0, 1, 1, 1, 2, 2, 3}, []uint64{0, 1, 2, 3}},
		[]interface{}{[]float64{0, 0, 1, 1, 1, 2, 2, 3}, []float64{0, 1, 2, 3}},
		[]interface{}{[]string{"a", "a", "b", "b", "b", "c"}, []string{"a", "b", "c"}},
		[]interface{}{[]bool{true, true, true, false}, []bool{true, false}},
		[]interface{}{[]interface{}{0, 1, 1, "1", 2}, []interface{}{0, 1, "1", 2}},
		[]interface{}{[]interface{}{[]int{0, 1}, []int{0, 1}, []int{1, 2}}, []interface{}{[]int{0, 1}, []int{1, 2}}},
		[]interface{}{[]interface{}{A{0, 1}, A{1, 2}, A{0, 1}, B{0, 1}}, []interface{}{A{0, 1}, A{1, 2}, B{0, 1}}},
	}

	for _, v := range tests {
		u := Unique(v[0])
		assert.Equal(t, u, v[1])
	}

	test := map[string]interface{}{"test": 1}
	assert.Panic(t, func() { Unique(test["test"]) })
}
