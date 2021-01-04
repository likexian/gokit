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
	CBC_AES_KEY    = []byte("1234567812345678")
	CBC_PLAINTEXT  = []byte("hello xaes!")
	CBC_CIPHERTEXT = []byte{32, 73, 238, 61, 249, 194, 179, 122, 136, 105, 227, 59, 55, 89, 10, 97}
)

func TestCBCEncrypt(t *testing.T) {
	ciphertext, err := CBCEncrypt(nil, CBC_AES_KEY, nil)
	assert.Nil(t, err)
	assert.Equal(t, ciphertext, []byte(nil))

	_, err = CBCEncrypt(CBC_PLAINTEXT, nil, nil)
	assert.NotNil(t, err)

	_, err = CBCEncrypt(CBC_PLAINTEXT, CBC_AES_KEY[:1], nil)
	assert.NotNil(t, err)

	_, err = CBCEncrypt(CBC_PLAINTEXT, CBC_AES_KEY, CBC_AES_KEY[:1])
	assert.NotNil(t, err)

	ciphertext, err = CBCEncrypt(CBC_PLAINTEXT, CBC_AES_KEY, nil)
	assert.Nil(t, err)
	assert.Equal(t, ciphertext, CBC_CIPHERTEXT)
}

func TestCBCDecrypt(t *testing.T) {
	CBC_PLAINTEXT, err := CBCDecrypt(nil, CBC_AES_KEY, nil)
	assert.Nil(t, err)
	assert.Equal(t, CBC_PLAINTEXT, []byte(nil))

	_, err = CBCDecrypt(CBC_CIPHERTEXT, nil, nil)
	assert.NotNil(t, err)

	_, err = CBCDecrypt(CBC_CIPHERTEXT, CBC_AES_KEY[:1], nil)
	assert.NotNil(t, err)

	_, err = CBCDecrypt(CBC_CIPHERTEXT, CBC_AES_KEY, CBC_AES_KEY[:1])
	assert.NotNil(t, err)

	CBC_PLAINTEXT, err = CBCDecrypt(CBC_CIPHERTEXT, CBC_AES_KEY, nil)
	assert.Nil(t, err)
	assert.Equal(t, CBC_PLAINTEXT, CBC_PLAINTEXT)
}

func TestCBC(t *testing.T) {
	data := ""
	for i := 0; i < 1000; i++ {
		data += strconv.Itoa(i)

		ciphertext, err := CBCEncrypt([]byte(data), CBC_AES_KEY, nil)
		assert.Nil(t, err)

		plaintext, err := CBCDecrypt(ciphertext, CBC_AES_KEY, nil)
		assert.Nil(t, err)

		assert.Equal(t, plaintext, []byte(data))
	}
}
