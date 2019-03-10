/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xos

import (
	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
	"os"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.NotEqual(t, Version(), "")
	assert.NotEqual(t, Author(), "")
	assert.NotEqual(t, License(), "")
}

func TestLookupUser(t *testing.T) {
	uid, gid, err := LookupUser("nobody")
	assert.Nil(t, err)
	assert.True(t, uid > 0)
	assert.True(t, gid > 0)
}

func TestWritePid(t *testing.T) {
	pid := "xos.pid"
	defer os.Remove(pid)

	err := WritePid(pid)
	assert.Nil(t, err)
	assert.True(t, xfile.Exists(pid))
}

func TestSetid(t *testing.T) {
	err := SetUid(0)
	assert.Nil(t, err)

	err = SetGid(0)
	assert.Nil(t, err)

	err = SetUser("nobody")
	assert.Nil(t, err)

	err = SetUid(0)
	assert.NotNil(t, err)

	err = SetGid(0)
	assert.NotNil(t, err)

	err = SetUser("root")
	assert.NotNil(t, err)
}
