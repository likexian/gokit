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
