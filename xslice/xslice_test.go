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
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
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
