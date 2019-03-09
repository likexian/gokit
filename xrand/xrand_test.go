/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xrand

import (
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.NotEqual(t, Version(), "")
	assert.NotEqual(t, Author(), "")
	assert.NotEqual(t, License(), "")
}

func TestInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := Int(0)
		assert.Equal(t, 0, v)

		v = Int(1)
		assert.Equal(t, 0, v)

		v = Int(2)
		assert.True(t, v == 0 || v == 1)
	}
}

func TestIntRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := IntRange(0, 0)
		assert.Equal(t, 0, v)

		v = IntRange(0, 1)
		assert.Equal(t, 0, v)

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
		assert.Equal(t, 10, len(v))
	}
}

func TestStringRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := StringRange(10, "")
		assert.Equal(t, "", v)

		v = StringRange(10, "abc")
		assert.Equal(t, len(v), 10)
		for _, vv := range []rune(v) {
			assert.True(t, vv == 'a' || vv == 'b' || vv == 'c')
		}
	}
}

func TestBytes(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v, err := Bytes(10)
		assert.Nil(t, err)
		assert.Equal(t, 10, len(v))
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
