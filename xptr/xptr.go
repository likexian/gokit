/*
 * Copyright 2012-2020 Li Kexian
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

package xptr

// Version returns package version
func Version() string {
	return "0.2.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Int converts an int to a pointer
func Int(v int) *int {
	return &v
}

// Int8 converts an int8 to a pointer
func Int8(v int8) *int8 {
	return &v
}

// Int16 converts an int16 to a pointer
func Int16(v int16) *int16 {
	return &v
}

// Int32 converts an int32 to a pointer
func Int32(v int32) *int32 {
	return &v
}

// Int64 converts an int64 to a pointer
func Int64(v int64) *int64 {
	return &v
}

// Uint converts an uint to a pointer
func Uint(v uint) *uint {
	return &v
}

// Uint8 converts an uint8 to a pointer
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint16 converts an uint16 to a pointer
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint32 converts an uint32 to a pointer
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint64 converts an uint64 to a pointer
func Uint64(v uint64) *uint64 {
	return &v
}

// Float32 converts an float32 to a pointer
func Float32(v float32) *float32 {
	return &v
}

// Float64 converts an float64 to a pointer
func Float64(v float64) *float64 {
	return &v
}

// Bool converts an bool to a pointer
func Bool(v bool) *bool {
	return &v
}

// Byte converts an byte to a pointer
func Byte(v byte) *byte {
	return &v
}

// Rune converts an rune to a pointer
func Rune(v rune) *rune {
	return &v
}

// String converts an string to a pointer
func String(v string) *string {
	return &v
}
