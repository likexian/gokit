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

package xaes

import (
	"strconv"
	"testing"

	"github.com/likexian/gokit/assert"
)

var (
	cbcAESKey     = []byte("1234567812345678")
	cbcPlaintext  = []byte("hello xaes!")
	cbcCiphertext = []byte{32, 73, 238, 61, 249, 194, 179, 122, 136, 105, 227, 59, 55, 89, 10, 97}
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
	cbcPlaintext, err := CBCDecrypt(nil, cbcAESKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, cbcPlaintext, []byte(nil))

	_, err = CBCDecrypt(cbcCiphertext, nil, nil)
	assert.NotNil(t, err)

	_, err = CBCDecrypt(cbcCiphertext, cbcAESKey[:1], nil)
	assert.NotNil(t, err)

	_, err = CBCDecrypt(cbcCiphertext, cbcAESKey, cbcAESKey[:1])
	assert.NotNil(t, err)

	cbcPlaintext, err = CBCDecrypt(cbcCiphertext, cbcAESKey, nil)
	assert.Nil(t, err)
	assert.Equal(t, cbcPlaintext, cbcPlaintext)
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
