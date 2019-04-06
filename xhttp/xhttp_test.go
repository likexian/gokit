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

package xhttp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

const (
	BASEURL = "https://httpbin.org/"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
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

	var h Host = "likexian.com"
	_, _ = req.Do("GET", BASEURL, h)
	assert.Equal(t, req.Request.Host, "likexian.com")
}

func TestSetHeader(t *testing.T) {
	req := New()
	author := req.GetHeader("X-Author")
	assert.Equal(t, author, "")
	req.SetHeader("X-Author", "likexian")
	assert.Equal(t, req.GetHeader("X-Author"), "likexian")

	h1 := Header{
		"X-Version": Version(),
	}
	_, _ = req.Do("GET", BASEURL, h1)
	assert.Equal(t, req.GetHeader("X-Author"), "likexian")
	assert.Equal(t, req.GetHeader("X-Version"), Version())

	h2 := http.Header{
		"X-License": []string{License()},
	}
	_, _ = req.Do("GET", BASEURL, h2)
	assert.Equal(t, req.GetHeader("X-Author"), "likexian")
	assert.Equal(t, req.GetHeader("X-Version"), Version())
	assert.Equal(t, req.GetHeader("X-License"), License())
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

func TestSetClient(t *testing.T) {
	req := New()
	rsp, err := req.Do("GET", BASEURL, &http.Client{})
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
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

func TestJson(t *testing.T) {
	req := New()

	rsp, err := req.Do("GET", BASEURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err := rsp.Json()
	assert.NotNil(t, err)

	rsp, err = req.Do("GET", BASEURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.Response.StatusCode, 200)
	s, err = rsp.Json()
	assert.Nil(t, err)
	assert.Equal(t, s.Get("url").MustString(""), BASEURL+"get")
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

	// set cookies by args
	cookie := &http.Cookie{Name: "k", Value: "likexian"}
	req.SetEnableCookie(true)
	rsp, err = req.Do("GET", BASEURL, cookie)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)
}

func TestQueryParam(t *testing.T) {
	req := New()

	query := QueryParam{"k": "v"}
	_, err := req.Do("GET", BASEURL+"get", query)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"get?k=v")

	query = QueryParam{"a": "1", "b": 2, "c": 3}
	_, err = req.Do("GET", req.Request.URL.String(), query)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"get?k=v&a=1&b=2&c=3")

	query = QueryParam{}
	_, err = req.Do("GET", BASEURL+"get", query)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"get")
}

func TestFormParam(t *testing.T) {
	req := New()

	form := FormParam{"k": "v"}
	rsp, err := req.Do("POST", BASEURL+"post", form)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err := rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"k": "v"`)

	form = FormParam{"a": "1", "b": 2, "c": 3}
	rsp, err = req.Do("POST", req.Request.URL.String(), form)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"a": "1"`)
	assert.Contains(t, text, `"b": "2"`)
	assert.Contains(t, text, `"c": "3"`)

	form = FormParam{}
	rsp, err = req.Do("POST", BASEURL+"post", form)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"form": {}`)

	data := map[string]interface{}{"a": "1", "b": 2, "c": 3}
	rsp, err = req.Do("POST", BASEURL+"post", FormParam(data))
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"a": "1"`)
	assert.Contains(t, text, `"b": "2"`)
	assert.Contains(t, text, `"c": "3"`)
}

func TestValuesParam(t *testing.T) {
	req := New()
	values := url.Values{"k": []string{"v"}}

	// url.Values as query string
	_, err := req.Do("GET", BASEURL+"get", values)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"get?k=v")

	// url.Values as form data
	rsp, err := req.Do("POST", BASEURL+"post", values)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err := rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"k": "v"`)
}

func TestPostBody(t *testing.T) {
	req := New()

	// Post string
	rsp, err := req.Do("POST", BASEURL+"post", "k=v")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err := rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"k": "v"`)

	// Post []byte
	rsp, err = req.Do("POST", req.Request.URL.String(), []byte("a=1&b=2&c=3"))
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"a": "1"`)
	assert.Contains(t, text, `"b": "2"`)
	assert.Contains(t, text, `"c": "3"`)

	// Post bytes.Buffer
	var b bytes.Buffer
	b.Write([]byte("k=v"))
	rsp, err = req.Do("POST", BASEURL+"post", b)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"k": "v"`)

	// Post json string
	rsp, err = req.Do("POST", BASEURL+"post", `{"k": "v"}`, Header{"Content-Type": "application/json"})
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"data": "{\"k\": \"v\"}"`)

	// Post map as json
	data := map[string]interface{}{"a": "1", "b": 2, "c": 3}
	rsp, err = req.Do("POST", BASEURL+"post", JsonParam(data))
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), BASEURL+"post")
	j, err := rsp.Json()
	assert.Nil(t, err)
	assert.Equal(t, j.Get("url").MustString(""), BASEURL+"post")
	assert.Equal(t, j.Get("json.a").MustString(""), "1")
	assert.Equal(t, j.Get("json.b").MustInt(0), 2)
	assert.Equal(t, j.Get("json.c").MustInt(0), 3)
}

func TestPostFile(t *testing.T) {
	req := New()

	// Test post one file
	rsp, err := req.Do("POST", BASEURL+"post", FormFile{"file": "../go.mod"})
	assert.Nil(t, err)
	defer rsp.Close()
	text, err := rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"Content-Type": "multipart/form-data`)
	assert.Contains(t, text, `"file": "module github.com/likexian/gokit`)

	// Test post more files
	rsp, err = req.Do("POST", BASEURL+"post", FormFile{"file_0": "../go.mod"}, FormFile{"file_1": "../go.sum"})
	assert.Nil(t, err)
	defer rsp.Close()
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"Content-Type": "multipart/form-data`)
	assert.Contains(t, text, `"file_0": "module github.com/likexian/gokit`)
	assert.Contains(t, text, `"file_1": "github.com/likexian/gokit`)

	// Test post file and form
	rsp, err = req.Do("POST", BASEURL+"post", FormParam{"k": "v"}, FormFile{"file": "../go.mod"})
	assert.Nil(t, err)
	defer rsp.Close()
	text, err = rsp.String()
	assert.Nil(t, err)
	assert.Contains(t, text, `"Content-Type": "multipart/form-data`)
	assert.Contains(t, text, `"file": "module github.com/likexian/gokit`)
	assert.Contains(t, text, `"k": "v"`)
}

func TestWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.AfterFunc(100*time.Millisecond, cancel)
	}()

	req := New()
	_, err := req.Do("GET", BASEURL+"get", ctx)
	assert.NotNil(t, err)
}

func TestSetRetries(t *testing.T) {
	req := New()
	assert.Panic(t, func() { req.SetRetries() })

	// no retry (default)
	rsp, err := req.Do("Get", "http://127.0.0.1:8080/")
	assert.NotNil(t, err)
	assert.Equal(t, rsp.Tracing.Retries, 0)

	// retry 3 times
	req.SetRetries(3)
	rsp, err = req.Do("Get", "http://127.0.0.1:8080/")
	assert.NotNil(t, err)
	assert.Equal(t, rsp.Tracing.Retries, 3)

	// start http server after 3 second, then request shall success
	go func() {
		time.AfterFunc(3*time.Second, func() {
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello!") })
			http.ListenAndServe("127.0.0.1:5555", nil)
		})
	}()

	// retry until success, sleep 1 second per request
	req.SetRetries(-1, time.Duration(1*time.Second))
	rsp, err = req.Do("Get", "http://127.0.0.1:5555/")
	assert.Nil(t, err)
	defer rsp.Close()
	text, err := rsp.String()
	assert.Nil(t, err)
	assert.Equal(t, text, "Hello!")
	assert.Equal(t, rsp.Tracing.Retries, 2)
}

func TestDump(t *testing.T) {
	req := New()

	req.SetDebug(true, false)
	rsp, err := req.Do("POST", BASEURL+"post", "k=v")
	assert.Nil(t, err)
	defer rsp.Close()
	dump := rsp.Dump()
	assert.NotContains(t, dump, "k=v")

	req.SetDebug(true, true)
	rsp, err = req.Do("POST", BASEURL+"post", "k=v")
	assert.Nil(t, err)
	defer rsp.Close()
	dump = rsp.Dump()
	assert.Contains(t, dump, "k=v\r\n")
}
