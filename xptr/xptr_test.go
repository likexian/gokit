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

package xptr

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
	v := int(1)
	p := Int(v)
	assert.Equal(t, *p, v)
}

func TestInt8(t *testing.T) {
	v := int8(1)
	p := Int8(v)
	assert.Equal(t, *p, v)
}

func TestInt16(t *testing.T) {
	v := int16(1)
	p := Int16(v)
	assert.Equal(t, *p, v)
}

func TestInt32(t *testing.T) {
	v := int32(1)
	p := Int32(v)
	assert.Equal(t, *p, v)
}

func TestInt64(t *testing.T) {
	v := int64(1)
	p := Int64(v)
	assert.Equal(t, *p, v)
}

func TestUint(t *testing.T) {
	v := uint(1)
	p := Uint(v)
	assert.Equal(t, *p, v)
}

func TestUint8(t *testing.T) {
	v := uint8(1)
	p := Uint8(v)
	assert.Equal(t, *p, v)
}

func TestUint16(t *testing.T) {
	v := uint16(1)
	p := Uint16(v)
	assert.Equal(t, *p, v)
}

func TestUint32(t *testing.T) {
	v := uint32(1)
	p := Uint32(v)
	assert.Equal(t, *p, v)
}

func TestUint64(t *testing.T) {
	v := uint64(1)
	p := Uint64(v)
	assert.Equal(t, *p, v)
}

func TestFloat32(t *testing.T) {
	v := float32(1)
	p := Float32(v)
	assert.Equal(t, *p, v)
}

func TestFloat64(t *testing.T) {
	v := float64(1)
	p := Float64(v)
	assert.Equal(t, *p, v)
}

func TestBool(t *testing.T) {
	v := bool(true)
	p := Bool(v)
	assert.Equal(t, *p, v)
}

func TestByte(t *testing.T) {
	v := byte(1)
	p := Byte(v)
	assert.Equal(t, *p, v)
}

func TestRune(t *testing.T) {
	v := rune(1)
	p := Rune(v)
	assert.Equal(t, *p, v)
}

func TestString(t *testing.T) {
	v := string("1")
	p := String(v)
	assert.Equal(t, *p, v)
}
