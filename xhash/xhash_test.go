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
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	h := Md5("12345678")
	assert.Equal(t, h.Hex(), "25d55ad283aa400af464c76d713c07ad")
	assert.Equal(t, h.B64(), "JdVa0oOqQAr0ZMdtcTwHrQ==")
}

func TestSha1(t *testing.T) {
	h := Sha1("12345678")
	assert.Equal(t, h.Hex(), "7c222fb2927d828af22f592134e8932480637c0d")
	assert.Equal(t, h.B64(), "fCIvspJ9goryL1khNOiTJIBjfA0=")
}

func TestSha256(t *testing.T) {
	h := Sha256("12345678")
	assert.Equal(t, h.Hex(), "ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f")
	assert.Equal(t, h.B64(), "73l8gRjwLftklgfdXT+MdiMEjJwGPVMsyVxe16iYpk8=")
}

func TestSha512(t *testing.T) {
	h := Sha512("12345678")
	assert.Equal(t, h.Hex(), "fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe")
	assert.Equal(t, h.B64(), "+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==")
}

func TestFileMd5(t *testing.T) {
	_, err := FileMd5("/i-am-not-exists")
	assert.NotNil(t, err)

	h, err := FileMd5("/dev/null")
	assert.Nil(t, err)

	assert.Equal(t, h.Hex(), "d41d8cd98f00b204e9800998ecf8427e")
	assert.Equal(t, h.B64(), "1B2M2Y8AsgTpgAmY7PhCfg==")
}

func TestFileSha1(t *testing.T) {
	_, err := FileSha1("/i-am-not-exists")
	assert.NotNil(t, err)

	h, err := FileSha1("/dev/null")
	assert.Nil(t, err)

	assert.Equal(t, h.Hex(), "da39a3ee5e6b4b0d3255bfef95601890afd80709")
	assert.Equal(t, h.B64(), "2jmj7l5rSw0yVb/vlWAYkK/YBwk=")
}

func TestFileSha256(t *testing.T) {
	_, err := FileSha256("/i-am-not-exists")
	assert.NotNil(t, err)

	h, err := FileSha256("/dev/null")
	assert.Nil(t, err)

	assert.Equal(t, h.Hex(), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	assert.Equal(t, h.B64(), "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=")
}

func TestFileSha512(t *testing.T) {
	_, err := FileSha512("/i-am-not-exists")
	assert.NotNil(t, err)

	h, err := FileSha512("/dev/null")
	assert.Nil(t, err)

	assert.Equal(t, h.Hex(), "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e")
	assert.Equal(t, h.B64(), "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==")
}
