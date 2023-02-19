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

package xaes

import (
	"crypto/aes"
	"strconv"
	"testing"

	"github.com/likexian/gokit/assert"
)

var (
	cbcAESKey         = []byte("1234567812345678")
	cbcPlaintext      = []byte("hello xaes!")
	cbcCiphertext     = []byte{32, 73, 238, 61, 249, 194, 179, 122, 136, 105, 227, 59, 55, 89, 10, 97}
	cbcHmacCiphertext = []byte{56, 91, 201, 6, 75, 116, 161, 43, 218, 129, 203,
		149, 23, 197, 144, 175, 49, 50, 51, 52, 53, 54, 55, 56, 49, 50, 51, 52, 53,
		54, 55, 56, 183, 182, 120, 158, 253, 33, 84, 16, 240, 33, 44, 163, 38, 26,
		103, 57, 96, 131, 128, 47, 251, 7, 180, 234, 107, 134, 0, 126, 1, 1, 227, 15,
	}
)

func TestCBCEncrypt(t *testing.T) {
	ciphertext, err := CBCEncrypt(nil, cbcAESKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, ciphertext, []byte(nil))

	_, err = CBCEncrypt(cbcPlaintext, nil, nil)
	assert.NotNil(t, err)

	_, err = CBCEncrypt(cbcPlaintext, cbcAESKey[:1], nil)
	assert.NotNil(t, err)

	_, err = CBCEncrypt(cbcPlaintext, cbcAESKey, cbcAESKey[:1])
	assert.NotNil(t, err)

	ciphertext, err = CBCEncrypt(cbcPlaintext, cbcAESKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, ciphertext, cbcCiphertext)
}

func TestCBCDecrypt(t *testing.T) {
	plaintext, err := CBCDecrypt(nil, cbcAESKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, plaintext, []byte(nil))

	_, err = CBCDecrypt(cbcCiphertext, nil, nil)
	assert.NotNil(t, err)

	_, err = CBCDecrypt(cbcCiphertext, cbcAESKey[:1], nil)
	assert.NotNil(t, err)

	_, err = CBCDecrypt(cbcCiphertext, cbcAESKey, cbcAESKey[:1])
	assert.NotNil(t, err)

	_, err = CBCDecrypt(cbcCiphertext, []byte("1234567812345677"), nil)
	assert.NotNil(t, err)

	plaintext, err = CBCDecrypt(cbcCiphertext, cbcAESKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, plaintext, cbcPlaintext)
}

func TestCBC(t *testing.T) {
	data := ""
	for i := 0; i < 1000; i++ {
		data += strconv.Itoa(i)

		ciphertext, err := CBCEncrypt([]byte(data), cbcAESKey, nil)
		assert.Nil(t, err)

		plaintext, err := CBCDecrypt(ciphertext, cbcAESKey, nil)
		assert.Nil(t, err)

		assert.Equal(t, plaintext, []byte(data))
	}
}

func TestCBCEncryptWithHmacWithHmac(t *testing.T) {
	ciphertext, err := CBCEncryptWithHmac(nil, cbcAESKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, ciphertext, []byte(nil))

	_, err = CBCEncryptWithHmac(cbcPlaintext, nil, nil)
	assert.NotNil(t, err)

	_, err = CBCEncryptWithHmac(cbcPlaintext, cbcAESKey[:1], nil)
	assert.Nil(t, err)

	_, err = CBCEncryptWithHmac(cbcPlaintext, cbcAESKey, cbcAESKey[:1])
	assert.NotNil(t, err)

	ciphertext, err = CBCEncryptWithHmac(cbcPlaintext, cbcAESKey, cbcAESKey)
	assert.Nil(t, err)
	assert.Equal(t, ciphertext, cbcHmacCiphertext)
}

func TestCBCDecryptWithHmac(t *testing.T) {
	plaintext, err := CBCDecryptWithHmac(nil, cbcAESKey)
	assert.Nil(t, err)
	assert.Equal(t, plaintext, []byte(nil))

	_, err = CBCDecryptWithHmac(cbcHmacCiphertext, nil)
	assert.NotNil(t, err)

	_, err = CBCDecryptWithHmac(cbcHmacCiphertext, cbcAESKey[:1])
	assert.NotNil(t, err)

	_, err = CBCDecryptWithHmac(cbcHmacCiphertext[:1], cbcAESKey)
	assert.NotNil(t, err)

	_, err = CBCDecryptWithHmac(append(cbcHmacCiphertext, '1'), cbcAESKey)
	assert.NotNil(t, err)

	plaintext, err = CBCDecryptWithHmac(cbcHmacCiphertext, cbcAESKey)
	assert.Nil(t, err)
	assert.Equal(t, plaintext, cbcPlaintext)
}

func TestCBCWithHmac(t *testing.T) {
	data := ""
	for i := 0; i < 1000; i++ {
		data += strconv.Itoa(i)

		ciphertext, err := CBCEncryptWithHmac([]byte(data), cbcAESKey, nil)
		assert.Nil(t, err)

		plaintext, err := CBCDecryptWithHmac(ciphertext, cbcAESKey)
		assert.Nil(t, err)

		assert.Equal(t, plaintext, []byte(data))
	}
}

func TestPKCS7Padding(t *testing.T) {
	_, err := PKCS7Padding(nil, aes.BlockSize)
	assert.NotNil(t, err)
}
