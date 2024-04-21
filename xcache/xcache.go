/*
 * Copyright 2012-2024 Li Kexian
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
	"github.com/likexian/gokit/xcache/memory"
)

// Cacher list
const (
	MemoryCache = iota
)

// Cachex is cache interface
type Cachex interface {
	Get(key string) interface{}
	MGet(key ...string) []interface{}
	Set(key string, val interface{}, ttl int64) error
	Has(key string) bool
	Del(key string) error
	Incr(key string) error
	Decr(key string) error
	SetGC(gcInterval, gcMaxOnce int)
	Flush() error
	Close() error
}

// Version returns package version
func Version() string {
	return "0.2.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// New returns a new cacher
func New(cacher int) Cachex {
	switch cacher {
	default:
		return memory.New()
	}
}
