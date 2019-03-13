/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xhttp

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

func TestNew(t *testing.T) {
	http := New("GET", "https://httpbin.org/")
	assert.Equal(t, http.Method, "GET")
	assert.Equal(t, http.URL, "https://httpbin.org/")
	_, err := http.GetHeader("")
	assert.Nil(t, err)
	requestid := http.Requestid
	processid := http.Processid
	assert.NotEqual(t, processid, "")
	assert.NotEqual(t, requestid, "")
	t.Logf("%+v", http.Statics)

	http = New("POST", "https://httpbin.org/get")
	assert.Equal(t, http.Method, "POST")
	assert.Equal(t, http.URL, "https://httpbin.org/get")
	_, err = http.GetHeader("")
	assert.Nil(t, err)
	assert.NotEqual(t, http.Processid, processid)
	assert.NotEqual(t, http.Requestid, requestid)
	t.Logf("%+v", http.Statics)

	http = New("CODE", "https://httpbin.org/get")
	_, err = http.GetHeader("")
	assert.NotNil(t, err)

	http = New("GET", "ftp://httpbin.org/get")
	_, err = http.GetHeader("")
	assert.NotNil(t, err)
}

func TestNext(t *testing.T) {
	http := New("GET", "https://httpbin.org/")
	assert.Equal(t, http.Method, "GET")
	assert.Equal(t, http.URL, "https://httpbin.org/")
	_, err := http.GetHeader("")
	assert.Nil(t, err)
	requestid := http.Requestid
	processid := http.Processid
	assert.NotEqual(t, processid, "")
	assert.NotEqual(t, requestid, "")
	t.Logf("%+v", http.Statics)

	http.Next("POST", "https://httpbin.org/get")
	assert.Equal(t, http.Method, "POST")
	assert.Equal(t, http.URL, "https://httpbin.org/get")
	_, err = http.GetHeader("")
	assert.Nil(t, err)
	assert.Equal(t, http.Processid, processid)
	assert.NotEqual(t, http.Requestid, requestid)
	t.Logf("%+v", http.Statics)
}

func TestSetHeader(t *testing.T) {
	http := New("GET", "https://httpbin.org/get")
	http.SetHeader("X-Test", "Testing")
	http.SetHeader("User-Agent", "Testing UserAgent")
	assert.Equal(t, http.Request.Header["X-Test"][0], "Testing")
	assert.Equal(t, http.Request.Header["User-Agent"][0], "Testing UserAgent")
	t.Logf("%+v", http.Statics)
}

func TestGetHeader(t *testing.T) {
	http := New("GET", "https://httpbin.org/get")
	_, err := http.GetHeader("")
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	t.Logf("%+v", http.Statics)

	for _, k := range []string{"Connection", "Content-Type", "Date", "Server"} {
		h, err := http.GetHeader(k)
		assert.Nil(t, err)
		assert.NotEqual(t, h, "")
	}
}

func TestFollowRedirect(t *testing.T) {
	http := New("GET", "https://httpbin.org/redirect/3")
	loc, err := http.GetHeader("Location")
	assert.Nil(t, err)
	assert.Equal(t, loc, "")
	assert.Equal(t, http.Response.StatusCode, 200)
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/redirect/3")
	http.FollowRedirect = false
	loc, err = http.GetHeader("Location")
	assert.Nil(t, err)
	assert.Equal(t, loc, "/relative-redirect/2")
	assert.Equal(t, http.Response.StatusCode, 302)
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/redirect/3")
	http.FollowRedirect = true
	loc, err = http.GetHeader("Location")
	assert.Nil(t, err)
	assert.Equal(t, loc, "")
	assert.Equal(t, http.Response.StatusCode, 200)
	t.Logf("%+v", http.Statics)
}

func TestBytes(t *testing.T) {
	http := New("GET", "https://httpbin.org/")
	b, err := http.Bytes()
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "<")
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/get")
	b, err = http.Bytes()
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "{")
	t.Logf("%+v", http.Statics)
}

func TestText(t *testing.T) {
	http := New("GET", "https://httpbin.org/")
	s, err := http.String()
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/get")
	s, err = http.String()
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "{")
	t.Logf("%+v", http.Statics)
}

func TestGzip(t *testing.T) {
	http := New("GET", "https://httpbin.org/")
	http.Gzip = false
	s, err := http.String()
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/get")
	http.Gzip = true
	s, err = http.String()
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "{")
	t.Logf("%+v", http.Statics)
}

func TestFile(t *testing.T) {
	defer func() {
		os.Remove("favicon.ico")
		os.Remove("index.html")
		os.Remove("get.html")
		os.RemoveAll("tmp")
	}()

	http := New("GET", "https://httpbin.org/static/favicon.ico")
	ss, err := http.File("")
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	fs, fe := xfile.Size("favicon.ico")
	assert.Nil(t, fe)
	assert.Equal(t, fs, ss)
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/")
	ss, err = http.File("")
	assert.Nil(t, err)
	assert.Equal(t, http.Response.StatusCode, 200)
	fs, fe = xfile.Size("index.html")
	assert.Nil(t, fe)
	assert.Equal(t, fs, ss)
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/get")
	_, err = http.File("./tmp/")
	assert.Nil(t, err)
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/get")
	_, err = http.File("get.html")
	assert.Nil(t, err)
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/get")
	_, err = http.File("get.html")
	assert.NotNil(t, err)
	t.Logf("%+v", http.Statics)

	http.Next("GET", "https://httpbin.org/404")
	_, err = http.File("404.html")
	assert.NotNil(t, err)
	t.Logf("%+v", http.Statics)
}
