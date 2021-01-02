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

func TestVersion(t *testing.T) {
	Contains(t, Version(), ".")
	Contains(t, Author(), "likexian")
	Contains(t, License(), "Apache License")
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

func TestZero(t *testing.T) {
	Zero(t, []interface{}{}, "testing expect to be zero")
	NotZero(t, true, "testing expect to be not zero")
}

func TestContains(t *testing.T) {
	Contains(t, []int{1, 2, 3}, 2, "testing expect to be contains")
	NotContains(t, []string{"a", "b", "c"}, "d", "testing expect to be not contains")
}

func TestMatch(t *testing.T) {
	Match(t, "li*", "likexian", "testing expect to be match")
	NotMatch(t, "li.kexian", "likexian", "testing expect to be not match")
}

func TestLen(t *testing.T) {
	Len(t, []int{0, 1, 2}, 3, "length expect to be 3")
	NotLen(t, []int{0, 1, 2}, 1, "length expect to be not 1")
}

func TestLtGt(t *testing.T) {
	Lt(t, 1, 2, "testing expect to be less")
	Le(t, 1, 2, "testing expect to be less or equal")
	Le(t, 1, 1, "testing expect to be less or equal")
	Gt(t, 2, 1, "testing expect to be greater")
	Ge(t, 2, 1, "testing expect to be greater or equal")
	Ge(t, 1, 1, "testing expect to be greater or equal")
}

func TestPanic(t *testing.T) {
	Panic(t, func() { panic("failed") })
	Panic(t, func() { panic("failed") }, "why not panic")
}

func TestNotPanic(t *testing.T) {
	NotPanic(t, func() {})
	NotPanic(t, func() {}, "why panic")
}
