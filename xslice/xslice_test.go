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
		[]interface{}{1, 1},
		[]interface{}{1.0, 1.0},
		[]interface{}{true, true},
	}

	for _, v := range tests {
		assert.Equal(t, Unique(v[0]), v[1])
	}
}
