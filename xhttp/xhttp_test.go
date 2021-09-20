/*
 * Copyright 2012-2021 Li Kexian
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
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xjson"
	"github.com/likexian/gokit/xtime"
)

var (
	LOCALURL = ServerForTesting("6666")
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestNew(t *testing.T) {
	req := New()
	ctx := context.Background()

	_, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "GET")
	assert.Equal(t, req.Request.URL.String(), LOCALURL)

	_, err = req.Do(ctx, "CODE", LOCALURL)
	assert.NotNil(t, err)
	_, err = req.Do(ctx, "GET", "")
	assert.NotNil(t, err)
	_, err = req.Do(ctx, "GET", "::")
	assert.NotNil(t, err)

	_, err = req.Do(ctx, "get", LOCALURL)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "GET")
	assert.Equal(t, req.Request.URL.String(), LOCALURL)

	_, err = req.Do(ctx, "POST", LOCALURL+"post")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.Method, "POST")
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")

	clientID := req.ClientID
	req = New()
	_, err = req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	assert.NotEqual(t, req.ClientID, clientID)
}

func TestMethod(t *testing.T) {
	ctx := context.Background()

	rsp, err := Get(ctx, LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = Head(ctx, LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = Post(ctx, LOCALURL+"post")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = Put(ctx, LOCALURL+"put")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = Patch(ctx, LOCALURL+"patch")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = Delete(ctx, LOCALURL+"delete")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = Options(ctx, LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	req := New()

	rsp, err = req.Get(ctx, LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = req.Head(ctx, LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = req.Post(ctx, LOCALURL+"post")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = req.Put(ctx, LOCALURL+"put")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = req.Patch(ctx, LOCALURL+"patch")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = req.Delete(ctx, LOCALURL+"delete")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	rsp, err = req.Options(ctx, LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
}

func TestSetClientKey(t *testing.T) {
	req := New()
	assert.Equal(t, req.ClientKey, "")
	req.SetClientKey(LOCALURL)
	assert.Equal(t, req.ClientKey, LOCALURL)
}

func TestSetHost(t *testing.T) {
	req := New()
	ctx := context.Background()

	host := req.Request.Host
	req.SetHost("likexian.com")
	assert.Equal(t, req.Request.Host, "likexian.com")
	assert.NotEqual(t, req.Request.Host, host)

	var h Host = "likexian.com"
	_, _ = req.Do(ctx, "GET", LOCALURL, h)
	assert.Equal(t, req.Request.Host, "likexian.com")
}

func TestSetHeader(t *testing.T) {
	req := New()
	ctx := context.Background()

	author := req.GetHeader("X-Author")
	assert.Equal(t, author, "")
	req.SetHeader("X-Author", "likexian")
	assert.Equal(t, req.GetHeader("X-Author"), "likexian")

	h1 := Header{
		"X-Version": Version(),
	}
	_, _ = req.Do(ctx, "GET", LOCALURL, h1)
	assert.Equal(t, req.GetHeader("X-Author"), "likexian")
	assert.Equal(t, req.GetHeader("X-Version"), Version())

	h2 := http.Header{
		"X-License": []string{License()},
	}
	_, _ = req.Do(ctx, "GET", LOCALURL, h2)
	assert.Equal(t, req.GetHeader("X-Author"), "likexian")
	assert.Equal(t, req.GetHeader("X-Version"), Version())
	assert.Equal(t, req.GetHeader("X-License"), License())
}

func TestSetUA(t *testing.T) {
	req := New()
	ua := req.GetHeader("User-Agent")
	assert.Equal(t, ua, fmt.Sprintf("GoKit XHTTP Client/%s", Version()))
	req.SetUA("HTTP Client by likexian")
	assert.Equal(t, req.GetHeader("User-Agent"), "HTTP Client by likexian")
}

func TestSetReferer(t *testing.T) {
	req := New()
	referer := req.GetHeader("referer")
	assert.Equal(t, referer, "")
	req.SetHeader("referer", LOCALURL)
	assert.Equal(t, req.GetHeader("referer"), LOCALURL)
	req.SetReferer(LOCALURL + "test")
	assert.Equal(t, req.GetHeader("referer"), LOCALURL+"test")
}

func TestGetHeader(t *testing.T) {
	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	for _, k := range []string{"Content-Type", "Date", "Server"} {
		h := rsp.GetHeader(k)
		assert.NotEqual(t, h, "")
	}
}

func TestSetClient(t *testing.T) {
	req := New()
	ctx := context.Background()
	rsp, err := req.Do(ctx, "GET", LOCALURL, &http.Client{})
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
}

func TestBytes(t *testing.T) {
	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	b, err := rsp.Bytes()
	assert.Nil(t, err)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "<")

	tracing := rsp.Tracing
	assert.NotEqual(t, tracing.Timestamp, "")
	assert.NotEqual(t, tracing.Nonce, "")

	rsp, err = req.Do(ctx, "GET", LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)

	b, err = rsp.Bytes()
	assert.Nil(t, err)
	assert.NotEqual(t, len(b), 0)
	assert.Equal(t, string(b[0:1]), "{")

	assert.NotEqual(t, rsp.Tracing.Timestamp, "")
	assert.NotEqual(t, rsp.Tracing.Nonce, tracing.Nonce)
	assert.Equal(t, rsp.Tracing.ClientID, tracing.ClientID)
	assert.NotEqual(t, rsp.Tracing.RequestID, tracing.RequestID)

	tracing = rsp.Tracing
	req = New()
	rsp, err = req.Do(ctx, "GET", LOCALURL+"status/404")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 404)
	assert.NotEqual(t, rsp.Tracing.ClientID, tracing.ClientID)
	assert.NotEqual(t, rsp.Tracing.RequestID, tracing.RequestID)
}

func TestString(t *testing.T) {
	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	s, err := rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	rsp, err = req.Do(ctx, "GET", LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "{")
}

func TestJSON(t *testing.T) {
	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	_, err = rsp.JSON()
	assert.NotNil(t, err)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	s, err := rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, s.Get("url").MustString(""), LOCALURL+"get")
}

func TestFile(t *testing.T) {
	defer func() {
		os.Remove("index.html")
		os.Remove("get.html")
		os.RemoveAll("tmp")
	}()

	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	ss, err := rsp.File()
	assert.Nil(t, err)
	fs, err := xfile.Size("index.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	ss, err = rsp.File("tmp/")
	assert.Nil(t, err)
	fs, err = xfile.Size("tmp/index.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	ss, err = rsp.File("get.html")
	assert.Nil(t, err)
	fs, err = xfile.Size("get.html")
	assert.Nil(t, err)
	assert.Equal(t, fs, ss)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"get")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	_, err = rsp.File("get.html")
	assert.NotNil(t, err)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"status/404")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 404)
	_, err = rsp.File("404.html")
	assert.NotNil(t, err)
}

func TestFollowRedirect(t *testing.T) {
	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL+"redirect/3")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	loc := rsp.GetHeader("Location")
	assert.Equal(t, loc, "")

	req.FollowRedirect(false)
	rsp, err = req.Do(ctx, "GET", LOCALURL+"redirect/3")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 302)
	loc = rsp.GetHeader("Location")
	assert.Equal(t, loc, "/redirect/2")

	req.FollowRedirect(true)
	rsp, err = req.Do(ctx, "GET", LOCALURL+"redirect/3")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	loc = rsp.GetHeader("Location")
	assert.Equal(t, loc, "")
}

func TestSetGzip(t *testing.T) {
	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	s, err := rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	req.SetGzip(false)
	rsp, err = req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")

	req.SetGzip(true)
	rsp, err = req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, rsp.StatusCode, 200)
	s, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, len(s), 0)
	assert.Equal(t, s[0:1], "<")
}

func TestSetVerifyTls(t *testing.T) {
	req := New()
	assert.False(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)

	req.SetVerifyTLS(false)
	assert.True(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)

	req.SetVerifyTLS(true)
	assert.False(t, req.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestSetKeepAliveTimeout(t *testing.T) {
	req := New()
	assert.False(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)

	req.SetKeepAliveTimeout(0)
	assert.True(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)

	req.SetKeepAliveTimeout(30)
	assert.False(t, req.Client.Transport.(*http.Transport).DisableKeepAlives)
}

func TestSetConnectTimeout(t *testing.T) {
	req := New()

	req.SetConnectTimeout(3)
	assert.Equal(t, req.GetTimeout().ConnectTimeout, 3)
}

func TestSetClientTimeout(t *testing.T) {
	req := New()

	req.SetClientTimeout(30)
	assert.Equal(t, req.Client.Timeout, time.Duration(30)*time.Second)
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
	req := New().SetProxy(func(req *http.Request) (*url.URL, error) {
		return url.ParseRequestURI("http://127.0.0.1:8080")
	})
	ctx := context.Background()
	_, err := req.Do(ctx, "GET", LOCALURL)
	assert.NotNil(t, err)
}

func TestSetProxyUrl(t *testing.T) {
	req := New().SetProxyURL("127.0.0.1:8080")
	ctx := context.Background()
	_, err := req.Do(ctx, "GET", LOCALURL)
	assert.NotNil(t, err)
}

func TestEnableCookie(t *testing.T) {
	// not enable cookies
	req := New()
	ctx := context.Background()

	req.FollowRedirect(false)
	rsp, err := req.Do(ctx, "GET", LOCALURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	// enable cookies
	req.EnableCookie(true)
	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	// delete cookies
	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies/delete?k=")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	// set cookie again
	req.EnableCookie(true)
	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)

	// disable cookies
	req.EnableCookie(false)
	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies/set/k/v")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"cookies")
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 0)

	// set cookies by args
	cookie := &http.Cookie{Name: "k", Value: "likexian"}
	req.EnableCookie(true)
	rsp, err = req.Do(ctx, "GET", LOCALURL, cookie)
	assert.Nil(t, err)
	defer rsp.Close()
	assert.Equal(t, len(req.Request.Cookies()), 1)
}

func TestQueryParam(t *testing.T) {
	req := New()
	ctx := context.Background()

	query := QueryParam{"k": "v"}
	_, err := req.Do(ctx, "GET", LOCALURL+"get", query)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"get?k=v")

	query = QueryParam{"a": "1", "b": 2, "c": 3}
	_, err = req.Do(ctx, "GET", req.Request.URL.String(), query)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"get?k=v&a=1&b=2&c=3")

	query = QueryParam{}
	_, err = req.Do(ctx, "GET", LOCALURL+"get", query)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"get")
}

func TestFormParam(t *testing.T) {
	req := New()
	ctx := context.Background()

	form := FormParam{"k": "v"}
	rsp, err := req.Do(ctx, "POST", LOCALURL+"post", form)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err := rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("form").Get("k.0").MustString(""), "v")

	form = FormParam{"a": "1", "b": 2, "c": 3}
	rsp, err = req.Do(ctx, "POST", req.Request.URL.String(), form)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err = rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("form").Get("a.0").MustString(""), "1")
	assert.Equal(t, json.Get("form").Get("b.0").MustString(""), "2")
	assert.Equal(t, json.Get("form").Get("c.0").MustString(""), "3")

	form = FormParam{}
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", form)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err = rsp.JSON()
	assert.Nil(t, err)
	m, _ := json.Get("form").Map()
	assert.Equal(t, m, map[string]interface{}{})

	data := map[string]interface{}{"a": "1", "b": 2, "c": 3}
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", FormParam(data))
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err = rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("form").Get("a.0").MustString(""), "1")
	assert.Equal(t, json.Get("form").Get("b.0").MustString(""), "2")
	assert.Equal(t, json.Get("form").Get("c.0").MustString(""), "3")
}

func TestValuesParam(t *testing.T) {
	req := New()
	ctx := context.Background()
	values := url.Values{"k": []string{"v"}}

	// url.Values as query string
	_, err := req.Do(ctx, "GET", LOCALURL+"get", values)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"get?k=v")

	// url.Values as form data
	rsp, err := req.Do(ctx, "POST", LOCALURL+"post", values)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err := rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("form").Get("k.0").MustString(""), "v")
}

func TestPostBody(t *testing.T) {
	req := New()
	ctx := context.Background()

	// Post string
	rsp, err := req.Do(ctx, "POST", LOCALURL+"post", "k=v")
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err := rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("form").Get("k.0").MustString(""), "v")

	// Post []byte
	rsp, err = req.Do(ctx, "POST", req.Request.URL.String(), []byte("a=1&b=2&c=3"))
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err = rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("form").Get("a.0").MustString(""), "1")
	assert.Equal(t, json.Get("form").Get("b.0").MustString(""), "2")
	assert.Equal(t, json.Get("form").Get("c.0").MustString(""), "3")

	// Post bytes.Buffer
	var b bytes.Buffer
	b.Write([]byte("k=v"))
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", b)
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err = rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("form").Get("k.0").MustString(""), "v")

	// Post json string
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", `{"k": "v"}`, Header{"Content-Type": "application/json"})
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	json, err = rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, json.Get("json").Get("k").MustString(""), "v")

	// Post map as json
	data := map[string]interface{}{"a": "1", "b": 2, "c": 3}
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", JSONParam(data))
	assert.Nil(t, err)
	assert.Equal(t, req.Request.URL.String(), LOCALURL+"post")
	j, err := rsp.JSON()
	assert.Nil(t, err)
	assert.Equal(t, j.Get("url").MustString(""), LOCALURL+"post")
	assert.Equal(t, j.Get("json.a").MustString(""), "1")
	assert.Equal(t, j.Get("json.b").MustInt(0), 2)
	assert.Equal(t, j.Get("json.c").MustInt(0), 3)
}

func TestPostFile(t *testing.T) {
	req := New()
	ctx := context.Background()

	// Test post one file
	rsp, err := req.Do(ctx, "POST", LOCALURL+"post", FormFile{"file": "../go.mod"})
	assert.Nil(t, err)
	defer rsp.Close()
	json, err := rsp.JSON()
	assert.Nil(t, err)
	assert.Contains(t, json.Get("headers.Content-Type.0").MustString(""), "multipart/form-data")
	assert.Contains(t, json.Get("file").Get("file").MustString(""), "module github.com/likexian/gokit")

	// Test post more files
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", FormFile{"file_0": "../go.mod"}, FormFile{"file_1": "../go.sum"})
	assert.Nil(t, err)
	defer rsp.Close()
	json, err = rsp.JSON()
	assert.Nil(t, err)
	assert.Contains(t, json.Get("headers.Content-Type.0").MustString(""), "multipart/form-data")
	assert.Contains(t, json.Get("file").Get("file_0").MustString(""), "module github.com/likexian/gokit")
	assert.Contains(t, json.Get("file").Get("file_1").MustString(""), "")

	// Test post file and form
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", FormParam{"k": "v"}, FormFile{"file": "../go.mod", "404": "404.md"})
	assert.Nil(t, err)
	defer rsp.Close()
	json, err = rsp.JSON()
	assert.Nil(t, err)
	assert.Contains(t, json.Get("headers.Content-Type.0").MustString(""), "multipart/form-data")
	assert.Contains(t, json.Get("file").Get("file").MustString(""), "module github.com/likexian/gokit")
	assert.Equal(t, json.Get("form").Get("k.0").MustString(""), "v")
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.AfterFunc(100*time.Millisecond, cancel)
	}()

	req := New()
	_, err := req.Do(ctx, "GET", LOCALURL+"sleep")
	assert.NotNil(t, err)
}

func TestSetRetries(t *testing.T) {
	req := New()
	ctx := context.Background()
	assert.Panic(t, func() { req.SetRetries() })

	// no retry (default)
	rsp, err := req.Do(ctx, "Get", "http://127.0.0.1:5555/")
	assert.NotNil(t, err)
	assert.Equal(t, rsp.Tracing.Retries, 0)

	// retry 3 times
	req.SetRetries(3)
	rsp, err = req.Do(ctx, "Get", "http://127.0.0.1:5555/")
	assert.NotNil(t, err)
	assert.Equal(t, rsp.Tracing.Retries, 3)

	// start http server after 3 second, then request shall success
	go func() {
		time.AfterFunc(100*time.Millisecond, func() {
			http.HandleFunc("/after3/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello!") })
			_ = http.ListenAndServe("127.0.0.1:5555", nil)
		})
	}()

	// retry until success, sleep 1 second per request
	req.SetRetries(-1, 100*time.Millisecond)
	rsp, err = req.Do(ctx, "Get", "http://127.0.0.1:5555/after3/")
	assert.Nil(t, err)
	defer rsp.Close()
	text, err := rsp.String()
	assert.Nil(t, err)
	assert.Equal(t, text, "Hello!")
	assert.Gt(t, rsp.Tracing.Retries, 0)
}

func TestDump(t *testing.T) {
	req := New()
	ctx := context.Background()

	req.SetDump(true, false)
	rsp, err := req.Do(ctx, "POST", LOCALURL+"post", "k=v")
	assert.Nil(t, err)
	defer rsp.Close()
	dump := rsp.Dump()
	assert.NotContains(t, string(dump[0]), "k=v")

	req.SetDump(true, true)
	rsp, err = req.Do(ctx, "POST", LOCALURL+"post", "k=v")
	assert.Nil(t, err)
	defer rsp.Close()
	dump = rsp.Dump()
	assert.Contains(t, string(dump[0]), "k=v")
}

func TestEnableCache(t *testing.T) {
	req := New()
	ctx := context.Background()

	rsp, err := req.Do(ctx, "GET", LOCALURL+"time")
	assert.Nil(t, err)
	defer rsp.Close()

	text, err := rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, text, "")

	newRsp, err := req.Do(ctx, "GET", LOCALURL+"time")
	assert.Nil(t, err)
	defer newRsp.Close()

	newText, err := newRsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, newText, text)
	assert.NotEqual(t, newRsp.Tracing.RequestID, rsp.Tracing.RequestID)

	// enable get cache
	req.EnableCache("GET", 300)

	rsp, err = req.Do(ctx, "GET", LOCALURL+"time")
	assert.Nil(t, err)
	defer rsp.Close()

	text, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, text, "")

	newRsp, err = req.Do(ctx, "GET", LOCALURL+"time")
	assert.Nil(t, err)
	defer newRsp.Close()

	newText, err = newRsp.String()
	assert.Nil(t, err)
	assert.Equal(t, newText, text)
	assert.Equal(t, newRsp.Tracing.RequestID, rsp.Tracing.RequestID)

	newRsp, err = req.Do(ctx, "GET", LOCALURL+"time", QueryParam{"q": "a"})
	assert.Nil(t, err)
	defer newRsp.Close()

	newText, err = newRsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, newText, text)
	assert.NotEqual(t, newRsp.Tracing.RequestID, rsp.Tracing.RequestID)

	// enable post cache
	req.EnableCache("post", 300)

	rsp, err = req.Do(ctx, "POST", LOCALURL+"time", QueryParam{"q": "a"}, FormParam{"d": "v", "x": "likexian"})
	assert.Nil(t, err)
	defer rsp.Close()

	text, err = rsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, text, "")

	newRsp, err = req.Do(ctx, "POST", LOCALURL+"time", QueryParam{"q": "a"}, FormParam{"d": "v", "x": "likexian"})
	assert.Nil(t, err)
	defer newRsp.Close()

	newText, err = newRsp.String()
	assert.Nil(t, err)
	assert.Equal(t, newText, text)
	assert.Equal(t, newRsp.Tracing.RequestID, rsp.Tracing.RequestID)

	newRsp, err = req.Do(ctx, "POST", LOCALURL+"time", QueryParam{"q": "a"},
		FormParam{"d": "v", "x": "likexian", "q": "a"})
	assert.Nil(t, err)
	defer newRsp.Close()

	newText, err = newRsp.String()
	assert.Nil(t, err)
	assert.NotEqual(t, newText, text)
	assert.NotEqual(t, newRsp.Tracing.RequestID, rsp.Tracing.RequestID)
}

func TestCheckClient(t *testing.T) {
	u, _ := url.Parse(LOCALURL)
	r := &http.Request{
		Header: http.Header{},
		Method: "POST",
		URL:    u,
	}

	err := CheckClient(r, "")
	assert.NotNil(t, err)

	r.Header.Set("X-Http-Gokit-Requestid", "test")
	err = CheckClient(r, "")
	assert.NotNil(t, err)

	r.Header.Set("X-Http-Gokit-Requestid", "test-test-test")
	err = CheckClient(r, "")
	assert.NotNil(t, err)

	r.Header.Set("X-Http-Gokit-Requestid", "1234-test-test")
	err = CheckClient(r, "")
	assert.NotNil(t, err)

	r.Header.Set("X-Http-Gokit-Requestid", fmt.Sprintf("%d-test-test", xtime.S()))
	err = CheckClient(r, "")
	assert.NotNil(t, err)

	r.Header.Set("X-Http-Gokit-Requestid", fmt.Sprintf("%d-1234-test", xtime.S()))
	err = CheckClient(r, "")
	assert.NotNil(t, err)

	req := New()
	ctx := context.Background()
	rsp, err := req.Do(ctx, "GET", LOCALURL)
	assert.Nil(t, err)
	defer rsp.Close()
	err = CheckClient(req.Request, "")
	assert.Nil(t, err)
}

func TestConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	ctx := context.Background()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := New()
			req.SetHeader("X-Test-Value", "Test")
			rsp, err := req.Do(ctx, "GET", LOCALURL)
			assert.Nil(t, err)
			defer rsp.Close()
			assert.Equal(t, rsp.StatusCode, 200)
			str, err := rsp.String()
			assert.Nil(t, err)
			assert.Equal(t, len(str), 128)
		}()
	}

	wg.Wait()
}

func TestGetClientIPs(t *testing.T) {
	u, _ := url.Parse(LOCALURL)
	r := &http.Request{
		RemoteAddr: "127.0.0.1:1234",
		Header:     http.Header{},
		Method:     "POST",
		URL:        u,
	}

	ips := GetClientIPs(r)
	assert.Equal(t, ips, []string{"127.0.0.1"})

	r.Header.Set("X-Real-Ip", "1.1.1.1")
	ips = GetClientIPs(r)
	assert.Equal(t, ips, []string{"1.1.1.1", "127.0.0.1"})

	r.Header.Set("X-Forwarded-For", "2.2.2.2")
	ips = GetClientIPs(r)
	assert.Equal(t, ips, []string{"1.1.1.1", "2.2.2.2", "127.0.0.1"})

	r.Header.Set("X-Forwarded-For", "2.2.2.2, 3.3.3.3")
	ips = GetClientIPs(r)
	assert.Equal(t, ips, []string{"1.1.1.1", "2.2.2.2", "3.3.3.3", "127.0.0.1"})
}

func ServerForTesting(listen string) string {
	defaultListenIP := "127.0.0.1"
	defaultListenPort := "8080"

	listen = strings.TrimSpace(strings.Replace(listen, " ", "", -1))
	listen = strings.Trim(listen, ":")
	if !strings.Contains(listen, ":") {
		if len(listen) == 0 {
			listen = fmt.Sprintf("%s:%s", defaultListenIP, defaultListenPort)
		} else if len(listen) < 5 {
			listen = fmt.Sprintf("%s:%s", defaultListenIP, listen)
		} else {
			listen = fmt.Sprintf("%s:%s", listen, defaultListenPort)
		}
	}

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, `<!DOCTYPE html><html><head><meta charset="UTF-8">`+
				`<title>HTTP Server For Testing</title></head><body>Hello Testing!</body></html>`)
		})
		http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			type Result struct {
				Args    url.Values  `json:"args"`
				Headers http.Header `json:"headers"`
				Origin  string      `json:"origin"`
				URL     string      `json:"url"`
			}
			result := Result{
				Args:    r.URL.Query(),
				Headers: r.Header,
				Origin:  strings.Split(r.RemoteAddr, ":")[0],
				URL:     fmt.Sprintf("http://%s%s", r.Host, r.URL.String()),
			}
			text, _ := xjson.Dumps(result)
			fmt.Fprint(w, text)
		})
		http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			type Result struct {
				Args    url.Values             `json:"args"`
				Form    url.Values             `json:"form"`
				JSON    map[string]interface{} `json:"json"`
				File    map[string]string      `json:"file"`
				Headers http.Header            `json:"headers"`
				Origin  string                 `json:"origin"`
				URL     string                 `json:"url"`
			}

			result := Result{
				Args:    r.URL.Query(),
				Headers: r.Header,
				Form:    url.Values{},
				JSON:    map[string]interface{}{},
				File:    map[string]string{},
				Origin:  strings.Split(r.RemoteAddr, ":")[0],
				URL:     fmt.Sprintf("http://%s%s", r.Host, r.URL.String()),
			}
			if r.Header.Get("Content-Type") == "application/json" {
				body, _ := ioutil.ReadAll(r.Body)
				json, _ := xjson.Loads(string(body))
				result.JSON, _ = json.Map()
			} else {
				err := r.ParseMultipartForm(32 << 20)
				if err != nil {
					result.Form = r.PostForm
				} else {
					result.Form = r.MultipartForm.Value
					for k, v := range r.MultipartForm.File {
						for _, f := range v {
							fd, err := f.Open()
							if err == nil {
								ss, _ := ioutil.ReadAll(fd)
								result.File[k] = string(ss)
							}
						}
					}
				}
			}
			text, _ := xjson.Dumps(result)
			fmt.Fprint(w, text)
		})
		http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
		})
		http.HandleFunc("/patch", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
		})
		http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
		})
		http.HandleFunc("/cookies", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			type Result struct {
				Cookies map[string]string `json:"cookies"`
			}
			result := Result{
				Cookies: map[string]string{},
			}
			for _, v := range r.Cookies() {
				result.Cookies[v.Name] = v.Value
			}
			text, _ := xjson.Dumps(result)
			fmt.Fprint(w, text)
		})
		http.HandleFunc("/cookies/set/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			l := r.URL.String()[13:]
			ls := strings.Split(l, "/")
			if len(ls) > 1 {
				cookie := http.Cookie{Name: ls[0], Value: ls[1], Path: "/"}
				http.SetCookie(w, &cookie)
			}
			http.Redirect(w, r, "/cookies", http.StatusFound)
		})
		http.HandleFunc("/cookies/delete", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			for k := range r.URL.Query() {
				cookie := http.Cookie{Name: k, Value: "", Path: "/", MaxAge: -1}
				http.SetCookie(w, &cookie)
			}
			http.Redirect(w, r, "/cookies", http.StatusFound)
		})
		http.HandleFunc("/redirect/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			s := "/get"
			l := r.URL.String()[10:]
			if len(l) > 0 {
				n, err := assert.ToInt64(l)
				if err == nil && n > 1 {
					s = fmt.Sprintf("/redirect/%d", n-1)
				}
			}
			http.Redirect(w, r, s, http.StatusFound)
		})
		http.HandleFunc("/status/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			s := 200
			l := r.URL.String()[8:]
			if len(l) > 0 {
				n, err := assert.ToInt64(l)
				if err == nil && n > 0 {
					s = int(n)
				}
			}
			w.WriteHeader(s)
		})
		http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%d", time.Now().UnixNano())
		})
		http.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			time.Sleep(1 * time.Second)
		})
		_ = http.ListenAndServe(listen, GzWrap(SetHeaderWrap(http.DefaultServeMux, Header{"Server": "Testing"})))
	}()

	req := New()
	for {
		_, err := req.Do(context.Background(), "GET", fmt.Sprintf("http://%s/", listen))
		if err == nil {
			break
		}
	}

	return fmt.Sprintf("http://%s/", listen)
}
