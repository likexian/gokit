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

package xjson

import (
	"github.com/likexian/simplejson-go"
)

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
	return "Licensed under the Apache License 2.0"
}

// Decode decode json string to Json object
func Decode(s string) (*simplejson.Json, error) {
	return simplejson.Loads(s)
}

// Encode encode value to json text
func Encode(v interface{}) (string, error) {
	return simplejson.New(v).Dumps()
}

// Load load json file to Json object
func Load(f string) (*simplejson.Json, error) {
	return simplejson.Load(f)
}

// Dump dump value to json file
func Dump(f string, v interface{}) error {
	return simplejson.New(v).Dump(f)
}
