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
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

var (
	// ErrMissingEncryptKey is missing encrypt key error
	ErrMissingEncryptKey = errors.New("xaes: key for encrypting is missing")
	// ErrMissingDecryptKey is missing decrypt key error
	ErrMissingDecryptKey = errors.New("xaes: key for decrypting is missing")
	// ErrInvalidIVSize is invalid IV size error
	ErrInvalidIVSize = fmt.Errorf("xaes: length of iv must be %d", aes.BlockSize)
)

// CBCEncrypt do AES CBC encrypt
func CBCEncrypt(plaintext, key, iv []byte) ([]byte, error) {
	if plaintext == nil {
		return nil, nil
	}

	if key == nil {
		return nil, ErrMissingEncryptKey
	}

	if iv == nil {
		iv = make([]byte, aes.BlockSize)
	}

	if len(iv) != aes.BlockSize {
		return nil, ErrInvalidIVSize
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// CBCDecrypt do AES CBC decrypt
func CBCDecrypt(ciphertext, key, iv []byte) ([]byte, error) {
	if ciphertext == nil {
		return nil, nil
	}

	if key == nil {
		return nil, ErrMissingDecryptKey
	}

	if iv == nil {
		iv = make([]byte, aes.BlockSize)
	}

	if len(iv) != aes.BlockSize {
		return nil, ErrInvalidIVSize
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS7Unpadding(plaintext)

	return plaintext, nil
}

// PKCS7Padding do PKCS7 padding
func PKCS7Padding(data []byte, blockSize int) []byte {
	paddingSize := blockSize - len(data)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(data, paddingText...)
}

// PKCS7Unpadding do PKCS7 unpadding
func PKCS7Unpadding(data []byte) []byte {
	dataSize := len(data)
	paddingSize := int(data[dataSize-1])
	return data[:(dataSize - paddingSize)]
}
