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

package xfile

import (
	"fmt"
	"github.com/likexian/gokit/assert"
	"os"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
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

	_, err = ReadText("tmp/not-exists")
	assert.NotNil(t, err)

	_, err = ReadLines("tmp/not-exists", 0)
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

	_, err = Size("tmp/not-exists")
	assert.NotNil(t, err)

	_, err = MTime("tmp/not-exists")
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

func TestListDir(t *testing.T) {
	defer os.RemoveAll("tmp")

	ls, err := ListDir("", "", -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 3)

	ls, err = ListDir("tmp", "", -1)
	assert.NotNil(t, err)
	assert.Equal(t, len(ls), 0)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			WriteText(fmt.Sprintf("tmp/%d/%d.txt", i, j), ".")
		}
	}

	ls, err = ListDir("tmp", "", -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 110)

	ls, err = ListDir("tmp", "dir", -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 10)

	ls, err = ListDir("tmp", "file", -1)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 100)

	ls, err = ListDir("tmp", "", 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)

	ls, err = ListDir("tmp", "dir", 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)

	ls, err = ListDir("tmp", "file", 5)
	assert.Nil(t, err)
	assert.Equal(t, len(ls), 5)
}

func TestCopy(t *testing.T) {
	defer os.RemoveAll("tmp")

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			WriteText(fmt.Sprintf("tmp/%d/%d.txt", i, j), fmt.Sprintf("%d", i+j))
		}
	}

	os.Symlink("tmp/0", "tmp/100")

	err := Copy("", "")
	assert.Equal(t, err, ErrHasExists)

	err = Copy("tmp/0", "tmp/1")
	assert.Equal(t, err, ErrHasExists)

	err = Copy("tmp/10", "tmp/11")
	assert.NotNil(t, err)

	err = Copy("tmp/100", "tmp/101")
	assert.Nil(t, err)
	assert.True(t, Lexists("tmp/101"))

	err = Copy("tmp/0/0.txt", "tmp/0/10.txt")
	assert.Nil(t, err)
	assert.True(t, Exists("tmp/0/10.txt"))

	err = Copy("tmp/0", "tmp/102")
	assert.Nil(t, err)
	assert.True(t, Exists("tmp/102"))
	ls, err := ListDir("tmp/0", "", -1)
	assert.Nil(t, err)
	for _, v := range ls {
		assert.True(t, Exists("tmp/102/"+v.Name))
	}
}
