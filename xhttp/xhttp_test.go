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
	"fmt"
	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	BASEURL = "https://httpbin.org/"
)

func TestVersion(t *testing.T) {
	assert.NotEqual(t, Version(), "")
	assert.NotEqual(t, Author(), "")
	assert.NotEqual(t, License(), "")
}

func TestNew(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "GET")
	assert.Equal(t, req.Request.URL.String(), BASEURL)

	req, err = New("CODE", BASEURL)
	assert.NotNil(t, err)

	req, err = New("GET", "")
	assert.NotNil(t, err)

	req, err = New("GET", "::")
	assert.NotNil(t, err)

	req, err = New("get", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "GET")
	assert.Equal(t, req.Request.URL.String(), BASEURL)

	req, err = New("POST", BASEURL+"post")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "POST")
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")

	req, err = New("GET", BASEURL)
	clientId := req.ClientId
	req, err = New("GET", BASEURL)
	assert.NotEqual(t, req.ClientId, clientId)
}

func TestNext(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	err = req.Next("POST", BASEURL+"post")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "POST")
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")

	err = req.Next("Code", BASEURL+"post")
	assert.NotNil(t, err)

	req, err = New("GET", BASEURL)
	clientId := req.ClientId
	err = req.Next("POST", BASEURL+"post")
	assert.Equal(t, req.ClientId, clientId)
}

func TestSetMethod(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	err = req.SetMethod("CODE")
	assert.NotNil(t, err)

	err = req.SetMethod("POST")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "POST")

	clientId := req.ClientId
	err = req.SetMethod("PUT")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "PUT")
	assert.Equal(t, req.ClientId, clientId)
}

func TestSetURL(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	err = req.SetURL("")
	assert.NotNil(t, err)

	err = req.SetURL("::")
	assert.NotNil(t, err)

	err = req.SetURL(BASEURL + "post")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")

	clientId := req.ClientId
	err = req.SetURL(BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL)
	assert.Equal(t, req.ClientId, clientId)
}

func TestSetClientKey(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.ClientKey, "")
	req.SetClientKey(BASEURL)
	assert.Equal(t, req.ClientKey, BASEURL)
}

func TestSetHost(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	host := req.Request.Host
	req.SetHost("likexian.com")
	assert.Equal(t, req.Request.Host, "likexian.com")
	assert.NotEqual(t, req.Request.Host, host)

	err = req.Next("GET", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Host, host)
}

func TestSetHeader(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	author := req.Request.Header.Get("X-Author")
	assert.Equal(t, author, "")
	req.SetHeader("X-Author", "likexian")
	assert.Equal(t, req.Request.Header.Get("X-Author"), "likexian")

	err = req.Next("GET", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Header.Get("X-Author"), "likexian")
}

func TestSetUA(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	ua := req.Request.Header.Get("User-Agent")
	assert.Equal(t, ua, fmt.Sprintf("GoKit XHTTP Client/%s", Version()))
	req.SetUA("Http Client by likexian")
	assert.Equal(t, req.Request.Header.Get("User-Agent"), "Http Client by likexian")

	err = req.Next("GET", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Header.Get("User-Agent"), "Http Client by likexian")
}

func TestSetReferer(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	referer := req.Request.Header.Get("referer")
	assert.Equal(t, referer, "")
	req.SetHeader("referer", BASEURL)
	assert.Equal(t, req.Request.Header.Get("referer"), BASEURL)

	err = req.Next("GET", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Header.Get("referer"), BASEURL)
}

func TestGetHeader(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	rsp, err := req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	for _, k := range []string{"Connection", "Content-Type", "Date", "Server"} {
		h, err := rsp.GetHeader(k)
		assert.Nil(t, err)
		assert.NotEqual(t, h, "")
	}
}

func TestBytes(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	rsp, err := req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	b, err := rsp.Bytes()
	assert.Nil(t, err)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "<")

	trace := rsp.Trace
	assert.NotEqual(t, trace.Timestamp, "")
	assert.NotEqual(t, trace.Nonce, "")

	err = req.Next("GET", BASEURL+"get")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	b, err = rsp.Bytes()
	assert.Nil(t, err)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "{")

	assert.NotEqual(t, rsp.Trace.Timestamp, "")
	assert.NotEqual(t, rsp.Trace.Nonce, trace.Nonce)
	assert.Equal(t, rsp.Trace.ClientId, trace.ClientId)
	assert.NotEqual(t, rsp.Trace.RequestId, trace.RequestId)

	trace = rsp.Trace
	req, err = New("GET", BASEURL+"status/404")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 404)
	assert.NotEqual(t, rsp.Trace.ClientId, trace.ClientId)
	assert.NotEqual(t, rsp.Trace.RequestId, trace.RequestId)
}

func TestString(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	rsp, err := req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	s, err := rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	err = req.Next("GET", BASEURL+"get")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "{")
}

func TestFile(t *testing.T) {
	defer func() {
		os.Remove("favicon.ico")
		os.Remove("index.html")
		os.Remove("get.html")
		os.RemoveAll("tmp")
	}()

	req, err := New("GET", BASEURL+"static/favicon.ico")
	assert.Nil(t, err)
	rsp, err := req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err := rsp.File()
	assert.Nil(t, err)
	fs, err := xfile.Size("favicon.ico")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	err = req.Next("GET", BASEURL)
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File()
	assert.Nil(t, err)
	fs, err = xfile.Size("index.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	err = req.Next("GET", BASEURL+"get")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File("tmp/")
	assert.Nil(t, err)
	fs, err = xfile.Size("tmp/index.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	err = req.Next("GET", BASEURL+"get")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File("get.html")
	assert.Nil(t, err)
	fs, err = xfile.Size("get.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	err = req.Next("GET", BASEURL+"get")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File("get.html")
	assert.NotNil(t, err)

	err = req.Next("GET", BASEURL+"404")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 404)
	ss, err = rsp.File("404.html")
	assert.NotNil(t, err)
}

func TestSetFollowRedirect(t *testing.T) {
	req, err := New("GET", BASEURL+"redirect/3")
	assert.Nil(t, err)
	rsp, err := req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	loc, err := rsp.GetHeader("Location")
	assert.Nil(t, err)
	assert.Equal(t, loc, "")

	req, err = New("GET", BASEURL+"redirect/3")
	assert.Nil(t, err)
	req.SetFollowRedirect(false)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 302)
	loc, err = rsp.GetHeader("Location")
	assert.Nil(t, err)
	assert.Equal(t, loc, "/relative-redirect/2")

	req, err = New("GET", BASEURL+"redirect/3")
	assert.Nil(t, err)
	req.SetFollowRedirect(true)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	loc, err = rsp.GetHeader("Location")
	assert.Nil(t, err)
	assert.Equal(t, loc, "")
}

func TestSetGzip(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	rsp, err := req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err := rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	req, err = New("GET", BASEURL)
	assert.Nil(t, err)
	req.SetGzip(false)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	req, err = New("GET", BASEURL)
	assert.Nil(t, err)
	req.SetGzip(true)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")
}

func TestSetVerifyTls(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	assert.False(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)

	req.SetVerifyTls(false)
	assert.True(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)

	req.SetVerifyTls(true)
	assert.False(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestSetKeepAlive(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	assert.False(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)

	req.SetKeepAlive(0)
	assert.True(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)

	req.SetKeepAlive(30)
	assert.False(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)
}

func TestSetTimeout(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)

	timeout := req.GetTimeout()
	timeout.ClientTimeout = 10
	timeout.ResponseHeaderTimeout = 3

	req.SetTimeout(timeout)
	assert.Equal(t, req.Client.Timeout, time.Duration(10)*time.Second)
	assert.Equal(t, req.Client.Transport.(*http.Transport).ResponseHeaderTimeout, time.Duration(3)*time.Second)
}

func TestSetProxy(t *testing.T) {
	req, err := New("GET", BASEURL)
	assert.Nil(t, err)
	req.SetProxy("127.0.0.1:8080")
	_, err = req.Do()
	assert.NotNil(t, err)
}

func TestSetEnableCookie(t *testing.T) {
	// not enable cookies
	req, err := New("GET", BASEURL+"cookies/set/k/v")
	req.SetFollowRedirect(false)
	assert.Nil(t, err)
	rsp, err := req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	err = req.Next("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	// enable cookies
	err = req.Next("GET", BASEURL+"cookies/set/k/v")
	req.SetEnableCookie(true)
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	err = req.Next("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	// delete cookies
	err = req.Next("GET", BASEURL+"cookies/delete?k=")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	err = req.Next("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	// set cookie again
	err = req.Next("GET", BASEURL+"cookies/set/k/v")
	req.SetEnableCookie(true)
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	err = req.Next("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	// disable cookies
	err = req.Next("GET", BASEURL+"cookies/set/k/v")
	req.SetEnableCookie(false)
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	err = req.Next("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	rsp, err = req.Do()
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)
}
