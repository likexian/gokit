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
	}
	for _, v := range tests {
		ok := Contains(v[0], v[1])
		assert.Equal(t, ok, v[2])
	}
}
