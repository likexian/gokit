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

type A struct {
	x, y int
}

type B struct {
	x, y int
}

func TestEqual(t *testing.T) {
	Equal(t, nil, nil, "testing Equal failed")
	Equal(t, true, true, "testing Equal failed")
	Equal(t, "string", "string", "testing Equal failed")
	Equal(t, int(1.0), int(1.0), "testing Equal failed")
	Equal(t, int64(1.0), int64(1.0), "testing Equal failed")
	Equal(t, uint64(1.0), uint64(1.0), "testing Equal failed")
	Equal(t, float64(1.0), float64(1.0), "testing Equal failed")
	Equal(t, []string{"a", "b", "c"}, []string{"a", "b", "c"}, "testing Equal failed")
	Equal(t, []int{1, 2, 3}, []int{1, 2, 3}, "testing Equal failed")
	Equal(t, []float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 3.0}, "testing Equal failed")
	Equal(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}, "testing Equal failed")
	Equal(t, map[string]interface{}{"a": 1, "b": "2"}, map[string]interface{}{"a": 1, "b": "2"}, "testing Equal failed")
	Equal(t, map[string]interface{}{"a": []int{1, 2}}, map[string]interface{}{"a": []int{1, 2}}, "testing Equal failed")
	Equal(t, A{1, 2}, A{1, 2}, "testing Equal failed")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, nil, "", "testing NotEqual failed")
	NotEqual(t, true, false, "testing NotEqual failed")
	NotEqual(t, "string", "strings", "testing NotEqual failed")
	NotEqual(t, int(1.0), int(2.0), "testing NotEqual failed")
	NotEqual(t, int64(1.0), int64(2.0), "testing NotEqual failed")
	NotEqual(t, uint64(1.0), uint64(2.0), "testing NotEqual failed")
	NotEqual(t, float64(1.0), float64(2.0), "testing NotEqual failed")
	NotEqual(t, []string{"a", "b", "c"}, []string{"a", "b", "d"}, "testing NotEqual failed")
	NotEqual(t, []int{1, 2, 3}, []int{1, 2, 4}, "testing NotEqual failed")
	NotEqual(t, []float64{1.0, 2.0, 3.0}, []float64{1.0, 2.0, 4.0}, "testing NotEqual failed")
	NotEqual(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 3}, "testing NotEqual failed")
	NotEqual(t, map[string]interface{}{"a": 1, "b": "2"}, map[string]interface{}{"a": 1, "b": "3"}, "testing NotEqual failed")
	NotEqual(t, map[string]interface{}{"a": []int{1, 2}}, map[string]interface{}{"a": []int{1, 3}}, "testing NotEqual failed")
	NotEqual(t, A{1, 1}, A{1, 2}, "testing NotEqual failed")
	NotEqual(t, A{1, 2}, B{1, 2}, "testing NotEqual failed")
}

func TestNil(t *testing.T) {
	Nil(t, nil, "testing expect to be nil")
	NotNil(t, true, "testing expect to be not nil")
}

func TestTrue(t *testing.T) {
	True(t, true, "testing expect to be true")
	False(t, false, "testing expect to be false")
}
