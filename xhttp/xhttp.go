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

type Statics struct {
	SendTime int64
	RecvTime int64
}

type Request struct {
	Method            string
	URL               string
	Host              string
	Params            map[string][]string
	Files             map[string]string
	UserAgent         string
	Proxy             string
	Gzip              bool
	VerifyTls         bool
	FollowRedirect    bool
	EnableCookie      bool
	Retries           int
	ConnTimeout       int
	HandshakeTimeout  int
	ReadHeaderTimeout int
	ReadWriteTimeout  int
	KeepAliveTimeout  int
	Request           *http.Request
	Response          *http.Response
	Timestamp         string
	Nonce             string
	Requestid         string
	Processid         string
	Statics           Statics
	Debug             bool
}

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
	return "0.2.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

func New(method, surl string) (request *Request) {
	return &Request{
		Method:            method,
		URL:               surl,
		Host:              "",
		Params:            map[string][]string{},
		Files:             map[string]string{},
		UserAgent:         fmt.Sprintf("GoKit XHTTP Client/%s+(LiKexian)", Version()),
		Proxy:             "",
		Gzip:              true,
		VerifyTls:         true,
		FollowRedirect:    true,
		EnableCookie:      true,
		Retries:           0,
		ConnTimeout:       5,
		HandshakeTimeout:  5,
		ReadHeaderTimeout: 30,
		ReadWriteTimeout:  60,
		KeepAliveTimeout:  60,
		Request:           &http.Request{Header: http.Header{}},
		Response:          &http.Response{},
		Timestamp:         fmt.Sprintf("%d", xtime.S()),
		Nonce:             fmt.Sprintf("%d", xrand.IntRange(1000000, 9999999)),
		Requestid:         "",
		Processid:         xhash.Sha1(fmt.Sprintf("xhttp-%d", xtime.Ns())).Hex(),
		Statics:           Statics{},
		Debug:             false,
	}
}

func (r *Request) Next(method, surl string) (request *Request) {
	r.Method = method
	r.URL = surl

	r.Timestamp = fmt.Sprintf("%d", xtime.S())
	r.Nonce = fmt.Sprintf("%d", xrand.IntRange(1000000, 9999999))
	r.Requestid = ""

	return r
}

func (r *Request) SetHeader(key, value string) {
	r.Request.Header.Set(key, value)
}

func (r *Request) doRequest() (err error) {
	if r.Requestid != "" {
		return
	}

	startAt := xtime.Ms()
	defer func() {
		r.Statics.SendTime = xtime.Ms() - startAt
	}()

	r.Method = strings.ToUpper(strings.TrimSpace(r.Method))
	r.URL = strings.TrimSpace(r.URL)

	if !xslice.Contains(SUPPORT_METHOD, r.Method) {
		err = fmt.Errorf("xhttp: not supported method: %s", r.Method)
		return
	}

	u, err := url.Parse(r.URL)
	if err != nil {
		err = fmt.Errorf("xhttp: parse url failed: %s", err.Error())
		return
	}

	r.Request.Method = r.Method
	r.Request.URL = u

	if r.Host != "" {
		r.Request.Host = r.Host
	}

	if r.UserAgent != "" {
		r.Request.Header.Set("User-Agent", r.UserAgent)
	}

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   time.Duration(r.ConnTimeout) * time.Second,
			KeepAlive: time.Duration(r.KeepAliveTimeout) * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   time.Duration(r.HandshakeTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(r.ReadHeaderTimeout) * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: r.VerifyTls},
		DisableCompression:    !r.Gzip,
	}

	if r.KeepAliveTimeout == 0 {
		transport.DisableKeepAlives = true
	}

	if r.Proxy != "" {
		if !strings.HasPrefix(r.Proxy, "http://") &&
			!strings.HasPrefix(r.Proxy, "https://") &&
			!strings.HasPrefix(r.Proxy, "socks5://") {
			r.Proxy = "http://" + r.Proxy
		}
		transport.Proxy = func(req *http.Request) (*url.URL, error) {
			return url.ParseRequestURI(r.Proxy)
		}
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(r.ReadWriteTimeout) * time.Second,
	}

	if !r.FollowRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	if r.EnableCookie {
		client.Jar, _ = cookiejar.New(nil)
	}

	r.Requestid = xhash.Sha1(fmt.Sprintf("xhttp-%s-%s-%s-%s", r.Timestamp, r.Nonce, r.Method, r.URL)).Hex()
	r.Request.Header.Set("X-XHTTP-Timestamp", r.Timestamp)
	r.Request.Header.Set("X-XHTTP-Nonce", r.Nonce)
	r.Request.Header.Set("X-XHTTP-Requestid", r.Requestid)

	r.Response, err = client.Do(r.Request)

	return
}

func (r *Request) GetHeader(name string) (v string, err error) {
	err = r.doRequest()
	if err != nil {
		return
	}

	if name == "" {
		return
	}

	if v, ok := r.Response.Header[name]; ok {
		return v[0], nil
	}

	return
}

func (r *Request) File(fpath string) (size int64, err error) {
	fpath = strings.TrimSpace(fpath)
	if fpath == "" {
		_, fpath = filepath.Split(r.URL)
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

	err = r.doRequest()
	if err != nil {
		return
	}

	defer r.Response.Body.Close()
	if r.Response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status code: %d", r.Response.StatusCode)
	}

	startAt := xtime.Ms()
	defer func() {
		r.Statics.RecvTime = xtime.Ms() - startAt
	}()

	fd, err := xfile.New(fpath)
	if err != nil {
		return
	}

	defer fd.Close()
	size, err = io.Copy(fd, r.Response.Body)

	return
}

func (r *Request) Bytes() (b []byte, err error) {
	err = r.doRequest()
	if err != nil {
		return
	}

	startAt := xtime.Ms()
	defer func() {
		r.Statics.RecvTime = xtime.Ms() - startAt
	}()

	defer r.Response.Body.Close()
	b, err = ioutil.ReadAll(r.Response.Body)

	return
}

func (r *Request) String() (s string, err error) {
	b, err := r.Bytes()
	if err != nil {
		return
	}

	return string(b), nil
}
