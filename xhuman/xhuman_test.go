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

package xhuman

import (
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestFormatByteSize(t *testing.T) {
	tests := []struct {
		in  int64
		p   int
		out string
	}{
		{0, 2, "0.00B"},
		{B, 2, "1.00B"},
		{KB, 2, "1.00KB"},
		{MB, 2, "1.00MB"},
		{GB, 2, "1.00GB"},
		{TB, 2, "1.00TB"},
		{PB, 2, "1.00PB"},
		{EB, 2, "1.00EB"},
		{100 * TB, 2, "100.00TB"},
		{1024 * GB, 2, "1.00TB"},
		{GB + MB, 2, "1.00GB"},
		{GB + 10*MB, 2, "1.01GB"},
		{GB + 100*MB, 2, "1.10GB"},
		{GB + 1000*MB, 2, "1.98GB"},
		{GB + 1024*MB, 2, "2.00GB"},
		{TB + 1000*GB, 0, "2TB"},
		{TB + 1000*GB, 1, "2.0TB"},
		{TB + 1000*GB, 2, "1.98TB"},
		{TB + 1000*GB, 3, "1.977TB"},
		{TB + 1000*GB, 4, "1.9766TB"},
		{TB + 1000*GB, 5, "1.97656TB"},
		{TB + 1000*GB, 6, "1.976562TB"},
		{TB + 1000*GB, 7, "1.9765625TB"},
		{TB + 1000*GB, 8, "1.97656250TB"},
		{TB + 1000*GB, 9, "1.976562500TB"},
		{TB + 1000*GB, 10, "1.9765625000TB"},
		{1024 * 1024 * 1024 * 1024 * 1024, 0, "1PB"},
	}

	for _, v := range tests {
		vv := FormatByteSize(v.in, v.p)
		assert.Equal(t, vv, v.out)
	}
}

func TestParseByteSize(t *testing.T) {
	tests := []struct {
		in  string
		out int64
	}{
		{"0", 0},
		{"10", 10},
		{"0B", 0},
		{"1B", B},
		{"1KB", KB},
		{"1MB", MB},
		{"1GB", GB},
		{"1TB", TB},
		{"1PB", PB},
		{"1EB", EB},
		{"1 EB", EB},
		{"1.0EB", EB},
		{"1.00EB", EB},
		{"1 G", GB},
		{"1 GB", GB},
		{"1.0 G", GB},
		{"1.00 GB", GB},
		{"0.1 KB", 102},
		{"0.10 KB", 102},
		{"100GB", 100 * GB},
		{"100 GB", 100 * GB},
	}

	for _, v := range tests {
		vv, err := ParseByteSize(v.in)
		assert.Nil(t, err)
		assert.Equal(t, vv, v.out)
	}

	for _, v := range []string{"K", "KB", "-1", "-1K", "-1KB", "1AB", "1 K B"} {
		_, err := ParseByteSize(v)
		assert.NotNil(t, err)
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		in  float64
		p   int
		out float64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 0},
		{1, 1, 1},
		{0, 2, 0},
		{1, 2, 1},
		{1.0, 0, 1},
		{1.4, 0, 1},
		{1.5, 0, 2},
		{1.6, 0, 2},
		{1.0, 1, 1.0},
		{1.4, 1, 1.4},
		{1.5, 1, 1.5},
		{1.6, 1, 1.6},
		{1.0, 2, 1.00},
		{1.4, 2, 1.40},
		{1.5, 2, 1.50},
		{1.6, 2, 1.60},
		{-1, 0, -1},
		{-1.4, 0, -1},
		{-1.5, 0, -2},
		{-1.6, 0, -2},
		{-1, 1, -1.0},
		{-1.4, 1, -1.4},
		{-1.5, 1, -1.5},
		{-1.6, 1, -1.6},
		{-1, 2, -1.00},
		{-1.4, 2, -1.40},
		{-1.5, 2, -1.50},
		{-1.6, 2, -1.60},
	}

	for _, v := range tests {
		vv := Round(v.in, v.p)
		assert.Equal(t, vv, v.out)
	}
}

func TestComma(t *testing.T) {
	tests := []struct {
		in  float64
		p   int
		out string
	}{
		{0, 0, "0"},
		{10, 0, "10"},
		{100, 0, "100"},
		{1000, 0, "1,000"},
		{10000, 0, "10,000"},
		{100000, 0, "100,000"},
		{1000000, 0, "1,000,000"},
		{10000000, 0, "10,000,000"},
		{100000000, 0, "100,000,000"},
		{1000000000, 0, "1,000,000,000"},
		{10000000000, 0, "10,000,000,000"},
		{100000000000, 0, "100,000,000,000"},
		{1000000000000, 0, "1,000,000,000,000"},
		{1000000000000.1, 0, "1,000,000,000,000"},
		{1000000000000.1, 1, "1,000,000,000,000.1"},
		{1000000000000.1, 2, "1,000,000,000,000.10"},
		{1000000000000.1, 3, "1,000,000,000,000.100"},
		{1000000000000.1, 4, "1,000,000,000,000.1000"},
		{123456789123456.1, 0, "123,456,789,123,456"},
		{123456789123456.1, 1, "123,456,789,123,456.1"},
		{0, 0, "0"},
		{-10, 0, "-10"},
		{-100, 0, "-100"},
		{-1000, 0, "-1,000"},
		{-10000, 0, "-10,000"},
		{-100000, 0, "-100,000"},
		{-1000000, 0, "-1,000,000"},
		{-10000000, 0, "-10,000,000"},
		{-100000000, 0, "-100,000,000"},
		{-1000000000, 0, "-1,000,000,000"},
		{-10000000000, 0, "-10,000,000,000"},
		{-100000000000, 0, "-100,000,000,000"},
		{-1000000000000, 0, "-1,000,000,000,000"},
		{-1000000000000.1, 0, "-1,000,000,000,000"},
		{-1000000000000.1, 1, "-1,000,000,000,000.1"},
		{-1000000000000.1, 2, "-1,000,000,000,000.10"},
		{-1000000000000.1, 3, "-1,000,000,000,000.100"},
		{-1000000000000.1, 4, "-1,000,000,000,000.1000"},
		{-123456789123456.1, 0, "-123,456,789,123,456"},
		{-123456789123456.1, 1, "-123,456,789,123,456.1"},
	}

	for _, v := range tests {
		vv := Comma(v.in, v.p)
		assert.Equal(t, vv, v.out)
	}
}
