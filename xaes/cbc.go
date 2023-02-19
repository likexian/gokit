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
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/likexian/gokit/xhash"
	"github.com/likexian/gokit/xrand"
)

var (
	// ErrMissingEncryptKey is missing encrypt key error
	ErrMissingEncryptKey = errors.New("xaes: key for encrypting is missing")
	// ErrMissingDecryptKey is missing decrypt key error
	ErrMissingDecryptKey = errors.New("xaes: key for decrypting is missing")
	// ErrInvalidIVSize is invalid IV size error
	ErrInvalidIVSize = fmt.Errorf("xaes: length of iv must be %d", aes.BlockSize)
	// ErrInvalidCiphertextSize is invalid ciphertext size error
	ErrInvalidCiphertextSize = fmt.Errorf("xaes: length of ciphertext must be greater than %d",
		aes.BlockSize+sha256.Size)
	// ErrInvalidHmac is invalid hmac error
	ErrInvalidHmac = fmt.Errorf("xaes: hmac is invalid")
	// ErrInvalidPKCS7Padding is invalid pkcs7 padding
	ErrInvalidPKCS7Padding = fmt.Errorf("xaes: invalid PKCS7 padding")
)

// CBCEncryptWithHmac do AES CBC encrypt then HMAC
func CBCEncryptWithHmac(plaintext, key, iv []byte) ([]byte, error) {
	if plaintext == nil {
		return nil, nil
	}

	if key == nil {
		return nil, ErrMissingEncryptKey
	}

	if iv == nil {
		var err error
		iv, err = xrand.Bytes(aes.BlockSize)
		if err != nil {
			return nil, err
		}
	}

	hashKey := xhash.Sha256(key).Bytes()

	ciphertext, err := CBCEncrypt(plaintext, hashKey[:sha256.Size/2], iv)
	if err != nil {
		return nil, err
	}

	hmactext := xhash.HmacSha256(string(hashKey[sha256.Size/2:]), iv).Bytes()
	hmactext = xhash.HmacSha256(string(hmactext), ciphertext).Bytes()

	ciphertext = append(ciphertext, iv...)
	ciphertext = append(ciphertext, hmactext...)

	return ciphertext, nil
}

// CBCDecryptWithHmac do AES CBC decrypt with HMAC
func CBCDecryptWithHmac(ciphertext, key []byte) ([]byte, error) {
	if ciphertext == nil {
		return nil, nil
	}

	if len(ciphertext) <= aes.BlockSize+sha256.Size {
		return nil, ErrInvalidCiphertextSize
	}

	if key == nil {
		return nil, ErrMissingDecryptKey
	}

	hashKey := xhash.Sha256(key).Bytes()

	hmactextIn := ciphertext[len(ciphertext)-sha256.Size:]
	iv := ciphertext[len(ciphertext)-sha256.Size-aes.BlockSize : len(ciphertext)-sha256.Size]
	ciphertext = ciphertext[:len(ciphertext)-sha256.Size-aes.BlockSize]

	hmactext := xhash.HmacSha256(string(hashKey[sha256.Size/2:]), iv).Bytes()
	hmactext = xhash.HmacSha256(string(hmactext), ciphertext).Bytes()

	if !bytes.Equal(hmactext, hmactextIn) {
		return nil, ErrInvalidHmac
	}

	return CBCDecrypt(ciphertext, hashKey[:sha256.Size/2], iv)
}

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
	plaintext, err = PKCS7Padding(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

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

	plaintext, err = PKCS7Unpadding(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// PKCS7Padding do PKCS7 padding
func PKCS7Padding(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	paddingSize := blockSize - len(data)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)

	return append(data, paddingText...), nil
}

// PKCS7Unpadding do PKCS7 unpadding
func PKCS7Unpadding(data []byte) ([]byte, error) {
	dataSize := len(data)
	paddingSize := int(data[dataSize-1])

	if dataSize < paddingSize {
		return nil, ErrInvalidPKCS7Padding
	}

	return data[:(dataSize - paddingSize)], nil
}
