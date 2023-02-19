/*
 * Copyright 2012-2023 Li Kexian
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

package xcache

import (
	"testing"

	"github.com/likexian/gokit/assert"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestNew(t *testing.T) {
	c := New(MemoryCache)
	defer c.Close()

	b := c.Has("x")
	assert.False(t, b)

	v := c.Get("x")
	assert.Equal(t, v, nil)

	err := c.Set("x", 1, 0)
	assert.Nil(t, err)

	b = c.Has("x")
	assert.True(t, b)

	v = c.Get("x")
	assert.Equal(t, v, 1)

	err = c.Del("x")
	assert.Nil(t, err)

	b = c.Has("x")
	assert.False(t, b)

	v = c.Get("x")
	assert.Equal(t, v, nil)
}
