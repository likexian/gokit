/*
 * Copyright 2012-2021 Li Kexian
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
	"encoding/base64"
	"os"
	"testing"

	"github.com/likexian/gokit/assert"
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
			"9053253e972cf40443a4083f452f24d4",
			"kFMlPpcs9ARDpAg/RS8k1A==",
		},
		{
			[]interface{}{1234, 5678},
			"9053253e972cf40443a4083f452f24d4",
			"kFMlPpcs9ARDpAg/RS8k1A==",
		},
		{
			[]interface{}{123, 456, 78},
			"37b07f6264727bdd5818a3093e91e6bd",
			"N7B/YmRye91YGKMJPpHmvQ==",
		},
	}

	for _, v := range tests {
		h := Md5(v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
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
			"495b10744ee2c4d5423d7e25e7632d31db08fc00",
			"SVsQdE7ixNVCPX4l52MtMdsI/AA=",
		},
		{
			[]interface{}{1234, 5678},
			"495b10744ee2c4d5423d7e25e7632d31db08fc00",
			"SVsQdE7ixNVCPX4l52MtMdsI/AA=",
		},
		{
			[]interface{}{123, 456, 78},
			"d00e00aca9d085d6b53a363c7a95e112b3064270",
			"0A4ArKnQhda1OjY8epXhErMGQnA=",
		},
	}

	for _, v := range tests {
		h := Sha1(v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
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
			"15b4d8e3c2d7987b16be2cfba411d5fbd980340d3736ad61913aba0ffaf3608c",
			"FbTY48LXmHsWviz7pBHV+9mANA03Nq1hkTq6D/rzYIw=",
		},
		{
			[]interface{}{1234, 5678},
			"15b4d8e3c2d7987b16be2cfba411d5fbd980340d3736ad61913aba0ffaf3608c",
			"FbTY48LXmHsWviz7pBHV+9mANA03Nq1hkTq6D/rzYIw=",
		},
		{
			[]interface{}{123, 456, 78},
			"ac817a17155f86736d8389b3943149f54da90ff16f6376c00b5aef47ed868ea9",
			"rIF6FxVfhnNtg4mzlDFJ9U2pD/FvY3bAC1rvR+2Gjqk=",
		},
	}

	for _, v := range tests {
		h := Sha256(v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
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
			"fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c70669" +
				"5a03cddce372f139eff4d13959da6f1f5d3eabe",
			"+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==",
		},
		{
			[]interface{}{"12345678"},
			"fa585d89c851dd338a70dcf535aa2a92fee7836dd6aff1226583e88e0996293f16bc009c652826e0fc5c70669" +
				"5a03cddce372f139eff4d13959da6f1f5d3eabe",
			"+lhdichR3TOKcNz1Naoqkv7ng23Wr/EiZYPojgmWKT8WvACcZSgm4PxccGaVoDzdzjcvE57/TROVnabx9dPqvg==",
		},
		{
			[]interface{}{"1234", "5678"},
			"5a3d898311876b67bffacdd911908301a6ceb0f1256ddc1bbdbf130beed31f3db726c9cc7eac9dc7f70911df3" +
				"19c30a4509db302a271e72bf8d64c0d73718c7e",
			"Wj2JgxGHa2e/+s3ZEZCDAabOsPElbdwbvb8TC+7THz23JsnMfqydx/cJEd8xnDCkUJ2zAqJx5yv41kwNc3GMfg==",
		},
		{
			[]interface{}{1234, 5678},
			"5a3d898311876b67bffacdd911908301a6ceb0f1256ddc1bbdbf130beed31f3db726c9cc7eac9dc7f70911df3" +
				"19c30a4509db302a271e72bf8d64c0d73718c7e",
			"Wj2JgxGHa2e/+s3ZEZCDAabOsPElbdwbvb8TC+7THz23JsnMfqydx/cJEd8xnDCkUJ2zAqJx5yv41kwNc3GMfg==",
		},
		{
			[]interface{}{123, 456, 78},
			"778475d192b6effb468014c31863e98fdb7b18be556f79940c4a0430866246beaf52e4075f7484a1b15cf5c10" +
				"16d98e244b07293b0cd314cecf08a8d5ae0383e",
			"d4R10ZK27/tGgBTDGGPpj9t7GL5Vb3mUDEoEMIZiRr6vUuQHX3SEobFc9cEBbZjiRLByk7DNMUzs8IqNWuA4Pg==",
		},
	}

	for _, v := range tests {
		h := Sha512(v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
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
			"b13909276f8ff40ee0acfc845ebe2b70",
			"sTkJJ2+P9A7grPyEXr4rcA==",
		},
		{
			[]interface{}{1234, 5678},
			"b13909276f8ff40ee0acfc845ebe2b70",
			"sTkJJ2+P9A7grPyEXr4rcA==",
		},
		{
			[]interface{}{123, 456, 78},
			"b3dc572ce8fd87c48d51c303ae315307",
			"s9xXLOj9h8SNUcMDrjFTBw==",
		},
	}

	for _, v := range tests {
		h := HmacMd5(key, v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
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
			"49498ad44ae78c3da4ce326648c7c2bb6e7838cc",
			"SUmK1ErnjD2kzjJmSMfCu254OMw=",
		},
		{
			[]interface{}{1234, 5678},
			"49498ad44ae78c3da4ce326648c7c2bb6e7838cc",
			"SUmK1ErnjD2kzjJmSMfCu254OMw=",
		},
		{
			[]interface{}{123, 456, 78},
			"74c84407f881334f889a0a0ec8ba43a1ae9d9509",
			"dMhEB/iBM0+ImgoOyLpDoa6dlQk=",
		},
	}

	for _, v := range tests {
		h := HmacSha1(key, v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
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
			"fb8c94cc6c4e7a2c3651160db8f96fd26cb05a230a204610a5a0498bc5e2a6f9",
			"+4yUzGxOeiw2URYNuPlv0mywWiMKIEYQpaBJi8Xipvk=",
		},
		{
			[]interface{}{1234, 5678},
			"fb8c94cc6c4e7a2c3651160db8f96fd26cb05a230a204610a5a0498bc5e2a6f9",
			"+4yUzGxOeiw2URYNuPlv0mywWiMKIEYQpaBJi8Xipvk=",
		},
		{
			[]interface{}{123, 456, 78},
			"821ab013ab6e38ac94d1b9ed996f019e2180f66d7c61be22660711ec29a34c6b",
			"ghqwE6tuOKyU0bntmW8BniGA9m18Yb4iZgcR7CmjTGs=",
		},
	}

	for _, v := range tests {
		h := HmacSha256(key, v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
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
			"defdfafdbdbd488d40691246cffca688c75255ce9bbc7260f63b6e00f5fc4453aff465e6430cb7c7303fb523d" +
				"bf80b99e1f8ea890fe8ab1de19a33d3da497dce",
			"3v36/b29SI1AaRJGz/ymiMdSVc6bvHJg9jtuAPX8RFOv9GXmQwy3xzA/tSPb+AuZ4fjqiQ/oqx3hmjPT2kl9zg==",
		},
		{
			[]interface{}{"12345678"},
			"defdfafdbdbd488d40691246cffca688c75255ce9bbc7260f63b6e00f5fc4453aff465e6430cb7c7303fb523d" +
				"bf80b99e1f8ea890fe8ab1de19a33d3da497dce",
			"3v36/b29SI1AaRJGz/ymiMdSVc6bvHJg9jtuAPX8RFOv9GXmQwy3xzA/tSPb+AuZ4fjqiQ/oqx3hmjPT2kl9zg==",
		},
		{
			[]interface{}{"1234", "5678"},
			"d5f6dcb8107a7f1059185380e5205b206fb22b5e3996c2999e01ae2bda3a1a1d017d4eb6e0be799b6a8d97e68" +
				"d1516e15ab3a9da323af561edd5b635d98a5df6",
			"1fbcuBB6fxBZGFOA5SBbIG+yK145lsKZngGuK9o6Gh0BfU624L55m2qNl+aNFRbhWrOp2jI69WHt1bY12Ypd9g==",
		},
		{
			[]interface{}{1234, 5678},
			"d5f6dcb8107a7f1059185380e5205b206fb22b5e3996c2999e01ae2bda3a1a1d017d4eb6e0be799b6a8d97e68" +
				"d1516e15ab3a9da323af561edd5b635d98a5df6",
			"1fbcuBB6fxBZGFOA5SBbIG+yK145lsKZngGuK9o6Gh0BfU624L55m2qNl+aNFRbhWrOp2jI69WHt1bY12Ypd9g==",
		},
		{
			[]interface{}{123, 456, 78},
			"4c859bcf0ff932b34a4dc0d8b5f1aa537d67ad2be26f67c18de4626a503d28cb69525538efd0f3c363b2b6f9d" +
				"9a7ac121bfd64389ca2d94abf68867acf1b2c45",
			"TIWbzw/5MrNKTcDYtfGqU31nrSvib2fBjeRialA9KMtpUlU479Dzw2OytvnZp6wSG/1kOJyi2Uq/aIZ6zxssRQ==",
		},
	}

	for _, v := range tests {
		h := HmacSha512(key, v.in...)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
	}
}

func TestFileMd5(t *testing.T) {
	fd, err := os.Open("/dev/null")
	assert.Nil(t, err)
	defer fd.Close()

	tests := []struct {
		in  interface{}
		hex string
		b64 string
	}{
		{"/dev/null", "d41d8cd98f00b204e9800998ecf8427e", "1B2M2Y8AsgTpgAmY7PhCfg=="},
		{fd, "d41d8cd98f00b204e9800998ecf8427e", "1B2M2Y8AsgTpgAmY7PhCfg=="},
	}

	for _, v := range tests {
		h, err := FileMd5(v.in)
		assert.Nil(t, err)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
	}

	_, err = FileMd5("/i-am-not-exists")
	assert.NotNil(t, err)

	assert.Panic(t, func() { _, _ = FileMd5(true) })
}

func TestFileSha1(t *testing.T) {
	fd, err := os.Open("/dev/null")
	assert.Nil(t, err)
	defer fd.Close()

	tests := []struct {
		in  interface{}
		hex string
		b64 string
	}{
		{"/dev/null", "da39a3ee5e6b4b0d3255bfef95601890afd80709", "2jmj7l5rSw0yVb/vlWAYkK/YBwk="},
		{fd, "da39a3ee5e6b4b0d3255bfef95601890afd80709", "2jmj7l5rSw0yVb/vlWAYkK/YBwk="},
	}

	for _, v := range tests {
		h, err := FileSha1(v.in)
		assert.Nil(t, err)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
	}

	_, err = FileSha1("/i-am-not-exists")
	assert.NotNil(t, err)

	assert.Panic(t, func() { _, _ = FileSha1(true) })
}

func TestFileSha256(t *testing.T) {
	fd, err := os.Open("/dev/null")
	assert.Nil(t, err)
	defer fd.Close()

	tests := []struct {
		in  interface{}
		hex string
		b64 string
	}{
		{
			"/dev/null",
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		},
		{
			fd,
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		},
	}

	for _, v := range tests {
		h, err := FileSha256(v.in)
		assert.Nil(t, err)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
	}

	_, err = FileSha256("/i-am-not-exists")
	assert.NotNil(t, err)

	assert.Panic(t, func() { _, _ = FileSha256(true) })
}

func TestFileSha512(t *testing.T) {
	fd, err := os.Open("/dev/null")
	assert.Nil(t, err)
	defer fd.Close()

	tests := []struct {
		in  interface{}
		hex string
		b64 string
	}{
		{
			"/dev/null",
			"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d28" +
				"77eec2f63b931bd47417a81a538327af927da3e",
			"z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==",
		},
		{
			fd,
			"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d28" +
				"77eec2f63b931bd47417a81a538327af927da3e",
			"z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==",
		},
	}

	for _, v := range tests {
		h, err := FileSha512(v.in)
		assert.Nil(t, err)
		bs, _ := base64.StdEncoding.DecodeString(v.b64)

		assert.Equal(t, h.Bytes(), bs)
		assert.Equal(t, h.Hex(), v.hex)
		assert.Equal(t, h.Base64(), v.b64)
	}

	_, err = FileSha512("/i-am-not-exists")
	assert.NotNil(t, err)

	assert.Panic(t, func() { _, _ = FileSha512(true) })
}
