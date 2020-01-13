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

package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestBase(t *testing.T) {
	c := New()
	defer c.Close()

	// has
	nx := c.Has("x")
	assert.False(t, nx)

	// set
	err := c.Set("x", 1, -1)
	assert.Nil(t, err)

	// check set
	nx = c.Has("x")
	assert.True(t, nx)

	// get
	v := c.Get("x")
	assert.Equal(t, v, 1)

	// del
	err = c.Del("x")
	assert.Nil(t, err)

	// check del
	nx = c.Has("x")
	assert.False(t, nx)
	v = c.Get("x")
	assert.Equal(t, v, nil)

	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("%d", i)
		err = c.Set(k, i, 0)
		assert.Nil(t, err)
		assert.True(t, c.Has(k))
	}

	// get multiple key
	vs := c.MGet("1", "2", "3")
	assert.Len(t, vs, 3)
	assert.Equal(t, vs[0], 1)
	assert.Equal(t, vs[1], 2)
	assert.Equal(t, vs[2], 3)

	// flush cache
	c.Flush()
	v = c.Get("1")
	assert.Equal(t, v, nil)

	// get on expired key
	err = c.Set("xx", 1, 1)
	assert.Nil(t, err)
	time.Sleep(1 * time.Second)
	v = c.Get("xx")
	assert.Equal(t, v, nil)
}

func TestGC(t *testing.T) {
	c := New()
	defer c.Close()

	c.SetGC(1, 1)

	err := c.Set("x", 1, 1)
	assert.Nil(t, err)

	nx := c.Has("x")
	assert.True(t, nx)

	time.Sleep(1 * time.Second)

	nx = c.Has("x")
	assert.False(t, nx)
}

func TestIncr(t *testing.T) {
	c := New()
	defer c.Close()

	tests := []struct {
		in  interface{}
		out interface{}
	}{
		{int(0), int(1)},
		{int32(0), int32(1)},
		{int64(0), int64(1)},
		{uint(0), uint(1)},
		{uint32(0), uint32(1)},
		{uint64(0), uint64(1)},
	}

	for _, v := range tests {
		_ = c.Set("k", v.in, 0)
		_ = c.Incr("k")
		assert.Equal(t, c.Get("k"), v.out)
	}

	err := c.Incr("x")
	assert.NotNil(t, err)

	err = c.Set("x", "o", 0)
	assert.Nil(t, err)

	err = c.Incr("x")
	assert.NotNil(t, err)
}

func TestDecr(t *testing.T) {
	c := New()
	defer c.Close()

	tests := []struct {
		in  interface{}
		out interface{}
	}{
		{int(1), int(0)},
		{int32(1), int32(0)},
		{int64(1), int64(0)},
		{uint(1), uint(0)},
		{uint32(1), uint32(0)},
		{uint64(1), uint64(0)},
	}

	for _, v := range tests {
		_ = c.Set("k", v.in, 0)
		_ = c.Decr("k")
		assert.Equal(t, c.Get("k"), v.out)
	}

	err := c.Decr("x")
	assert.NotNil(t, err)

	err = c.Set("x", "o", 0)
	assert.Nil(t, err)

	err = c.Decr("x")
	assert.NotNil(t, err)

	for _, v := range []interface{}{
		uint(0),
		uint32(0),
		uint64(0),
	} {
		_ = c.Set("k", v, 0)
		err = c.Decr("k")
		assert.NotNil(t, err)
	}
}
