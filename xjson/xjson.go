/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
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
	return "Apache License, Version 2.0"
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
