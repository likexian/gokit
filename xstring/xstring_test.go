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

package xstring

import (
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		in  uint8
		out bool
	}{
		{'a', true},
		{'z', true},
		{'A', true},
		{'Z', true},
		{'0', false},
		{'9', false},
		{'+', false},
		{'@', false},
		{'\t', false},
		{'\n', false},
	}

	for _, v := range tests {
		assert.Equal(t, IsLetter(v.in), v.out)
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"a", false},
		{"-", false},
		{"--1", false},
		{"a1", false},
		{"1a", false},
		{"-1", true},
		{"0", true},
		{"1", true},
		{"-1.1", true},
		{"0.1", true},
		{"1.1", true},
	}

	for _, v := range tests {
		assert.Equal(t, IsNumeric(v.in), v.out)
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"a", "a"},
		{"abc", "cba"},
		{"a123b", "b321a"},
		{"中文可以吗?", "?吗以可文中"},
	}

	for _, v := range tests {
		assert.Equal(t, Reverse(v.in), v.out)
	}
}

func TestExpand(t *testing.T) {
	h := map[string]interface{}{"hello": "world"}
	m := map[string]interface{}{"name": "Li Kexian", "money": 100}

	tests := []struct {
		in  string
		mv  map[string]interface{}
		out string
	}{
		{"", m, ""},
		{"hello", m, "hello"},
		{"i am {}", m, "i am {}"},
		{"i am name}", m, "i am name}"},
		{"i am {name", m, "i am {name"},
		{"i am }name{", m, "i am }name{"},
		{"i am {name}", h, "i am %!name(MISSING)"},
		{"i am {name}", m, "i am Li Kexian"},
		{"i am {{name}}", m, "i am {Li Kexian}"},
		{"i am {{{{{{name}", m, "i am {{{{{Li Kexian"},
		{"i am {{{{{{name}}", m, "i am {{{{{Li Kexian}"},
		{"i am {{{{{{name}}}}}}", m, "i am {{{{{Li Kexian}}}}}"},
		{"i have ${money}", m, "i have $100"},
		{"{name} have ${money}, call {name}.", m, "Li Kexian have $100, call Li Kexian."},
	}

	for _, v := range tests {
		assert.Equal(t, Expand(v.in, v.mv), v.out)
	}
}

func TestLastInIndex(t *testing.T) {
	tests := []struct {
		s   string
		f   string
		out int
	}{
		{"a", "b", -1},
		{"a", "a", 0},
		{"ab", "b", 1},
		{"abc", "c", 2},
		{"{a}", "{", 0},
		{"{{a}", "{", 1},
		{"{{{a}", "{", 2},
	}

	for _, v := range tests {
		assert.Equal(t, LastInIndex(v.s, v.f), v.out)
	}
}
