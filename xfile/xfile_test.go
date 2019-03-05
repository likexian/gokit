/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xfile

import (
	"github.com/likexian/gokit/assert"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	err := os.RemoveAll("tmp")
	assert.Nil(t, err)

	err = os.Mkdir("tmp", 0755)
	assert.Nil(t, err)

	ok := FileExists("tmp/dir")
	assert.False(t, ok, "file expect to be not exists")

	err = os.Mkdir("tmp/dir", 0755)
	assert.Nil(t, err)

	ok = FileExists("tmp/dir")
	assert.True(t, ok, "file expect to be exists")

	ok = IsDir("tmp/dir")
	assert.True(t, ok, "file expect to be dir")

	ok = IsFile("tmp/dir")
	assert.False(t, ok, "file expect to be not file")

	ok = IsFile("tmp/file")
	assert.False(t, ok, "file expect to be not file")

	err = WriteText("tmp/file", "likexian")
	assert.Nil(t, err)

	text, err := ReadText("tmp/file")
	assert.Nil(t, err)
	assert.Equal(t, text, "likexian")

	ok = IsFile("tmp/file")
	assert.True(t, ok, "file expect to be file")

	ok = IsDir("tmp/file")
	assert.False(t, ok, "file expect to be not dir")

	n, err := FileSize("tmp/file")
	assert.Nil(t, err)
	assert.Equal(t, n, int64(8))

	m, err := FileMtime("tmp/file")
	assert.Nil(t, err)
	assert.True(t, m > 0)

	ok = IsSymlink("tmp/link")
	assert.False(t, ok, "file expect to be not expect")

	err = os.Symlink("file", "tmp/link")
	assert.Nil(t, err)

	ok = IsSymlink("tmp/link")
	assert.True(t, ok, "file expect to be symbolic link")

	ok = IsFile("tmp/link")
	assert.True(t, ok, "file expect to be file")

	ok = IsDir("tmp/link")
	assert.False(t, ok, "file expect to be not dir")

	err = ChmodAll("tmp", 0777)
	assert.Nil(t, err)

	err = ChownAll("tmp", 0, 0)
	assert.Nil(t, err)

	os.RemoveAll("tmp")

	pwd := GetPwd()
	assert.NotEqual(t, pwd, "", "pwd expect to be not empty")

	pwd = GetProcPwd()
	assert.NotEqual(t, pwd, "", "pwd expect to be not empty")
}
