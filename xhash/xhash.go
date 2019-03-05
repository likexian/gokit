/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

// XHash storing hash object
type XHash struct {
	Hash hash.Hash
}

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
	return "Apache License, Version 2.0"
}

// Md5 returns md5 hash of string
func Md5(s string) (h XHash) {
	h.Hash = md5.New()
	h.Hash.Write([]byte(s))
	return
}

// Sha1 returns sha1 hash of string
func Sha1(s string) (h XHash) {
	h.Hash = sha1.New()
	h.Hash.Write([]byte(s))
	return
}

// Sha256 returns sha256 hash of string
func Sha256(s string) (h XHash) {
	h.Hash = sha256.New()
	h.Hash.Write([]byte(s))
	return
}

// Sha512 returns sha512 hash of string
func Sha512(s string) (h XHash) {
	h.Hash = sha512.New()
	h.Hash.Write([]byte(s))
	return
}

// FileMd5 returns md5 hash of file
func FileMd5(p string) (h XHash, err error) {
	h.Hash = md5.New()
	err = h.readFile(p)
	return
}

// FileSha1 returns sha1 hash of file
func FileSha1(p string) (h XHash, err error) {
	h.Hash = sha1.New()
	err = h.readFile(p)
	return
}

// FileSha256 returns sha256 hash of file
func FileSha256(p string) (h XHash, err error) {
	h.Hash = sha256.New()
	err = h.readFile(p)
	return
}

// FileSha512 returns sha512 hash of file
func FileSha512(p string) (h XHash, err error) {
	h.Hash = sha512.New()
	err = h.readFile(p)
	return
}

// Hex encoding hash sum as hex string
func (h *XHash) Hex() string {
	return hex.EncodeToString(h.Hash.Sum(nil))
}

// B64 encoding hash sum as base64 string
func (h *XHash) B64() string {
	return base64.StdEncoding.EncodeToString(h.Hash.Sum(nil))
}

// readFile write file content to hash
func (h *XHash) readFile(p string) (err error) {
	fd, err := os.Open(p)
	if err != nil {
		return
	}

	defer fd.Close()
	_, err = io.Copy(h.Hash, fd)

	return
}
