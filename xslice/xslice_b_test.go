/*
 * Copyright 2012-2020 Li Kexian
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
	"strconv"
	"testing"
)

var (
	minData []string
	midData []string
	maxData []string
)

func init() {
	for i := 0; i < 100; i++ {
		minData = append(minData, strconv.Itoa(i))
	}
	for i := 0; i < 1000; i++ {
		midData = append(midData, strconv.Itoa(i))
	}
	for i := 0; i < 10000; i++ {
		maxData = append(maxData, strconv.Itoa(i))
	}
}

func benchmarkIsUnique(b *testing.B, v interface{}) {
	for i := 0; i < b.N; i++ {
		IsUnique(v)
	}
}

func benchmarkUnique(b *testing.B, v interface{}) {
	for i := 0; i < b.N; i++ {
		Unique(v)
	}
}

func benchmarkReverse(b *testing.B, v interface{}) {
	for i := 0; i < b.N; i++ {
		Reverse(v)
	}
}

func benchmarkShuffle(b *testing.B, v interface{}) {
	for i := 0; i < b.N; i++ {
		Shuffle(v)
	}
}

func BenchmarkIsUniqueMin(b *testing.B) { benchmarkIsUnique(b, minData) }
func BenchmarkIsUniqueMid(b *testing.B) { benchmarkIsUnique(b, midData) }
func BenchmarkIsUniqueMax(b *testing.B) { benchmarkIsUnique(b, maxData) }

func BenchmarkUniqueMin(b *testing.B) { benchmarkUnique(b, minData) }
func BenchmarkUniqueMid(b *testing.B) { benchmarkUnique(b, midData) }
func BenchmarkUniqueMax(b *testing.B) { benchmarkUnique(b, maxData) }

func BenchmarkReverseMin(b *testing.B) { benchmarkReverse(b, minData) }
func BenchmarkReverseMid(b *testing.B) { benchmarkReverse(b, midData) }
func BenchmarkReverseMax(b *testing.B) { benchmarkReverse(b, maxData) }

func BenchmarkShuffleMin(b *testing.B) { benchmarkShuffle(b, minData) }
func BenchmarkShuffleMid(b *testing.B) { benchmarkShuffle(b, midData) }
func BenchmarkShuffleMax(b *testing.B) { benchmarkShuffle(b, maxData) }
