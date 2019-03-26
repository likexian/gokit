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
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xhash"
	"github.com/likexian/gokit/xrand"
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
	Request  *http.Request
	Timeout  Timeout
	Client   *http.Client
	ClientId string
	SignKey  string
	Retries  int
	Debug    bool
}

// Host is http host
type Host string

// Header is http request header
type Header map[string]string

// QueryParam is query param map pass to xhttp
type QueryParam map[string]interface{}

// FormParam is form param map pass to xhttp
type FormParam map[string]interface{}

// param storing QueryParam and FormParam data set by Do
type param struct {
	url.Values
}

// getValues return values pointer
func (p *param) getValues() url.Values {
	if p.Values == nil {
		p.Values = make(url.Values)
	}

	return p.Values
}

func (p *param) Update(m param) {
	if m.Values == nil {
		return
	}
	vv := p.getValues()
	for m, n := range m.Values {
		for _, nn := range n {
			vv.Set(m, nn)
		}
	}
}

// Adds add map data to param
func (p *param) Adds(m map[string]interface{}) {
	if len(m) == 0 {
		return
	}

	vv := p.getValues()
	for k, v := range m {
		vv.Add(k, fmt.Sprint(v))
	}
}

// IsEmpty returns param is empty
func (p *param) IsEmpty() bool {
	return p.Values == nil
}

// Tracing storing tracing data
type Tracing struct {
	ClientId  string
	RequestId string
	Timestamp string
	Nonce     string
	SendTime  int64
	RecvTime  int64
	Retries   int64
}

// Response storing response data
type Response struct {
	Method   string
	URL      *url.URL
	Response *http.Response
	Tracing  Tracing
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
		"OPTIONS",
	}
)

// DefaultRequest is default request
var DefaultRequest = New()

// Version returns package version
func Version() string {
	return "0.7.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// New init a new xhttp client
func New() (r *Request) {
	timeout := Timeout{
		ConnTimeout:           10,
		TLSHandshakeTimeout:   5,
		ResponseHeaderTimeout: 30,
		ExpectContinueTimeout: 5,
		ClientTimeout:         60,
		KeepAliveTimeout:      60,
	}

	request := &http.Request{
		Header: http.Header{
			"User-Agent": []string{fmt.Sprintf("GoKit XHTTP Client/%s", Version())},
		},
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: false},
			DisableCompression: false,
		},
	}

	r = &Request{
		Request:  request,
		Timeout:  timeout,
		Client:   client,
		ClientId: UniqueId(fmt.Sprintf("%d", xtime.Ns())),
		SignKey:  "",
		Retries:  0,
		Debug:    false,
	}

	return
}

// Get do http GET request and returns response
func Get(surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do("GET", surl, args...)
}

// Head do http HEAD request and returns response
func Head(surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do("HEAD", surl, args...)
}

// Post do http POST request and returns response
func Post(surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do("POST", surl, args...)
}

// Put do http PUT request and returns response
func Put(surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do("PUT", surl, args...)
}

// Patch do http PATCH request and returns response
func Patch(surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do("PATCH", surl, args...)
}

// Delete do http DELETE request and returns response
func Delete(surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do("DELETE", surl, args...)
}

// Options do http OPTIONS request and returns response
func Options(surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do("OPTIONS", surl, args...)
}

// Get do http GET request and returns response
func (r *Request) Get(surl string, args ...interface{}) (s *Response, err error) {
	return r.Do("GET", surl, args...)
}

// Head do http HEAD request and returns response
func (r *Request) Head(surl string, args ...interface{}) (s *Response, err error) {
	return r.Do("HEAD", surl, args...)
}

// Post do http POST request and returns response
func (r *Request) Post(surl string, args ...interface{}) (s *Response, err error) {
	return r.Do("POST", surl, args...)
}

// Put do http PUT request and returns response
func (r *Request) Put(surl string, args ...interface{}) (s *Response, err error) {
	return r.Do("PUT", surl, args...)
}

// Patch do http PATCH request and returns response
func (r *Request) Patch(surl string, args ...interface{}) (s *Response, err error) {
	return r.Do("PATCH", surl, args...)
}

// Delete do http DELETE request and returns response
func (r *Request) Delete(surl string, args ...interface{}) (s *Response, err error) {
	return r.Do("DELETE", surl, args...)
}

// Options do http OPTIONS request and returns response
func (r *Request) Options(surl string, args ...interface{}) (s *Response, err error) {
	return r.Do("OPTIONS", surl, args...)
}

// GetHeader return request header value by name
func (r *Request) GetHeader(name string) string {
	return r.Request.Header.Get(name)
}

// SetSignKey set key for signing requestid
func (r *Request) SetSignKey(key string) *Request {
	r.SignKey = key
	return r
}

// SetHost set http request host
func (r *Request) SetHost(host string) *Request {
	r.Request.Host = host
	return r
}

// SetHeader set http request header
func (r *Request) SetHeader(key, value string) *Request {
	r.Request.Header.Set(key, value)
	return r
}

// SetUA set http request user-agent
func (r *Request) SetUA(ua string) *Request {
	r.SetHeader("User-Agent", ua)
	return r
}

// SetReferer set http request referer
func (r *Request) SetReferer(referer string) *Request {
	r.SetHeader("Referer", referer)
	return r
}

// SetGzip set http request gzip
func (r *Request) SetGzip(gzip bool) *Request {
	r.Client.Transport.(*http.Transport).DisableCompression = !gzip
	return r
}

// SetVerifyTls set http request tls verify
func (r *Request) SetVerifyTls(verify bool) *Request {
	r.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = !verify
	return r
}

// SetKeepAlive set http keepalive timeout
func (r *Request) SetKeepAlive(timeout int) *Request {
	r.Timeout.KeepAliveTimeout = timeout
	r.SetTimeout(r.Timeout)
	return r
}

// SetTimeout set http request timeout
func (r *Request) SetTimeout(timeout Timeout) *Request {
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
	return r
}

// GetTimeout get http request timeout
func (r *Request) GetTimeout() Timeout {
	return r.Timeout
}

// SetProxy set http request proxy
func (r *Request) SetProxy(proxy string) *Request {
	if !strings.HasPrefix(proxy, "http://") &&
		!strings.HasPrefix(proxy, "https://") &&
		!strings.HasPrefix(proxy, "socks5://") {
		proxy = "http://" + proxy
	}

	r.Client.Transport.(*http.Transport).Proxy = func(req *http.Request) (*url.URL, error) {
		return url.ParseRequestURI(proxy)
	}

	return r
}

// SetFollowRedirect set http request follow redirect
func (r *Request) SetFollowRedirect(follow bool) *Request {
	if follow {
		r.Client.CheckRedirect = nil
	} else {
		r.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return r
}

// SetEnableCookie set http request enable cookie
func (r *Request) SetEnableCookie(enable bool) *Request {
	if enable {
		if r.Client.Jar == nil {
			r.Client.Jar, _ = cookiejar.New(nil)
		}
	} else {
		if r.Client.Jar != nil {
			r.Client.Jar = nil
		}
	}

	return r
}

// Do send http request and return response
func (r *Request) Do(method, surl string, args ...interface{}) (s *Response, err error) {
	r.Request.Host = ""
	r.Request.Header.Del("Cookie")

	method = strings.ToUpper(strings.TrimSpace(method))
	if !assert.IsContains(SUPPORT_METHOD, method) {
		return nil, fmt.Errorf("xhttp: not supported method: %s", method)
	}

	r.Request.Method = method

	surl = strings.TrimSpace(surl)
	if surl == "" {
		return nil, fmt.Errorf("xhttp: no request url specify")
	}

	var formParam param
	var queryParam param

	for _, v := range args {
		switch vv := v.(type) {
		case Host:
			r.SetHost(string(vv))
		case Header:
			for k, v := range vv {
				r.SetHeader(k, v)
			}
		case http.Header:
			for k, v := range vv {
				for _, vv := range v {
					r.SetHeader(k, vv)
				}
			}
		case FormParam:
			formParam.Adds(vv)
		case QueryParam:
			queryParam.Adds(vv)
		case url.Values:
			if assert.IsContains([]string{"POST", "PUT"}, method) {
				formParam.Update(param{vv})
			} else {
				queryParam.Update(param{vv})
			}
		}
	}

	if !formParam.IsEmpty() {
		q := formParam.Encode()
		r.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(q)))
		r.Request.ContentLength = int64(len(q))
		r.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	} else {
		r.Request.Body = nil
		r.Request.ContentLength = 0
	}

	if !queryParam.IsEmpty() {
		q := queryParam.Encode()
		if strings.Contains(surl, "?") {
			surl += "&" + q
		} else {
			surl += "?" + q
		}
	}

	u, err := url.Parse(surl)
	if err != nil {
		return nil, fmt.Errorf("xhttp: parse url failed: %s", err.Error())
	}

	r.Request.URL = u

	s = &Response{
		Method: r.Request.Method,
		URL:    r.Request.URL,
		Tracing: Tracing{
			Timestamp: fmt.Sprintf("%d", xtime.S()),
			Nonce:     fmt.Sprintf("%d", xrand.IntRange(1000000, 9999999)),
			ClientId:  r.ClientId,
		},
	}

	startAt := xtime.Ms()
	defer func() {
		s.Tracing.SendTime = xtime.Ms() - startAt
	}()

	s.Tracing.RequestId = UniqueId(s.Tracing.Timestamp, s.Tracing.Nonce, s.Method, s.URL.String(), r.SignKey)
	r.Request.Header.Set("X-XHTTP-RequestId", fmt.Sprintf("%s-%s-%s", s.Tracing.Timestamp,
		s.Tracing.Nonce, s.Tracing.RequestId))

	s.Response, err = r.Client.Do(r.Request)

	return
}

// Close close response body
func (r *Response) Close() {
	r.Response.Body.Close()
}

// GetHeader return response header value by name
func (r *Response) GetHeader(name string) string {
	return r.Response.Header.Get(name)
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
		r.Tracing.RecvTime = xtime.Ms() - startAt
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
		r.Tracing.RecvTime = xtime.Ms() - startAt
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
