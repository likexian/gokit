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
	req := New()

	_, err := req.Do("GET", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "GET")
	assert.Equal(t, req.Request.URL.String(), BASEURL)

	_, err = req.Do("CODE", BASEURL)
	assert.NotNil(t, err)
	_, err = req.Do("GET", "")
	assert.NotNil(t, err)
	_, err = req.Do("GET", "::")
	assert.NotNil(t, err)

	_, err = req.Do("get", BASEURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "GET")
	assert.Equal(t, req.Request.URL.String(), BASEURL)

	_, err = req.Do("POST", BASEURL+"post")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "POST")
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")

	clientId := req.ClientId
	req = New()
	_, err = req.Do("GET", BASEURL)
	assert.NotEqual(t, req.ClientId, clientId)
}

func TestMethod(t *testing.T) {
	rsp, err := Get(BASEURL + "get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = Head(BASEURL + "get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = Post(BASEURL + "post")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = Put(BASEURL + "put")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = Patch(BASEURL + "patch")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = Delete(BASEURL + "delete")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = Options(BASEURL + "get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	req := New()

	rsp, err = req.Get(BASEURL + "get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = req.Head(BASEURL + "get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = req.Post(BASEURL + "post")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = req.Put(BASEURL + "put")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = req.Patch(BASEURL + "patch")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = req.Delete(BASEURL + "delete")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	rsp, err = req.Options(BASEURL + "get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
}

func TestSetSignKey(t *testing.T) {
	req := New()
	assert.Equal(t, req.SignKey, "")
	req.SetSignKey(BASEURL)
	assert.Equal(t, req.SignKey, BASEURL)
}

func TestSetHost(t *testing.T) {
	req := New()
	host := req.Request.Host
	req.SetHost("likexian.com")
	assert.Equal(t, req.Request.Host, "likexian.com")
	assert.NotEqual(t, req.Request.Host, host)
}

func TestSetHeader(t *testing.T) {
	req := New()
	author := req.GetHeader("X-Author")
	assert.Equal(t, author, "")
	req.SetHeader("X-Author", "likexian")
	assert.Equal(t, req.GetHeader("X-Author"), "likexian")
}

func TestSetUA(t *testing.T) {
	req := New()
	ua := req.GetHeader("User-Agent")
	assert.Equal(t, ua, fmt.Sprintf("GoKit XHTTP Client/%s", Version()))
	req.SetUA("Http Client by likexian")
	assert.Equal(t, req.GetHeader("User-Agent"), "Http Client by likexian")
}

func TestSetReferer(t *testing.T) {
	req := New()
	referer := req.GetHeader("referer")
	assert.Equal(t, referer, "")
	req.SetHeader("referer", BASEURL)
	assert.Equal(t, req.GetHeader("referer"), BASEURL)
}

func TestGetHeader(t *testing.T) {
	req := New()
	rsp, err := req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	for _, k := range []string{"Connection", "Content-Type", "Date", "Server"} {
		h := rsp.GetHeader(k)
		assert.NotEqual(t, h, "")
	}
}

func TestBytes(t *testing.T) {
	req := New()
	rsp, err := req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	b, err := rsp.Bytes()
	assert.Nil(t, err)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "<")

	tracing := rsp.Tracing
	assert.NotEqual(t, tracing.Timestamp, "")
	assert.NotEqual(t, tracing.Nonce, "")

	rsp, err = req.Do("GET", BASEURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	b, err = rsp.Bytes()
	assert.Nil(t, err)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "{")

	assert.NotEqual(t, rsp.Tracing.Timestamp, "")
	assert.NotEqual(t, rsp.Tracing.Nonce, tracing.Nonce)
	assert.Equal(t, rsp.Tracing.ClientId, tracing.ClientId)
	assert.NotEqual(t, rsp.Tracing.RequestId, tracing.RequestId)

	tracing = rsp.Tracing
	req = New()
	rsp, err = req.Do("GET", BASEURL+"status/404")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 404)
	assert.NotEqual(t, rsp.Tracing.ClientId, tracing.ClientId)
	assert.NotEqual(t, rsp.Tracing.RequestId, tracing.RequestId)
}

func TestString(t *testing.T) {
	req := New()
	rsp, err := req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)

	s, err := rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	rsp, err = req.Do("GET", BASEURL+"get")
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

	req := New()
	rsp, err := req.Do("GET", BASEURL+"static/favicon.ico")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err := rsp.File()
	assert.Nil(t, err)
	fs, err := xfile.Size("favicon.ico")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	rsp, err = req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File()
	assert.Nil(t, err)
	fs, err = xfile.Size("index.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	rsp, err = req.Do("GET", BASEURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File("tmp/")
	assert.Nil(t, err)
	fs, err = xfile.Size("tmp/index.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	rsp, err = req.Do("GET", BASEURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File("get.html")
	assert.Nil(t, err)
	fs, err = xfile.Size("get.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	rsp, err = req.Do("GET", BASEURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	ss, err = rsp.File("get.html")
	assert.NotNil(t, err)

	rsp, err = req.Do("GET", BASEURL+"404")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 404)
	ss, err = rsp.File("404.html")
	assert.NotNil(t, err)
}

func TestSetFollowRedirect(t *testing.T) {
	req := New()
	rsp, err := req.Do("GET", BASEURL+"redirect/3")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	loc := rsp.GetHeader("Location")
	assert.Equal(t, loc, "")

	req.SetFollowRedirect(false)
	rsp, err = req.Do("GET", BASEURL+"redirect/3")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 302)
	loc = rsp.GetHeader("Location")
	assert.Equal(t, loc, "/relative-redirect/2")

	req.SetFollowRedirect(true)
	rsp, err = req.Do("GET", BASEURL+"redirect/3")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	loc = rsp.GetHeader("Location")
	assert.Equal(t, loc, "")
}

func TestSetGzip(t *testing.T) {
	req := New()
	rsp, err := req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err := rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	req.SetGzip(false)
	rsp, err = req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	req.SetGzip(true)
	rsp, err = req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")
}

func TestSetVerifyTls(t *testing.T) {
	req := New()
	assert.False(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)

	req.SetVerifyTls(false)
	assert.True(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)

	req.SetVerifyTls(true)
	assert.False(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestSetKeepAlive(t *testing.T) {
	req := New()
	assert.False(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)

	req.SetKeepAlive(0)
	assert.True(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)

	req.SetKeepAlive(30)
	assert.False(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)
}

func TestSetTimeout(t *testing.T) {
	req := New()

	timeout := req.GetTimeout()
	timeout.ClientTimeout = 10
	timeout.ResponseHeaderTimeout = 3

	req.SetTimeout(timeout)
	assert.Equal(t, req.Client.Timeout, time.Duration(10)*time.Second)
	assert.Equal(t, req.Client.Transport.(*http.Transport).ResponseHeaderTimeout, time.Duration(3)*time.Second)
}

func TestSetProxy(t *testing.T) {
	req := New().SetProxy("127.0.0.1:8080")
	_, err := req.Do("GET", BASEURL)
	assert.NotNil(t, err)
}

func TestSetEnableCookie(t *testing.T) {
	// not enable cookies
	req := New()
	req.SetFollowRedirect(false)
	rsp, err := req.Do("GET", BASEURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	// enable cookies
	req.SetEnableCookie(true)
	rsp, err = req.Do("GET", BASEURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	// delete cookies
	rsp, err = req.Do("GET", BASEURL+"cookies/delete?k=")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	rsp, err = req.Do("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	// set cookie again
	req.SetEnableCookie(true)
	rsp, err = req.Do("GET", BASEURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	// disable cookies
	req.SetEnableCookie(false)
	rsp, err = req.Do("GET", BASEURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do("GET", BASEURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)
}
