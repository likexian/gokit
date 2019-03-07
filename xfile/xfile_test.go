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

func TestVersion(t *testing.T) {
	assert.NotEqual(t, Version(), "")
	assert.NotEqual(t, Author(), "")
	assert.NotEqual(t, License(), "")
}

func TestFile(t *testing.T) {
	defer os.RemoveAll("tmp")

	err := os.Mkdir("tmp", 0755)
	assert.Nil(t, err)

	ok := Exists("tmp/dir")
	assert.False(t, ok, "file expect to be not exists")

	err = os.Mkdir("tmp/dir", 0755)
	assert.Nil(t, err)

	ok = Exists("tmp/dir")
	assert.True(t, ok, "file expect to be exists")

	ok = IsDir("tmp/dir")
	assert.True(t, ok, "file expect to be dir")

	ok = IsFile("tmp/dir")
	assert.False(t, ok, "file expect to be not file")

	ok = IsFile("tmp/file")
	assert.False(t, ok, "file expect to be not file")

	fd, err := New("tmp/dir/test/")
	assert.NotNil(t, err)

	fd, err = New("tmp/file")
	assert.Nil(t, err)
	err = fd.Close()
	assert.Nil(t, err)

	fd, err = New("tmp/file/test")
	assert.NotNil(t, err)

	ok = IsFile("tmp/file")
	assert.True(t, ok, "file expect to be file")

	err = Write("tmp/file", []byte("likexian"))
	assert.Nil(t, err)

	err = Write("tmp/file/test", []byte("likexian"))
	assert.NotNil(t, err)

	text, err := ReadText("tmp/file")
	assert.Nil(t, err)
	assert.Equal(t, text, "likexian")

	err = WriteText("tmp/file", "1\n2\n3\n4\n5")
	assert.Nil(t, err)

	lines, err := ReadLines("tmp/file", 0)
	assert.Nil(t, err)
	assert.Equal(t, len(lines), 5)

	lines, err = ReadLines("tmp/file", 1)
	assert.Nil(t, err)
	assert.Equal(t, len(lines), 1)

	text, err = ReadText("tmp/not-exists")
	assert.NotNil(t, err)

	lines, err = ReadLines("tmp/not-exists", 0)
	assert.NotNil(t, err)

	err = WriteText("tmp/file", "likexian")
	assert.Nil(t, err)

	text, err = ReadText("tmp/file")
	assert.Nil(t, err)
	assert.Equal(t, text, "likexian")

	ok = IsFile("tmp/file")
	assert.True(t, ok, "file expect to be file")

	ok = IsDir("tmp/file")
	assert.False(t, ok, "file expect to be not dir")

	n, err := Size("tmp/file")
	assert.Nil(t, err)
	assert.Equal(t, n, int64(8))

	m, err := MTime("tmp/file")
	assert.Nil(t, err)
	assert.True(t, m > 0)

	n, err = Size("tmp/not-exists")
	assert.NotNil(t, err)

	n, err = MTime("tmp/not-exists")
	assert.NotNil(t, err)

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

	err = ChmodAll("tmp/not-exists", 0777)
	assert.NotNil(t, err)

	err = ChownAll("tmp/not-exists", 0, 0)
	assert.NotNil(t, err)

	pwd := GetPwd()
	assert.NotEqual(t, pwd, "", "pwd expect to be not empty")

	pwd = GetProcPwd()
	assert.NotEqual(t, pwd, "", "pwd expect to be not empty")
}
