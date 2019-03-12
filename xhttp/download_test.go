/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package main

import (
	"os"
	"testing"
)

func TestHttpDownload(t *testing.T) {
	var n, c int64
	var err error

	downName := "LICENSE"
	downUrl := "https://raw.githubusercontent.com/likexian/stathub-go/master/"
	zipUrl := "https://github.com/likexian/stathub-go/archive/"

	_, _, err = HttpDownload(downUrl, "", false)
	if err == nil {
		t.Error("No filename shall return error")
		return
	}

	n, c, err = HttpDownload(downUrl+downName, "", false)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("Download successful, file size: %d, cost: %dms", n, c)
	}

	if !FileExists(downName) {
		t.Error("File shall exists but not")
		return
	} else {
		os.Remove(downName)
	}

	_, _, err = HttpDownload(downUrl+downName, "/", false)
	if err == nil {
		t.Error("No filename shall return error")
		return
	}

	_, _, err = HttpDownload(downUrl+"404", downName, false)
	if err == nil {
		t.Error("Http got 404 shall return error")
		return
	}

	n, c, err = HttpDownload(downUrl+downName, downName, false)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("Download successful, file size: %d, cost: %dms", n, c)
	}

	_, _, err = HttpDownload(downUrl+downName, downName, false)
	if err == nil {
		t.Error("File name exists shall return error")
		return
	}

	if !FileExists(downName) {
		t.Error("File shall exists but not")
		return
	} else {
		os.Remove(downName)
	}

	n, c, err = HttpDownload(zipUrl+"master.zip", "", false)
	if err != nil {
		t.Error(err)
	} else {
		os.Remove("master.zip")
		t.Logf("Download successful, file size: %d, cost: %dms", n, c)
	}
}
