/*
 * Copyright 2012-2019 Li Kexian
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
	"github.com/likexian/gokit/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestMd5(t *testing.T) {
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"25d55ad283aa400af464c76d713c07ad",
			"JdVa0oOqQAr0ZMdtcTwHrQ==",
		},
		{
			[]interface{}{"12345678"},
			"25d55ad283aa400af464c76d713c07ad",
			"JdVa0oOqQAr0ZMdtcTwHrQ==",
		},
		{
			[]interface{}{"1234", "5678"},
			"25d55ad283aa400af464c76d713c07ad",
			"JdVa0oOqQAr0ZMdtcTwHrQ==",
		},
		{
			[]interface{}{1234, 5678},
			"25d55ad283aa400af464c76d713c07ad",
			"JdVa0oOqQAr0ZMdtcTwHrQ==",
		},
		{
			[]interface{}{123, 456, 78},
			"25d55ad283aa400af464c76d713c07ad",
			"JdVa0oOqQAr0ZMdtcTwHrQ==",
		},
	}

	for _, v := range tests {
		h := Md5(v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
}

func TestSha1(t *testing.T) {
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"7c222fb2927d828af22f592134e8932480637c0d",
			"fCIvspJ9goryL1khNOiTJIBjfA0=",
		},
		{
			[]interface{}{"12345678"},
			"7c222fb2927d828af22f592134e8932480637c0d",
			"fCIvspJ9goryL1khNOiTJIBjfA0=",
		},
		{
			[]interface{}{"1234", "5678"},
			"7c222fb2927d828af22f592134e8932480637c0d",
			"fCIvspJ9goryL1khNOiTJIBjfA0=",
		},
		{
			[]interface{}{1234, 5678},
			"7c222fb2927d828af22f592134e8932480637c0d",
			"fCIvspJ9goryL1khNOiTJIBjfA0=",
		},
		{
			[]interface{}{123, 456, 78},
			"7c222fb2927d828af22f592134e8932480637c0d",
			"fCIvspJ9goryL1khNOiTJIBjfA0=",
		},
	}

	for _, v := range tests {
		h := Sha1(v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
}

func TestSha256(t *testing.T) {
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f",
			"73l8gRjwLftklgfdXT+MdiMEjJwGPVMsyVxe16iYpk8=",
		},
		{
			[]interface{}{"12345678"},
			"ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f",
			"73l8gRjwLftklgfdXT+MdiMEjJwGPVMsyVxe16iYpk8=",
		},
		{
			[]interface{}{"1234", "5678"},
			"ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f",
			"73l8gRjwLftklgfdXT+MdiMEjJwGPVMsyVxe16iYpk8=",
		},
		{
			[]interface{}{1234, 5678},
			"ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f",
			"73l8gRjwLftklgfdXT+MdiMEjJwGPVMsyVxe16iYpk8=",
		},
		{
			[]interface{}{123, 456, 78},
			"ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f",
			"73l8gRjwLftklgfdXT+MdiMEjJwGPVMsyVxe16iYpk8=",
		},
	}

	for _, v := range tests {
		h := Sha256(v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
}

func TestSha512(t *testing.T) {
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe",
			"+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==",
		},
		{
			[]interface{}{"12345678"},
			"fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe",
			"+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==",
		},
		{
			[]interface{}{"1234", "5678"},
			"fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe",
			"+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==",
		},
		{
			[]interface{}{1234, 5678},
			"fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe",
			"+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==",
		},
		{
			[]interface{}{123, 456, 78},
			"fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c706695a03cddce372f139eff4d13959da6f1f5d3eabe",
			"+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==",
		},
	}

	for _, v := range tests {
		h := Sha512(v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
}

func TestHmacMd5(t *testing.T) {
	key := "87654321"
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"2589a5e790d014bf42e049126624cbdd",
			"JYml55DQFL9C4EkSZiTL3Q==",
		},
		{
			[]interface{}{"12345678"},
			"2589a5e790d014bf42e049126624cbdd",
			"JYml55DQFL9C4EkSZiTL3Q==",
		},
		{
			[]interface{}{"1234", "5678"},
			"2589a5e790d014bf42e049126624cbdd",
			"JYml55DQFL9C4EkSZiTL3Q==",
		},
		{
			[]interface{}{1234, 5678},
			"2589a5e790d014bf42e049126624cbdd",
			"JYml55DQFL9C4EkSZiTL3Q==",
		},
		{
			[]interface{}{123, 456, 78},
			"2589a5e790d014bf42e049126624cbdd",
			"JYml55DQFL9C4EkSZiTL3Q==",
		},
	}

	for _, v := range tests {
		h := HmacMd5(key, v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
}

func TestHmacSha1(t *testing.T) {
	key := "87654321"
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"3f271885b5503055cf2b93facc5cde88f94f7708",
			"PycYhbVQMFXPK5P6zFzeiPlPdwg=",
		},
		{
			[]interface{}{"12345678"},
			"3f271885b5503055cf2b93facc5cde88f94f7708",
			"PycYhbVQMFXPK5P6zFzeiPlPdwg=",
		},
		{
			[]interface{}{"1234", "5678"},
			"3f271885b5503055cf2b93facc5cde88f94f7708",
			"PycYhbVQMFXPK5P6zFzeiPlPdwg=",
		},
		{
			[]interface{}{1234, 5678},
			"3f271885b5503055cf2b93facc5cde88f94f7708",
			"PycYhbVQMFXPK5P6zFzeiPlPdwg=",
		},
		{
			[]interface{}{123, 456, 78},
			"3f271885b5503055cf2b93facc5cde88f94f7708",
			"PycYhbVQMFXPK5P6zFzeiPlPdwg=",
		},
	}

	for _, v := range tests {
		h := HmacSha1(key, v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
}

func TestHmacSha256(t *testing.T) {
	key := "87654321"
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"18cef3462f052e9fad5f4198f4ef397783189c6e25ab9dafbc7071401065ac76",
			"GM7zRi8FLp+tX0GY9O85d4MYnG4lq52vvHBxQBBlrHY=",
		},
		{
			[]interface{}{"12345678"},
			"18cef3462f052e9fad5f4198f4ef397783189c6e25ab9dafbc7071401065ac76",
			"GM7zRi8FLp+tX0GY9O85d4MYnG4lq52vvHBxQBBlrHY=",
		},
		{
			[]interface{}{"1234", "5678"},
			"18cef3462f052e9fad5f4198f4ef397783189c6e25ab9dafbc7071401065ac76",
			"GM7zRi8FLp+tX0GY9O85d4MYnG4lq52vvHBxQBBlrHY=",
		},
		{
			[]interface{}{1234, 5678},
			"18cef3462f052e9fad5f4198f4ef397783189c6e25ab9dafbc7071401065ac76",
			"GM7zRi8FLp+tX0GY9O85d4MYnG4lq52vvHBxQBBlrHY=",
		},
		{
			[]interface{}{123, 456, 78},
			"18cef3462f052e9fad5f4198f4ef397783189c6e25ab9dafbc7071401065ac76",
			"GM7zRi8FLp+tX0GY9O85d4MYnG4lq52vvHBxQBBlrHY=",
		},
	}

	for _, v := range tests {
		h := HmacSha256(key, v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
}

func TestHmacSha512(t *testing.T) {
	key := "87654321"
	tests := []struct {
		in  []interface{}
		hex string
		b64 string
	}{
		{
			[]interface{}{[]byte("12345678")},
			"defdfafdbdbd488d40691246cffca688c75255ce9bbc7260f63b6e00f5fc4453aff465e6430cb7c7303fb523dbf80b99e1f8ea890fe8ab1de19a33d3da497dce",
			"3v36/b29SI1AaRJGz/ymiMdSVc6bvHJg9jtuAPX8RFOv9GXmQwy3xzA/tSPb+AuZ4fjqiQ/oqx3hmjPT2kl9zg==",
		},
		{
			[]interface{}{"12345678"},
			"defdfafdbdbd488d40691246cffca688c75255ce9bbc7260f63b6e00f5fc4453aff465e6430cb7c7303fb523dbf80b99e1f8ea890fe8ab1de19a33d3da497dce",
			"3v36/b29SI1AaRJGz/ymiMdSVc6bvHJg9jtuAPX8RFOv9GXmQwy3xzA/tSPb+AuZ4fjqiQ/oqx3hmjPT2kl9zg==",
		},
		{
			[]interface{}{"1234", "5678"},
			"defdfafdbdbd488d40691246cffca688c75255ce9bbc7260f63b6e00f5fc4453aff465e6430cb7c7303fb523dbf80b99e1f8ea890fe8ab1de19a33d3da497dce",
			"3v36/b29SI1AaRJGz/ymiMdSVc6bvHJg9jtuAPX8RFOv9GXmQwy3xzA/tSPb+AuZ4fjqiQ/oqx3hmjPT2kl9zg==",
		},
		{
			[]interface{}{1234, 5678},
			"defdfafdbdbd488d40691246cffca688c75255ce9bbc7260f63b6e00f5fc4453aff465e6430cb7c7303fb523dbf80b99e1f8ea890fe8ab1de19a33d3da497dce",
			"3v36/b29SI1AaRJGz/ymiMdSVc6bvHJg9jtuAPX8RFOv9GXmQwy3xzA/tSPb+AuZ4fjqiQ/oqx3hmjPT2kl9zg==",
		},
		{
			[]interface{}{123, 456, 78},
			"defdfafdbdbd488d40691246cffca688c75255ce9bbc7260f63b6e00f5fc4453aff465e6430cb7c7303fb523dbf80b99e1f8ea890fe8ab1de19a33d3da497dce",
			"3v36/b29SI1AaRJGz/ymiMdSVc6bvHJg9jtuAPX8RFOv9GXmQwy3xzA/tSPb+AuZ4fjqiQ/oqx3hmjPT2kl9zg==",
		},
	}

	for _, v := range tests {
		h := HmacSha512(key, v.in...)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.B64(), v.b64)
	}
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
