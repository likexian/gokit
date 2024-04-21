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

package memory

import (
	"errors"
	"sync"
	"time"
)

var (
	// ErrKeyNotExists is key not exists error
	ErrKeyNotExists = errors.New("xcache: the key is not exists")
	// ErrDataTypeNotSupported is data type not supported error
	ErrDataTypeNotSupported = errors.New("xcache: data type is not supported")
	// ErrValueLessThanZero is value less than zero error
	ErrValueLessThanZero = errors.New("xcache: object value is less than zero")
)

// Object is storing single object
type Object struct {
	value  interface{}
	expire int64
}

// Objects is storing all object
type Objects struct {
	values     map[string]*Object
	gcInterval int
	gcMaxOnce  int
	gcExit     chan int
	sync.RWMutex
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

// New init a new cache
func New() *Objects {
	o := &Objects{
		values:     map[string]*Object{},
		gcInterval: 60,
		gcMaxOnce:  100,
		gcExit:     make(chan int),
	}

	go o.gc()

	return o
}

// Set set key value to cache
func (o *Objects) Set(key string, val interface{}, ttl int64) error {
	if ttl > 0 {
		ttl = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	o.Lock()
	defer o.Unlock()
	o.values[key] = &Object{val, ttl}

	return nil
}

// Get get value from cache
func (o *Objects) Get(key string) interface{} {
	o.RLock()
	defer o.RUnlock()

	v, ok := o.values[key]
	if !ok {
		return nil
	}

	if v.expired() {
		return nil
	}

	return v.value
}

// MGet get multiple value from cache
func (o *Objects) MGet(key ...string) []interface{} {
	r := []interface{}{}
	for _, k := range key {
		r = append(r, o.Get(k))
	}

	return r
}

// Has returns key is exists
func (o *Objects) Has(key string) bool {
	o.RLock()
	defer o.RUnlock()
	v, ok := o.values[key]
	if !ok {
		return false
	}

	return !v.expired()
}

// Del remove key from cache
func (o *Objects) Del(key string) error {
	o.Lock()
	defer o.Unlock()
	delete(o.values, key)
	return nil
}

// Incr increase cache counter
func (o *Objects) Incr(key string) error {
	o.Lock()
	defer o.Unlock()

	v, ok := o.values[key]
	if !ok {
		return ErrKeyNotExists
	}

	switch vv := v.value.(type) {
	case int:
		v.value = vv + 1
	case int32:
		v.value = vv + 1
	case int64:
		v.value = vv + 1
	case uint:
		v.value = vv + 1
	case uint32:
		v.value = vv + 1
	case uint64:
		v.value = vv + 1
	default:
		return ErrDataTypeNotSupported
	}

	return nil
}

// Decr decrease cache counter
func (o *Objects) Decr(key string) error {
	o.Lock()
	defer o.Unlock()

	v, ok := o.values[key]
	if !ok {
		return ErrKeyNotExists
	}

	switch vv := v.value.(type) {
	case int:
		v.value = vv - 1
	case int32:
		v.value = vv - 1
	case int64:
		v.value = vv - 1
	case uint:
		if vv <= 0 {
			return ErrValueLessThanZero
		}
		v.value = vv - 1
	case uint32:
		if vv <= 0 {
			return ErrValueLessThanZero
		}
		v.value = vv - 1
	case uint64:
		if vv <= 0 {
			return ErrValueLessThanZero
		}
		v.value = vv - 1
	default:
		return ErrDataTypeNotSupported
	}

	return nil
}

// Flush empty the cache
func (o *Objects) Flush() error {
	o.Lock()
	defer o.Unlock()
	o.values = map[string]*Object{}
	return nil
}

// Close stop the cache service
func (o *Objects) Close() error {
	o.gcExit <- 1
	o.Flush()
	return nil
}

// SetGC set gc interval and max once
func (o *Objects) SetGC(gcInterval, gcMaxOnce int) {
	o.Lock()
	o.gcInterval = gcInterval
	o.gcMaxOnce = gcMaxOnce
	o.Unlock()

	o.gcExit <- 1
	go o.gc()
}

// gc do gc check
func (o *Objects) gc() {
	o.RLock()
	gcInterval := o.gcInterval
	o.RUnlock()

	t := time.NewTicker(time.Duration(gcInterval) * time.Second)
	for {
		select {
		case <-o.gcExit:
			t.Stop()
			return
		case <-t.C:
			o.RLock()
			e := []string{}
			for k, v := range o.values {
				if v.expired() {
					e = append(e, k)
					if len(e) >= o.gcMaxOnce {
						break
					}
				}
			}
			o.RUnlock()
			o.Lock()
			for _, k := range e {
				delete(o.values, k)
			}
			o.Unlock()
		}
	}
}

// expired returns object is expired
func (b *Object) expired() bool {
	if b.expire <= 0 {
		return false
	}

	return time.Now().Unix() >= b.expire
}
