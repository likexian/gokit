/*
 * Copyright 2012-2026 Li Kexian
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

package xfile

import (
	"fmt"
	"os"
	"testing"

	"github.com/likexian/gokit/assert"
)

const (
	testDir  = "tmp"
	testFile = "tmp/file"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestFile(t *testing.T) {
	defer os.RemoveAll(testDir)

	err := os.Mkdir(testDir, 0755)
	assert.Nil(t, err)

	ok := Exists(testDir + "/dir")
	assert.False(t, ok, "file expect to be not exists")

	err = os.Mkdir(testDir+"/dir", 0755)
	assert.Nil(t, err)

	ok = Exists(testDir + "/dir")
	assert.True(t, ok, "file expect to be exists")

	ok = IsDir(testDir + "/dir")
	assert.True(t, ok, "file expect to be dir")

	ok = IsFile(testDir + "/dir")
	assert.False(t, ok, "file expect to be not file")

	ok = IsFile(testDir + "/file")
	assert.False(t, ok, "file expect to be not file")

	_, err = New(testDir + "/dir/test/")
	assert.NotNil(t, err)

	fd, err := New(testFile)
	assert.Nil(t, err)
	err = fd.Close()
	assert.Nil(t, err)

	_, err = New(testFile + "/test")
	assert.NotNil(t, err)

	ok = IsFile(testFile)
	assert.True(t, ok, "file expect to be file")

	err = Write(testFile, []byte("likexian"))
	assert.Nil(t, err)

	err = Write(testFile+"/test", []byte("likexian"))
	assert.NotNil(t, err)

	text, err := ReadText(testFile)
	assert.Nil(t, err)
	assert.Equal(t, text, "likexian")

	err = WriteText(testFile, "1\n2\n3\n4\n5")
	assert.Nil(t, err)

	lines, err := ReadLines(testFile, 0)
	assert.Nil(t, err)
	assert.Equal(t, len(lines), 5)

	lines, err = ReadLines(testFile, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(lines), 1)

	_, err = ReadText(testDir + "/not-exists")
	assert.NotNil(t, err)

	_, err = ReadLines(testDir+"/not-exists", 0)
	assert.NotNil(t, err)

	err = WriteText(testFile, "likexian")
	assert.Nil(t, err)

	text, err = ReadText(testFile)
	assert.Nil(t, err)
	assert.Equal(t, text, "likexian")

	ok = IsFile(testFile)
	assert.True(t, ok, "file expect to be file")

	ok = IsDir(testFile)
	assert.False(t, ok, "file expect to be not dir")

	n, err := Size(testFile)
	assert.Nil(t, err)
	assert.Equal(t, n, int64(8))

	m, err := MTime(testFile)
	assert.Nil(t, err)
	assert.True(t, m > 0)

	_, err = Size(testDir + "/not-exists")
	assert.NotNil(t, err)

	_, err = MTime(testDir + "/not-exists")
	assert.NotNil(t, err)

	ok = IsSymlink(testDir + "/link")
	assert.False(t, ok, "file expect to be not expect")

	err = os.Symlink("file", testDir+"/link")
	assert.Nil(t, err)

	ok = IsSymlink(testDir + "/link")
	assert.True(t, ok, "file expect to be symbolic link")

	ok = IsFile(testDir + "/link")
	assert.True(t, ok, "file expect to be file")

	ok = IsDir(testDir + "/link")
	assert.False(t, ok, "file expect to be not dir")

	err = Chmod(testDir, 0777)
	assert.Nil(t, err)

	err = ChmodAll(testDir, 0777)
	assert.Nil(t, err)

	err = Chown(testDir, 0, 0)
	assert.Nil(t, err)

	err = ChownAll(testDir, 0, 0)
	assert.Nil(t, err)

	err = ChmodAll(testDir+"/not-exists", 0777)
	assert.NotNil(t, err)

	err = ChownAll(testDir+"/not-exists", 0, 0)
	assert.NotNil(t, err)
}

func TestNewAndAppend(t *testing.T) {
	defer os.RemoveAll(testDir)

	// init test file
	fd, err := New(testFile)
	assert.Nil(t, err)
	_, _ = fd.Write([]byte("1"))
	text, err := ReadText(testFile)
	assert.Nil(t, err)
	assert.Equal(t, text, "1")

	// test new mode
	fd, err = New(testFile)
	assert.Nil(t, err)
	_, _ = fd.Write([]byte("1"))
	text, err = ReadText(testFile)
	assert.Nil(t, err)
	assert.Equal(t, text, "1")

	// test append mode
	fd, err = NewFile(testFile, true)
	assert.Nil(t, err)
	_, _ = fd.Write([]byte("1"))
	text, err = ReadText(testFile)
	assert.Nil(t, err)
	assert.Equal(t, text, "11")

	// test write text
	err = WriteText(testFile, "1")
	assert.Nil(t, err)
	text, err = ReadText(testFile)
	assert.Nil(t, err)
	assert.Equal(t, text, "1")

	// test append text
	err = AppendText(testFile, "1")
	assert.Nil(t, err)
	text, err = ReadText(testFile)
	assert.Nil(t, err)
	assert.Equal(t, text, "11")
}

func TestReadFirstLine(t *testing.T) {
	defer os.RemoveAll(testDir)

	err := os.Mkdir(testDir, 0755)
	assert.Nil(t, err)

	_, err = ReadFirstLine(testFile)
	assert.NotNil(t, err)

	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"\n", ""},
		{"\n\n", ""},
		{"\n\n\n", ""},
		{"abc\ndef\nghi", "abc"},
		{"\nabc\ndef\nghi", "abc"},
		{"\n\nabc\ndef\nghi", "abc"},
		{"\n\n\nabc\ndef\nghi", "abc"},
	}

	for _, v := range tests {
		err = WriteText(testFile, v.in)
		assert.Nil(t, err)
		line, err := ReadFirstLine(testFile)
		assert.Nil(t, err)
		assert.Equal(t, line, v.out)
	}
}

func TestReadLastLine(t *testing.T) {
	defer os.RemoveAll(testDir)

	err := os.Mkdir(testDir, 0755)
	assert.Nil(t, err)

	_, err = ReadLastLine(testFile)
	assert.NotNil(t, err)

	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"\n", ""},
		{"\n\n", ""},
		{"\n\n\n", ""},
		{"abc\ndef\nghi", "ghi"},
		{"abc\ndef\nghi\n", "ghi"},
		{"abc\ndef\nghi\n\n", "ghi"},
		{"abc\ndef\nghi\n\n\n", "ghi"},
	}

	for _, v := range tests {
		err = WriteText(testFile, v.in)
		assert.Nil(t, err)
		line, err := ReadLastLine(testFile)
		assert.Nil(t, err)
		assert.Equal(t, line, v.out)
	}
}

func TestListDir(t *testing.T) {
	defer os.RemoveAll(testDir)

	ls, err := ListDir("", TypeAll, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 3)

	ls, err = ListDir(testDir, TypeAll, -1)
	assert.NotNil(t, err)
	assert.Equal(t, len(ls), 0)

	for i := 0; i < 10; i++ {
		_ = WriteText(fmt.Sprintf("%s/%d.txt", testDir, i), ".")
		for j := 0; j < 10; j++ {
			_ = WriteText(fmt.Sprintf("%s/%d/%d.txt", testDir, i, j), ".")
		}
	}

	ls, err = ListDir(testDir, TypeAll, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 20)

	ls, err = ListDir(testDir, TypeDir, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 10)

	ls, err = ListDir(testDir, TypeFile, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 10)

	ls, err = ListDir(testDir, TypeAll, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)

	ls, err = ListDir(testDir, TypeDir, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)

	ls, err = ListDir(testDir, TypeFile, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)
}

func TestListDirAll(t *testing.T) {
	defer os.RemoveAll(testDir)

	ls, err := ListDirAll("", TypeAll, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 3)

	ls, err = ListDirAll(testDir, TypeAll, -1)
	assert.NotNil(t, err)
	assert.Equal(t, len(ls), 0)

	for i := 0; i < 10; i++ {
		_ = WriteText(fmt.Sprintf("%s/%d.txt", testDir, i), ".")
		for j := 0; j < 10; j++ {
			_ = WriteText(fmt.Sprintf("%s/%d/%d.txt", testDir, i, j), ".")
		}
	}

	ls, err = ListDirAll(testDir, TypeAll, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 120)

	ls, err = ListDirAll(testDir, TypeDir, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 10)

	ls, err = ListDirAll(testDir, TypeFile, -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 110)

	ls, err = ListDirAll(testDir, TypeAll, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)

	ls, err = ListDirAll(testDir, TypeDir, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)

	ls, err = ListDirAll(testDir, TypeFile, 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)
}

func TestCopy(t *testing.T) {
	defer os.RemoveAll(testDir)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			_ = WriteText(fmt.Sprintf("%s/%d/%d.txt", testDir, i, j), fmt.Sprintf("%d", i+j))
		}
	}

	_ = os.Symlink(testDir+"/0", testDir+"/100")

	err := Copy("", "")
	assert.Equal(t, err, ErrHasExists)

	err = Copy(testDir+"/0", testDir+"/1")
	assert.Equal(t, err, ErrHasExists)

	err = Copy(testDir+"/10", testDir+"/11")
	assert.NotNil(t, err)

	err = Copy(testDir+"/100", testDir+"/101")
	assert.Nil(t, err)
	assert.True(t, Lexists(testDir+"/101"))

	err = Copy(testDir+"/0/0.txt", testDir+"/0/10.txt")
	assert.Nil(t, err)
	assert.True(t, Exists(testDir+"/0/10.txt"))

	err = Copy(testDir+"/0", testDir+"/102")
	assert.Nil(t, err)
	assert.True(t, Exists(testDir+"/102"))
	ls, err := ListDir(testDir+"/0", TypeAll, -1)
	assert.Nil(t, err)
	for _, v := range ls {
		assert.True(t, Exists(testDir+"/102/"+v.Name))
	}
}
