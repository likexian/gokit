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

package xstruct

import (
	"errors"
	"fmt"
	"reflect"
)

// Structx storing struct data
type Structx struct {
	data  interface{}
	value reflect.Value
}

// Fieldx storing struct field
type Fieldx struct {
	data  reflect.StructField
	value reflect.Value
}

// ErrNotStruct is data not a valid struct
var ErrNotStruct = errors.New("xstruct: not a valid struct")

// ErrNotField is field not found
var ErrNotField = errors.New("xstruct: not a valid field name")

// ErrNotExported is field not a export field
var ErrNotExported = errors.New("xstruct: not a exported field")

// errNotSettable is field is not settable
var errNotSettable = errors.New("xstruct: not a settable field")

// Version returns package version
func Version() string {
	return "0.3.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// IsStruct returns if v is a struct
func IsStruct(v interface{}) bool {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	return vv.Kind() == reflect.Struct
}

// Name returns name of struct
func Name(v interface{}) string {
	return New(v).Name()
}

// Struct returns nested struct with name, it panic if not field
func Struct(v interface{}, name string) *Structx {
	return New(v).Struct(name)
}

// Map returns struct name value as map
func Map(v interface{}) map[string]interface{} {
	return New(v).Map()
}

// Names returns names of struct
func Names(v interface{}) []string {
	return New(v).Names()
}

// Tags returns tags of struct
func Tags(v interface{}, key string) (map[string]string, error) {
	return New(v).Tags(key)
}

// Values returns values of struct
func Values(v interface{}) []interface{} {
	return New(v).Values()
}

// Fields return fields of struct
func Fields(v interface{}) []*Fieldx {
	return New(v).Fields()
}

// MustField returns a field with name, panic if error
func MustField(v interface{}, name string) *Fieldx {
	return New(v).MustField(name)
}

// Field returns a field with name
func Field(v interface{}, name string) (*Fieldx, bool) {
	return New(v).Field(name)
}

// Set set value to the field name, must be exported field
func Set(v interface{}, name string, value interface{}) error {
	return New(v).Set(name, value)
}

// Zero set zero value to the field name, must be exported field
func Zero(v interface{}, name string) error {
	return New(v).Zero(name)
}

// New return a new xstruct object, it panic if not struct
func New(v interface{}) *Structx {
	if !IsStruct(v) {
		panic(ErrNotStruct)
	}

	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	s := &Structx{
		data:  v,
		value: vv,
	}

	return s
}

// Name returns name of struct
func (s *Structx) Name() string {
	return s.value.Type().Name()
}

// Struct returns nested struct with name, it panic if not field
func (s *Structx) Struct(name string) *Structx {
	f, ok := s.Field(name)
	if !ok {
		panic(ErrNotField)
	}

	return New(f.Value())
}

// Map returns struct name value as map
func (s *Structx) Map() map[string]interface{} {
	result := map[string]interface{}{}

	fs := s.Fields()
	for _, v := range fs {
		if !v.IsExport() {
			continue
		}
		result[v.Name()] = v.Value()
	}

	return result
}

// Names returns names of struct
func (s *Structx) Names() []string {
	var result []string

	fs := s.Fields()
	for _, v := range fs {
		result = append(result, v.Name())
	}

	return result
}

// Tags returns tags of struct
func (s *Structx) Tags(key string) (map[string]string, error) {
	result := map[string]string{}

	fs := s.Fields()
	for _, v := range fs {
		if !v.IsExport() {
			continue
		}
		result[v.Name()] = v.Tag(key)
	}

	return result, nil
}

// Values returns values of struct
func (s *Structx) Values() []interface{} {
	var result []interface{}

	fs := s.Fields()
	for _, v := range fs {
		if !v.IsExport() {
			continue
		}
		result = append(result, v.Value())
	}

	return result
}

// Fields return fields of struct
func (s *Structx) Fields() []*Fieldx {
	tt := s.value.Type()
	fields := []*Fieldx{}

	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		f := &Fieldx{
			data:  field,
			value: s.value.FieldByName(field.Name),
		}
		fields = append(fields, f)
	}

	return fields
}

// HasField returns field is exists
func (s *Structx) HasField(name string) bool {
	_, ok := s.value.Type().FieldByName(name)
	return ok
}

// MustField returns a field with name, panic if error
func (s *Structx) MustField(name string) *Fieldx {
	f, ok := s.Field(name)
	if !ok {
		panic(ErrNotField)
	}

	return f
}

// Field returns a field with name
func (s *Structx) Field(name string) (*Fieldx, bool) {
	f, ok := s.value.Type().FieldByName(name)
	if !ok {
		return nil, false
	}

	ff := &Fieldx{
		data:  f,
		value: s.value.FieldByName(name),
	}

	return ff, true
}

// IsStruct returns if field name is a struct
func (s *Structx) IsStruct(name string) bool {
	f, ok := s.Field(name)
	if !ok {
		return false
	}

	return IsStruct(f.Value())
}

// Set set value to the field name, must be exported field
func (s *Structx) Set(name string, value interface{}) error {
	f, ok := s.Field(name)
	if !ok {
		return ErrNotField
	}

	return f.Set(value)
}

// Zero set zero value to the field name, must be exported field
func (s *Structx) Zero(name string) error {
	f, ok := s.Field(name)
	if !ok {
		return ErrNotField
	}

	return f.Zero()
}

// Name returns name of field
func (f *Fieldx) Name() string {
	return f.data.Name
}

// Kind returns kind of field
func (f *Fieldx) Kind() reflect.Kind {
	return f.value.Kind()
}

// Tag returns tag of field by key
func (f *Fieldx) Tag(key string) string {
	return f.data.Tag.Get(key)
}

// Value returns value of field
func (f *Fieldx) Value() interface{} {
	return f.value.Interface()
}

// IsAnonymous returns if field is anonymous
func (f *Fieldx) IsAnonymous() bool {
	return f.data.Anonymous
}

// IsExport returns if field is exported
func (f *Fieldx) IsExport() bool {
	return f.data.PkgPath == ""
}

// IsZero returns if field have zero value, for example not initialized
// it panic if field is not exported
func (f *Fieldx) IsZero() bool {
	zero := reflect.Zero(f.value.Type()).Interface()
	return reflect.DeepEqual(f.Value(), zero)
}

// Set set value to the field, must be exported field
func (f *Fieldx) Set(v interface{}) error {
	if !f.IsExport() {
		return ErrNotExported
	}

	if !f.value.CanSet() {
		return errNotSettable
	}

	vv := reflect.ValueOf(v)
	if f.Kind() != vv.Kind() {
		return fmt.Errorf("xstruct: value kind not match, want: %s but got %s", f.Kind(), vv.Kind())
	}

	f.value.Set(vv)

	return nil
}

// Zero set field to zero value, must be exported field
func (f *Fieldx) Zero() error {
	zero := reflect.Zero(f.value.Type()).Interface()
	return f.Set(zero)
}
