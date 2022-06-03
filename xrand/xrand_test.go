/*
 * Copyright 2012-2022 Li Kexian
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

package xrand

import (
	"testing"

	"github.com/likexian/gokit/assert"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := Int(0)
		assert.Equal(t, v, 0)

		v = Int(1)
		assert.Equal(t, v, 0)

		v = Int(2)
		assert.True(t, v == 0 || v == 1)
	}
}

func TestIntRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := IntRange(0, 0)
		assert.Equal(t, v, 0)

		v = IntRange(0, 1)
		assert.Equal(t, v, 0)

		v = IntRange(0, 2)
		assert.True(t, v == 0 || v == 1)

		v = IntRange(2, 0)
		assert.True(t, v == 0 || v == 1)

		v = IntRange(2, 2)
		assert.True(t, v == 2)

		v = IntRange(100, 10000)
		assert.True(t, v >= 100 && v < 10000)
	}
}

func TestString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := String(10)
		assert.Equal(t, len(v), 10)
	}
}

func TestStringRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := StringRange(10, "")
		assert.Equal(t, v, "")

		v = StringRange(10, "abc")
		assert.Equal(t, len(v), 10)
		for _, vv := range v {
			assert.True(t, vv == 'a' || vv == 'b' || vv == 'c')
		}
	}
}

func TestBytes(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v, err := Bytes(10)
		assert.Nil(t, err)
		assert.Equal(t, len(v), 10)
	}
}

func TestHex(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v, err := Hex(10)
		assert.Nil(t, err)
		assert.True(t, len(v) > 10)
	}
}

func TestBase64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v, err := Base64(10)
		assert.Nil(t, err)
		assert.True(t, len(v) > 10)
	}
}
