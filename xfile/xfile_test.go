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
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	err := os.RemoveAll("tmp")
	if err != nil {
		t.Fatal(err)
	}

	err = os.Mkdir("tmp", 0755)
	if err != nil {
		t.Fatal(err)
	}

	ok := FileExists("tmp/dir")
	if ok {
		t.Fatal("file expect to be not exists")
	}

	err = os.Mkdir("tmp/dir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	ok = FileExists("tmp/dir")
	if !ok {
		t.Fatal("file expect to be exists")
	}

	ok = IsDir("tmp/dir")
	if !ok {
		t.Fatal("file expect to be dir")
	}

	ok = IsFile("tmp/dir")
	if ok {
		t.Fatal("file expect to be not file")
	}

	ok = IsFile("tmp/file")
	if ok {
		t.Fatal("file expect to be not file")
	}

	err = WriteText("tmp/file", "likexian")
	if err != nil {
		t.Fatal(err)
	}

	text, err := ReadText("tmp/file")
	if err != nil {
		t.Fatal(err)
	} else {
		if text != "likexian" {
			t.Fatalf("file text expect to be likexian but got %s", text)
		}
	}

	ok = IsFile("tmp/file")
	if !ok {
		t.Fatal("file expect to be file")
	}

	ok = IsDir("tmp/file")
	if ok {
		t.Fatal("file expect to be not dir")
	}

	n, err := FileSize("tmp/file")
	if err != nil {
		t.Fatal(err)
	} else {
		if n != 8 {
			t.Fatalf("file size expect to be 8 but got %d", n)
		}
	}

	m, err := FileMtime("tmp/file")
	if err != nil {
		t.Fatal(err)
	} else {
		if m <= 0 {
			t.Fatalf("get fail mtime failed")
		}
	}

	ok = IsSymlink("tmp/link")
	if ok {
		t.Fatal("file expect to be not exists")
	}

	err = os.Symlink("file", "tmp/link")
	if err != nil {
		t.Fatal(err)
	}

	ok = IsSymlink("tmp/link")
	if !ok {
		t.Fatal("file expect to be symbolic link")
	}

	ok = IsFile("tmp/link")
	if !ok {
		t.Fatal("file expect to be file")
	}

	ok = IsDir("tmp/link")
	if ok {
		t.Fatal("file expect to be not dir")
	}

	err = ChmodAll("tmp", 0777)
	if err != nil {
		t.Fatal(err)
	}

	err = ChownAll("tmp", 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	os.RemoveAll("tmp")

	pwd := GetPwd()
	if pwd == "" {
		t.Fatal("pwd expect to be not empty")
	}

	pwd = GetProcPwd()
	if pwd == "" {
		t.Fatal("pwd expect to be not empty")
	}
}
