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
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"
)

// Version returns package version
func Version() string {
	return "0.1.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Int returns random int in [0, max)
func Int(max int) int {
	if max <= 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max)
}

// IntRange returns random int in [min, max)
func IntRange(min, max int) int {
	if min > max {
		min, max = max, min
	}

	return Int(max-min) + min
}

// String returns n random string from 0-9,a-z,A-Z
func String(n int) string {
	sources := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return StringRange(n, sources)
}

// StringRange returns n random string base on source
func StringRange(n int, source string) string {
	if source == "" {
		return ""
	}

	ss := []rune(source)
	bs := make([]rune, n)
	for i := range bs {
		bs[i] = ss[Int(len(ss))]
	}

	return string(bs)
}

// Bytes returns n random bytes
func Bytes(n int) (bs []byte, err error) {
	bs = make([]byte, n)
	_, err = crand.Read(bs)
	return
}

// Hex returns hex string of n random bytes
func Hex(n int) (ss string, err error) {
	bs, err := Bytes(n)
	if err != nil {
		return
	}

	ss = hex.EncodeToString(bs)

	return
}

// Base64 returns base64 string of n random bytes
func Base64(n int) (ss string, err error) {
	bs, err := Bytes(n)
	if err != nil {
		return
	}

	ss = base64.StdEncoding.EncodeToString(bs)

	return
}
