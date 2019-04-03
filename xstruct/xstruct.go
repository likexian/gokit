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

// Struct storing struct data
type Struct struct {
	data  interface{}
	value reflect.Value
}

// Field storing struct field
type Field struct {
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

// New return a new xstruct object, it panic if not struct
func New(v interface{}) *Struct {
	if !IsStruct(v) {
		panic(ErrNotStruct)
	}

	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	s := &Struct{
		data:  v,
		value: vv,
	}

	return s
}

// Name returns name of struct
func (s *Struct) Name() string {
	return s.value.Type().Name()
}

// Struct returns nested struct with name, it panic if not field
func (s *Struct) Struct(name string) *Struct {
	tt := s.value.Type()
	_, ok := tt.FieldByName(name)
	if !ok {
		panic(ErrNotField)
	}

	return New(s.value.FieldByName(name).Interface())
}

// Map returns struct name value as map
func (s *Struct) Map() map[string]interface{} {
	result := map[string]interface{}{}

	fs := s.Fields()
	for _, v := range fs {
		if !v.IsExport() {
			continue
		}
		result[v.data.Name] = v.value.Interface()
	}

	return result
}

// Names returns names of struct
func (s *Struct) Names() []string {
	var result []string

	fs := s.Fields()
	for _, v := range fs {
		result = append(result, v.data.Name)
	}

	return result
}

// Tags returns tags of struct
func (s *Struct) Tags(key string) (map[string]string, error) {
	result := map[string]string{}

	fs := s.Fields()
	for _, v := range fs {
		if !v.IsExport() {
			continue
		}
		result[v.data.Name] = v.Tag(key)
	}

	return result, nil
}

// Values returns values of struct
func (s *Struct) Values() []interface{} {
	var result []interface{}

	fs := s.Fields()
	for _, v := range fs {
		if !v.IsExport() {
			continue
		}
		vv := s.value.FieldByName(v.data.Name)
		result = append(result, vv.Interface())
	}

	return result
}

// Fields return fields of struct
func (s *Struct) Fields() []*Field {
	tt := s.value.Type()
	fields := []*Field{}

	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		f := &Field{
			data:  field,
			value: s.value.FieldByName(field.Name),
		}
		fields = append(fields, f)
	}

	return fields
}

// MustField returns a field with name, panic if error
func (s *Struct) MustField(name string) *Field {
	f, ok := s.Field(name)
	if !ok {
		panic(ErrNotField)
	}

	return f
}

// Field returns a field with name
func (s *Struct) Field(name string) (*Field, bool) {
	tt := s.value.Type()
	field, ok := tt.FieldByName(name)
	if !ok {
		return nil, false
	}

	ff := &Field{
		data:  field,
		value: s.value.FieldByName(name),
	}

	return ff, true
}

// Set set value to the field name, must be exported field
func (s *Struct) Set(name string, value interface{}) error {
	f, ok := s.Field(name)
	if !ok {
		return ErrNotField
	}

	return f.Set(value)
}

// Zero set zero value to the field name, must be exported field
func (s *Struct) Zero(name string) error {
	f, ok := s.Field(name)
	if !ok {
		return ErrNotField
	}

	return f.Zero()
}

// Name returns name of field
func (f *Field) Name() string {
	return f.data.Name
}

// Kind returns kind of field
func (f *Field) Kind() reflect.Kind {
	return f.value.Kind()
}

// Tag returns tag of field by key
func (f *Field) Tag(key string) string {
	return f.data.Tag.Get(key)
}

// Value returns value of field
func (f *Field) Value() interface{} {
	return f.value.Interface()
}

// IsAnonymous returns if field is anonymous
func (f *Field) IsAnonymous() bool {
	return f.data.Anonymous
}

// IsExport returns if field is exported
func (f *Field) IsExport() bool {
	return f.data.PkgPath == ""
}

// IsZero returns if field have zero value, for example not initialized
// it panic if field is not exported
func (f *Field) IsZero() bool {
	zero := reflect.Zero(f.value.Type()).Interface()
	return reflect.DeepEqual(f.Value(), zero)
}

// Set set value to the field, must be exported field
func (f *Field) Set(v interface{}) error {
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
func (f *Field) Zero() error {
	zero := reflect.Zero(f.value.Type()).Interface()
	return f.Set(zero)
}

// IsStruct returns if v is a struct
func IsStruct(v interface{}) bool {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	return vv.Kind() == reflect.Struct
}
