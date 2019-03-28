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
	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
	"os"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestEncodeDecode(t *testing.T) {
	// Encode struct to string
	v := struct {
		Code    int
		Message string
		Student []string
	}{
		1,
		"Success",
		[]string{"Li Kexian"},
	}
	text, err := Encode(v)
	assert.Nil(t, err)
	assert.Equal(t, text, `{"Code":1,"Message":"Success","Student":["Li Kexian"]}`)

	// Encode map to string
	m := map[string]interface{}{
		"Code":    1,
		"Message": "Success",
		"Student": []string{"Li Kexian"},
	}
	text, err = Encode(m)
	assert.Nil(t, err)
	assert.Equal(t, text, `{"Code":1,"Message":"Success","Student":["Li Kexian"]}`)

	// Decode json to object
	json, err := Decode(text)
	assert.Nil(t, err)
	assert.Equal(t, json.Get("Code").MustInt(), 1)
	assert.Equal(t, json.Get("Message").MustString(), "Success")
	assert.Equal(t, json.Get("Student.0").MustString(), "Li Kexian")
}

func TestDumpLoad(t *testing.T) {
	f := "test.json"
	defer os.Remove(f)

	// Dump struct to file
	v := struct {
		Code    int
		Message string
		Student []string
	}{
		1,
		"Success",
		[]string{"Li Kexian"},
	}
	err := Dump(f, v)
	assert.Nil(t, err)
	assert.True(t, xfile.Exists(f))

	// Load json from file
	json, err := Load(f)
	assert.Nil(t, err)
	assert.Equal(t, json.Get("Code").MustInt(), 1)
	assert.Equal(t, json.Get("Message").MustString(), "Success")
	assert.Equal(t, json.Get("Student.0").MustString(), "Li Kexian")
}
