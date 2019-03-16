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
	"crypto/tls"
	"fmt"
	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xhash"
	"github.com/likexian/gokit/xrand"
	"github.com/likexian/gokit/xslice"
	"github.com/likexian/gokit/xtime"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Timeout storing timeout setting
type Timeout struct {
	ConnTimeout           int
	TLSHandshakeTimeout   int
	ResponseHeaderTimeout int
	ExpectContinueTimeout int
	ClientTimeout         int
	KeepAliveTimeout      int
}

// Request storing request data
type Request struct {
	Request   *http.Request
	Timeout   Timeout
	Client    *http.Client
	ClientId  string
	ClientKey string
	Retries   int
	Debug     bool
}

// Trace storing trace data
type Trace struct {
	ClientId  string
	RequestId string
	Timestamp string
	Nonce     string
	SendTime  int64
	RecvTime  int64
}

// Response storing response data
type Response struct {
	Method   string
	URL      *url.URL
	Response *http.Response
	Trace    Trace
}

// SUPPORT_METHOD list all supported http method
var (
	SUPPORT_METHOD = []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
	}
)

// Version returns package version
func Version() string {
	return "0.3.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// New init a new http request
func New(method, surl string) (r *Request, err error) {
	timeout := Timeout{
		ConnTimeout:           10,
		TLSHandshakeTimeout:   5,
		ResponseHeaderTimeout: 30,
		ExpectContinueTimeout: 5,
		ClientTimeout:         60,
		KeepAliveTimeout:      60,
	}

	request := &http.Request{
		Header: make(http.Header),
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: false},
			DisableCompression: false,
		},
	}

	r = &Request{
		Request:   request,
		Timeout:   timeout,
		Client:    client,
		ClientId:  UniqueId(fmt.Sprintf("%d", xtime.Ns())),
		ClientKey: "",
		Retries:   0,
		Debug:     false,
	}

	r.SetHeader("User-Agent", fmt.Sprintf("GoKit XHTTP Client/%s", Version()))
	err = r.Next(method, surl)

	return
}

// Next set next http request method and url
func (r *Request) Next(method, surl string) (err error) {
	r.Request.Host = ""
	r.Request.Header.Del("Cookie")

	err = r.SetMethod(method)
	if err != nil {
		return
	}

	err = r.SetURL(surl)
	if err != nil {
		return
	}

	return
}

// SetMethod set http request method
func (r *Request) SetMethod(method string) error {
	method = strings.ToUpper(strings.TrimSpace(method))
	if !xslice.Contains(SUPPORT_METHOD, method) {
		return fmt.Errorf("xhttp: not supported method: %s", method)
	}

	r.Request.Method = method

	return nil
}

// SetURL set http request url
func (r *Request) SetURL(surl string) error {
	surl = strings.TrimSpace(surl)
	if surl == "" {
		return fmt.Errorf("xhttp: no request url specify")
	}

	u, err := url.Parse(surl)
	if err != nil {
		return fmt.Errorf("xhttp: parse url failed: %s", err.Error())
	}

	r.Request.URL = u

	return nil
}

// SetClientKey set key for generating requestid
func (r *Request) SetClientKey(key string) {
	r.ClientKey = key
}

// SetHost set http request host
func (r *Request) SetHost(host string) {
	r.Request.Host = host
}

// SetHeader set http request header
func (r *Request) SetHeader(key, value string) {
	r.Request.Header.Set(key, value)
}

// SetUA set http request user-agent
func (r *Request) SetUA(ua string) {
	r.SetHeader("User-Agent", ua)
}

// SetReferer set http request referer
func (r *Request) SetReferer(referer string) {
	r.SetHeader("Referer", referer)
}

// SetGzip set http request gzip
func (r *Request) SetGzip(gzip bool) {
	r.Client.Transport.(*http.Transport).DisableCompression = !gzip
}

// SetVerifyTls set http request tls verify
func (r *Request) SetVerifyTls(verify bool) {
	r.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = !verify
}

// SetKeepAlive set http keepalive timeout
func (r *Request) SetKeepAlive(timeout int) {
	r.Timeout.KeepAliveTimeout = timeout
	r.SetTimeout(r.Timeout)
}

// GetTimeout get http request timeout
func (r *Request) GetTimeout() Timeout {
	return r.Timeout
}

// SetTimeout set http request timeout
func (r *Request) SetTimeout(timeout Timeout) {
	r.Timeout = timeout
	if r.Timeout.KeepAliveTimeout <= 0 {
		r.Client.Transport.(*http.Transport).DisableKeepAlives = true
	} else {
		r.Client.Transport.(*http.Transport).DisableKeepAlives = false
	}
	r.Client.Transport.(*http.Transport).DialContext = (&net.Dialer{
		Timeout:   time.Duration(r.Timeout.ConnTimeout) * time.Second,
		KeepAlive: time.Duration(r.Timeout.KeepAliveTimeout) * time.Second,
	}).DialContext
	r.Client.Transport.(*http.Transport).TLSHandshakeTimeout = time.Duration(r.Timeout.TLSHandshakeTimeout) * time.Second
	r.Client.Transport.(*http.Transport).ResponseHeaderTimeout = time.Duration(r.Timeout.ResponseHeaderTimeout) * time.Second
	r.Client.Transport.(*http.Transport).ExpectContinueTimeout = time.Duration(r.Timeout.ExpectContinueTimeout) * time.Second
	r.Client.Timeout = time.Duration(r.Timeout.ClientTimeout) * time.Second
}

// SetProxy set http request proxy
func (r *Request) SetProxy(proxy string) {
	if !strings.HasPrefix(proxy, "http://") &&
		!strings.HasPrefix(proxy, "https://") &&
		!strings.HasPrefix(proxy, "socks5://") {
		proxy = "http://" + proxy
	}
	r.Client.Transport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.ParseRequestURI(proxy)
	}
}

// SetFollowRedirect set http request follow redirect
func (r *Request) SetFollowRedirect(follow bool) {
	if follow {
		r.Client.CheckRedirect = nil
	} else {
		r.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
}

// SetEnableCookie set http request enable cookie
func (r *Request) SetEnableCookie(enable bool) {
	if enable {
		if r.Client.Jar == nil {
			r.Client.Jar, _ = cookiejar.New(nil)
		}
	} else {
		if r.Client.Jar != nil {
			r.Client.Jar = nil
		}
	}
}

// Do send http request and return response
func (r *Request) Do() (s *Response, err error) {
	s = &Response{
		Method: r.Request.Method,
		URL:    r.Request.URL,
		Trace: Trace{
			Timestamp: fmt.Sprintf("%d", xtime.S()),
			Nonce:     fmt.Sprintf("%d", xrand.IntRange(1000000, 9999999)),
			ClientId:  r.ClientId,
		},
	}

	startAt := xtime.Ms()
	defer func() {
		s.Trace.SendTime = xtime.Ms() - startAt
	}()

	s.Trace.RequestId = UniqueId(s.Trace.Timestamp, s.Trace.Nonce, s.Method, s.URL.String(), r.ClientKey)

	r.Request.Header.Set("X-XHTTP-Timestamp", s.Trace.Timestamp)
	r.Request.Header.Set("X-XHTTP-Nonce", s.Trace.Nonce)
	r.Request.Header.Set("X-XHTTP-RequestId", s.Trace.RequestId)

	s.Response, err = r.Client.Do(r.Request)

	return
}

// Close close response body
func (r *Response) Close() {
	r.Response.Body.Close()
}

// GetHeader return response header value by name
func (r *Response) GetHeader(name string) (v string, err error) {
	if v, ok := r.Response.Header[name]; ok {
		return v[0], nil
	}

	return
}

// File save response body to file
func (r *Response) File(paths ...string) (size int64, err error) {
	fpath := ""
	if len(paths) > 0 {
		fpath = paths[0]
	}

	fpath = strings.TrimSpace(fpath)
	if fpath == "" {
		_, fpath = filepath.Split(r.URL.String())
		if fpath == "" {
			fpath = "index.html"
		}
	} else {
		dir, name := filepath.Split(fpath)
		if name == "" {
			fpath = dir + "index.html"
		}
		if dir != "" && !xfile.Exists(dir) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return
			}
		}
	}

	if xfile.Exists(fpath) {
		return 0, fmt.Errorf("file %s is exists", fpath)
	}

	defer r.Response.Body.Close()
	if r.Response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status code: %d", r.Response.StatusCode)
	}

	startAt := xtime.Ms()
	defer func() {
		r.Trace.RecvTime = xtime.Ms() - startAt
	}()

	fd, err := xfile.New(fpath)
	if err != nil {
		return
	}

	defer fd.Close()
	size, err = io.Copy(fd, r.Response.Body)

	return
}

// Bytes returns response body as bytes
func (r *Response) Bytes() (b []byte, err error) {
	startAt := xtime.Ms()
	defer func() {
		r.Trace.RecvTime = xtime.Ms() - startAt
	}()

	defer r.Response.Body.Close()
	b, err = ioutil.ReadAll(r.Response.Body)

	return
}

// String returns response body as string
func (r *Response) String() (s string, err error) {
	b, err := r.Bytes()
	if err != nil {
		return
	}

	return string(b), nil
}

// UniqueId returns unique id of string list
func UniqueId(args ...string) string {
	s := "xhttp-" + strings.Join(args, "-")
	return xhash.Sha1(s).Hex()
}
