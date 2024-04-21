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

package xhash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"io"
	"os"

	"github.com/likexian/gokit/xstring"
)

// Hashx storing hash object
type Hashx struct {
	Hash hash.Hash
}

// Version returns package version
func Version() string {
	return "0.9.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Md5 returns md5 hash of string
func Md5(s ...interface{}) (h Hashx) {
	h.Hash = md5.New()
	h.writeString(s...)
	return
}

// Sha1 returns sha1 hash of string
func Sha1(s ...interface{}) (h Hashx) {
	h.Hash = sha1.New()
	h.writeString(s...)
	return
}

// Sha256 returns sha256 hash of string
func Sha256(s ...interface{}) (h Hashx) {
	h.Hash = sha256.New()
	h.writeString(s...)
	return
}

// Sha512 returns sha512 hash of string
func Sha512(s ...interface{}) (h Hashx) {
	h.Hash = sha512.New()
	h.writeString(s...)
	return
}

// HmacMd5 returns hmac md5 hash of string with key
func HmacMd5(key string, s ...interface{}) (h Hashx) {
	h.Hash = hmac.New(md5.New, []byte(key))
	h.writeString(s...)
	return
}

// HmacSha1 returns hmac sha1 hash of string with key
func HmacSha1(key string, s ...interface{}) (h Hashx) {
	h.Hash = hmac.New(sha1.New, []byte(key))
	h.writeString(s...)
	return
}

// HmacSha256 returns hmac sha256 hash of string with key
func HmacSha256(key string, s ...interface{}) (h Hashx) {
	h.Hash = hmac.New(sha256.New, []byte(key))
	h.writeString(s...)
	return
}

// HmacSha512 returns hmac sha512 hash of string with key
func HmacSha512(key string, s ...interface{}) (h Hashx) {
	h.Hash = hmac.New(sha512.New, []byte(key))
	h.writeString(s...)
	return
}

// FileMd5 returns md5 hash of file
func FileMd5(f interface{}) (h Hashx, err error) {
	h.Hash = md5.New()
	err = h.writeFile(f)
	return
}

// FileSha1 returns sha1 hash of file
func FileSha1(f interface{}) (h Hashx, err error) {
	h.Hash = sha1.New()
	err = h.writeFile(f)
	return
}

// FileSha256 returns sha256 hash of file
func FileSha256(f interface{}) (h Hashx, err error) {
	h.Hash = sha256.New()
	err = h.writeFile(f)
	return
}

// FileSha512 returns sha512 hash of file
func FileSha512(f interface{}) (h Hashx, err error) {
	h.Hash = sha512.New()
	err = h.writeFile(f)
	return
}

// Bytes returns hash sum as bytes
func (h Hashx) Bytes() []byte {
	return h.Hash.Sum(nil)
}

// Hex encoding hash sum as hex string
func (h Hashx) Hex() string {
	return hex.EncodeToString(h.Hash.Sum(nil))
}

// Base64 encoding hash sum as base64 string
func (h Hashx) Base64() string {
	return base64.StdEncoding.EncodeToString(h.Hash.Sum(nil))
}

// writeString write string content to hash
func (h Hashx) writeString(s ...interface{}) {
	length := len(s)
	for _, v := range s {
		switch v := v.(type) {
		case []byte:
			_, _ = h.Hash.Write(v)
		default:
			_, _ = h.Hash.Write([]byte(xstring.ToString(v)))
		}
		if length > 1 {
			_, _ = h.Hash.Write([]byte("\n"))
		}
	}
}

// writeFile write file content to hash
func (h Hashx) writeFile(f interface{}) error {
	switch f := f.(type) {
	case string:
		fd, err := os.Open(f)
		if err != nil {
			return err
		}
		defer fd.Close()
		_, err = io.Copy(h.Hash, fd)
		return err
	case *os.File:
		_, err := io.Copy(h.Hash, f)
		return err
	default:
		panic("xhash: not supported file type")
	}
}
