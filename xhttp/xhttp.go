/*
 * Copyright 2012-2023 Li Kexian
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
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xcache"
	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xhash"
	"github.com/likexian/gokit/xjson"
	"github.com/likexian/gokit/xrand"
	"github.com/likexian/gokit/xtime"
)

var (
	// DefaultRequest is default request
	DefaultRequest = New()

	// Caching is http request cache
	caching = xcache.New(xcache.MemoryCache)

	// supportMethod list all supported http method
	supportMethod = []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
	}
)

// Timeout storing timeout setting
type Timeout struct {
	ConnectTimeout        int
	TLSHandshakeTimeout   int
	ResponseHeaderTimeout int
	ExpectContinueTimeout int
	ClientTimeout         int
	KeepAliveTimeout      int
}

// Retries storing retry setting
type Retries struct {
	Times int
	Sleep time.Duration
}

// Dumping storing http dump setting
type Dumping struct {
	DumpHTTP bool
	DumpBody bool
}

// Caching storing cache method and ttl
type Caching struct {
	Method map[string]int64
}

// Request storing request data
type Request struct {
	ClientID  string
	Request   *http.Request
	Client    *http.Client
	ClientKey string
	Timeout   Timeout
	Caching   Caching
	Retries   Retries
	Dumping   Dumping
}

// Tracing storing tracing data
type Tracing struct {
	ClientID  string
	RequestID string
	Timestamp string
	Nonce     string
	SendTime  int64
	RecvTime  int64
	Retries   int
}

// Response storing response data
type Response struct {
	Method        string
	URL           *url.URL
	Response      *http.Response
	StatusCode    int
	ContentLength int64
	CacheKey      string
	Tracing       Tracing
	Dumping       [][]byte
}

// Host is http host
type Host string

// Header is http request header
type Header map[string]string

// QueryParam is query param map pass to xhttp
type QueryParam map[string]interface{}

// FormParam is form param map pass to xhttp
type FormParam map[string]interface{}

// JSONParam is json param map pass to xhttp
type JSONParam map[string]interface{}

// FormFile is form file for upload, formfield: filename
type FormFile map[string]string

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

// Version returns package version
func Version() string {
	return "0.19.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// New init a new xhttp client
func New() (r *Request) {
	timeout := Timeout{
		ConnectTimeout:        15,
		TLSHandshakeTimeout:   5,
		ResponseHeaderTimeout: 30,
		ExpectContinueTimeout: 5,
		ClientTimeout:         120,
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

	cache := Caching{
		Method: map[string]int64{},
	}

	r = &Request{
		ClientID:  xhash.Sha1("xhttp", xtime.Ns()).Hex(),
		Request:   request,
		Client:    client,
		ClientKey: "",
		Timeout:   timeout,
		Caching:   cache,
		Retries:   Retries{},
		Dumping:   Dumping{},
	}

	return
}

// Get do http GET request and returns response
func Get(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do(ctx, "GET", surl, args...)
}

// Head do http HEAD request and returns response
func Head(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do(ctx, "HEAD", surl, args...)
}

// Post do http POST request and returns response
func Post(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do(ctx, "POST", surl, args...)
}

// Put do http PUT request and returns response
func Put(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do(ctx, "PUT", surl, args...)
}

// Patch do http PATCH request and returns response
func Patch(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do(ctx, "PATCH", surl, args...)
}

// Delete do http DELETE request and returns response
func Delete(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do(ctx, "DELETE", surl, args...)
}

// Options do http OPTIONS request and returns response
func Options(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return DefaultRequest.Do(ctx, "OPTIONS", surl, args...)
}

// Get do http GET request and returns response
func (r *Request) Get(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return r.Do(ctx, "GET", surl, args...)
}

// Head do http HEAD request and returns response
func (r *Request) Head(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return r.Do(ctx, "HEAD", surl, args...)
}

// Post do http POST request and returns response
func (r *Request) Post(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return r.Do(ctx, "POST", surl, args...)
}

// Put do http PUT request and returns response
func (r *Request) Put(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return r.Do(ctx, "PUT", surl, args...)
}

// Patch do http PATCH request and returns response
func (r *Request) Patch(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return r.Do(ctx, "PATCH", surl, args...)
}

// Delete do http DELETE request and returns response
func (r *Request) Delete(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return r.Do(ctx, "DELETE", surl, args...)
}

// Options do http OPTIONS request and returns response
func (r *Request) Options(ctx context.Context, surl string, args ...interface{}) (s *Response, err error) {
	return r.Do(ctx, "OPTIONS", surl, args...)
}

// GetHeader return request header value by name
func (r *Request) GetHeader(name string) string {
	return r.Request.Header.Get(name)
}

// SetClientKey set key for signing requestid
func (r *Request) SetClientKey(key string) *Request {
	r.ClientKey = key
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

// SetVerifyTLS set http request tls verify
func (r *Request) SetVerifyTLS(verify bool) *Request {
	r.Client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = !verify
	return r
}

// SetKeepAliveTimeout set http keepalive timeout
func (r *Request) SetKeepAliveTimeout(timeout int) *Request {
	r.Timeout.KeepAliveTimeout = timeout
	r.SetTimeout(r.Timeout)
	return r
}

// SetConnectTimeout set http connect timeout
func (r *Request) SetConnectTimeout(timeout int) *Request {
	r.Timeout.ConnectTimeout = timeout
	r.SetTimeout(r.Timeout)
	return r
}

// SetClientTimeout set http client timeout
func (r *Request) SetClientTimeout(timeout int) *Request {
	r.Timeout.ClientTimeout = timeout
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
		Timeout:   time.Duration(r.Timeout.ConnectTimeout) * time.Second,
		KeepAlive: time.Duration(r.Timeout.KeepAliveTimeout) * time.Second,
	}).DialContext
	r.Client.Transport.(*http.Transport).TLSHandshakeTimeout =
		time.Duration(r.Timeout.TLSHandshakeTimeout) * time.Second
	r.Client.Transport.(*http.Transport).ResponseHeaderTimeout =
		time.Duration(r.Timeout.ResponseHeaderTimeout) * time.Second
	r.Client.Transport.(*http.Transport).ExpectContinueTimeout =
		time.Duration(r.Timeout.ExpectContinueTimeout) * time.Second
	r.Client.Timeout = time.Duration(r.Timeout.ClientTimeout) * time.Second
	return r
}

// GetTimeout get http request timeout
func (r *Request) GetTimeout() Timeout {
	return r.Timeout
}

// SetProxy set http request proxy
func (r *Request) SetProxy(proxy func(*http.Request) (*url.URL, error)) *Request {
	r.Client.Transport.(*http.Transport).Proxy = proxy
	return r
}

// SetProxyURL set http request proxy url
func (r *Request) SetProxyURL(proxy string) *Request {
	if !strings.HasPrefix(proxy, "http://") &&
		!strings.HasPrefix(proxy, "https://") &&
		!strings.HasPrefix(proxy, "socks5://") {
		proxy = "http://" + proxy
	}

	r.SetProxy(func(req *http.Request) (*url.URL, error) {
		return url.ParseRequestURI(proxy)
	})

	return r
}

// FollowRedirect set http request follow redirect
func (r *Request) FollowRedirect(follow bool) *Request {
	if follow {
		r.Client.CheckRedirect = nil
	} else {
		r.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return r
}

// EnableCookie set http request enable cookie
func (r *Request) EnableCookie(enable bool) *Request {
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

// EnableCache enable http client cache
func (r *Request) EnableCache(method string, ttl int64) *Request {
	method = strings.ToUpper(strings.TrimSpace(method))
	if assert.IsContains(supportMethod, method) {
		r.Caching.Method[method] = ttl
	}

	return r
}

// SetRetries set retry param
// int arg is setting retry times, time.Duration is setting retry sleep duration
// 0: no retry (default), -1: retry until success, > 1: retry x times
func (r *Request) SetRetries(args ...interface{}) *Request {
	if len(args) == 0 {
		panic("xhttp: the arguments is empty")
	}

	for i := 0; i < len(args); i++ {
		switch args[i].(type) {
		case int:
			r.Retries.Times = args[i].(int)
		case time.Duration:
			r.Retries.Sleep = args[i].(time.Duration)
		}
	}

	return r
}

// SetDump set http dump
func (r *Request) SetDump(dumpHTTP, dumpBody bool) *Request {
	r.Dumping.DumpHTTP = dumpHTTP
	r.Dumping.DumpBody = dumpBody
	return r
}

// Do send http request and return response
func (r *Request) Do(ctx context.Context, //nolint:cyclop
	method, surl string, args ...interface{}) (s *Response, err error) {
	r.Request.Host = ""
	r.Request.Header.Del("Cookie")
	r.Request.Header.Del("Content-Type")

	method = strings.ToUpper(strings.TrimSpace(method))
	if !assert.IsContains(supportMethod, method) {
		return nil, fmt.Errorf("xhttp: not supported method: %s", method)
	}

	r.Request.Method = method

	surl = strings.TrimSpace(surl)
	if surl == "" {
		return nil, fmt.Errorf("xhttp: no request url specify")
	}

	var formBody string
	var formParam param
	var queryParam param

	formFile := FormFile{}

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
		case *http.Client:
			r.Client = vv
		case *http.Cookie:
			r.Request.AddCookie(vv)
		case FormParam:
			formParam.Adds(vv)
		case QueryParam:
			queryParam.Adds(vv)
		case url.Values:
			if assert.IsContains([]string{"POST", "PUT", "PATCH"}, method) {
				formParam.Update(param{vv})
			} else {
				queryParam.Update(param{vv})
			}
		case JSONParam:
			formBody, err = xjson.Dumps(vv)
			if err != nil {
				return nil, fmt.Errorf("xhttp: encode json param failed: %w", err)
			}
			r.Request.Header.Set("Content-Type", "application/json")
		case string:
			formBody = vv
		case []byte:
			formBody = string(vv)
		case bytes.Buffer:
			formBody = vv.String()
		case FormFile:
			for k, v := range vv {
				formFile[k] = v
			}
		}
	}

	r.Request = r.Request.WithContext(ctx)
	r.Request.Body = nil
	r.Request.ContentLength = 0

	if assert.IsContains([]string{"POST", "PUT", "PATCH"}, method) {
		if len(formFile) > 0 {
			pr, pw := io.Pipe()
			bw := multipart.NewWriter(pw)
			go func() {
				for k, v := range formFile {
					fw, err := bw.CreateFormFile(k, v)
					if err != nil {
						continue
					}
					fd, err := os.Open(v)
					if err != nil {
						continue
					}
					_, err = io.Copy(fw, fd)
					fd.Close()
					if err != nil {
						continue
					}
				}
				for k, v := range formParam.Values {
					for _, vv := range v {
						_ = bw.WriteField(k, vv)
					}
				}
				bw.Close()
				pw.Close()
			}()
			r.SetHeader("Content-Type", bw.FormDataContentType())
			r.Request.Body = io.NopCloser(pr)
		} else {
			if !formParam.IsEmpty() {
				formBody += formParam.Encode()
			}
			if formBody != "" {
				r.Request.Body = io.NopCloser(bytes.NewReader([]byte(formBody)))
				r.Request.ContentLength = int64(len(formBody))
				if r.Request.Header.Get("Content-Type") == "" {
					r.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				}
			}
		}
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
		return nil, fmt.Errorf("xhttp: parse url failed: %w", err)
	}
	r.Request.URL = u

	s = &Response{
		Method: r.Request.Method,
		URL:    r.Request.URL,
		Tracing: Tracing{
			Timestamp: fmt.Sprintf("%d", xtime.S()),
			Nonce:     fmt.Sprintf("%d", xrand.IntRange(1000000, 9999999)),
			ClientID:  r.ClientID,
			Retries:   -1,
		},
	}

	startAt := xtime.Ms()
	defer func() {
		s.Tracing.SendTime = xtime.Ms() - startAt
	}()

	s.Tracing.RequestID = xhash.Sha1("xhttp", s.Tracing.Timestamp,
		s.Tracing.Nonce, s.Method, s.URL.Path, s.URL.RawQuery, r.ClientKey).Hex()
	r.Request.Header.Set("X-HTTP-GoKit-RequestId", fmt.Sprintf("%s-%s-%s", s.Tracing.Timestamp,
		s.Tracing.Nonce, s.Tracing.RequestID))

	cacheTTL, cacheEnabled := r.Caching.Method[s.Method]
	if cacheEnabled || r.Dumping.DumpHTTP {
		dumpBody := r.Dumping.DumpBody
		if cacheEnabled {
			dumpBody = true
		}
		d, err := httputil.DumpRequestOut(r.Request, dumpBody)
		if err == nil {
			s.Dumping = append(s.Dumping, d)
		}
	}

	if cacheEnabled {
		body := ""
		if len(s.Dumping) > 0 {
			d := strings.Split(string(s.Dumping[0]), "\r\n\r\n")
			if len(d) > 1 {
				body = d[1]
			}
			s.CacheKey = xhash.Sha1(s.Method, s.URL.String(), body).Hex()
			cacheVal := caching.Get(s.CacheKey)
			if cacheVal != nil {
				s = cacheVal.(*Response)
				return
			}
		}
	}

	for i := 0; r.Retries.Times == -1 || i <= r.Retries.Times; i++ {
		s.Tracing.Retries++
		s.Response, err = r.Client.Do(r.Request)
		if err == nil {
			break
		}
		if r.Retries.Sleep > 0 {
			time.Sleep(r.Retries.Sleep)
		}
	}

	if err == nil {
		s.StatusCode = s.Response.StatusCode
		s.ContentLength = s.Response.ContentLength
	}

	if r.Dumping.DumpHTTP {
		d, err := httputil.DumpResponse(s.Response, r.Dumping.DumpBody)
		if err == nil {
			s.Dumping = append(s.Dumping, d)
		}
	}

	if s.CacheKey != "" {
		_ = caching.Set(s.CacheKey, s, cacheTTL)
	}

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
		return 0, fmt.Errorf("xhttp: file %s is exists", fpath)
	}

	defer r.Response.Body.Close()
	if r.Response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("xhttp: bad status code: %d", r.Response.StatusCode)
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

	b, err = io.ReadAll(r.Response.Body)
	r.Response.Body.Close()

	if r.CacheKey != "" {
		r.Response.Body = io.NopCloser(bytes.NewBuffer(b))
	}

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

// JSON returns response body as *xjson.JSON
// For more please refer to gokit/xjson
func (r *Response) JSON() (*xjson.JSON, error) {
	s, err := r.String()
	if err != nil {
		return &xjson.JSON{}, err
	}

	return xjson.Loads(s)
}

// Dump returns http dump of request and response
// [bytes[request], bytes[response]]
func (r *Response) Dump() [][]byte {
	return r.Dumping
}

// CheckClient returns is a valid client request
// used by http server, it will check the requestId
func CheckClient(r *http.Request, ClientKey string) error {
	id := r.Header.Get("X-Http-Gokit-Requestid")
	if id == "" {
		return fmt.Errorf("xhttp: missing request id")
	}

	ids := strings.Split(id, "-")
	if len(ids) != 3 {
		return fmt.Errorf("xhttp: request id invalid")
	}

	tm, err := assert.ToInt64(ids[0])
	if err != nil {
		return fmt.Errorf("xhttp: time stamp invalid")
	}

	now := xtime.S()
	if tm-now > 300 || tm-now < -300 {
		return fmt.Errorf("xhttp: time stamp expired")
	}

	nonce, err := assert.ToInt64(ids[1])
	if err != nil {
		return fmt.Errorf("xhttp: request nonce invalid")
	}

	sum := xhash.Sha1("xhttp", tm, nonce, r.Method, r.URL.Path, r.URL.RawQuery, ClientKey).Hex()
	if sum != ids[2] {
		return fmt.Errorf("xhttp: hash value not matched")
	}

	return nil
}

// GetClientIPs returns all ips from http client
func GetClientIPs(r *http.Request) []string {
	ips := []string{}

	for _, h := range []string{"X-Real-Ip", "X-Forwarded-For"} {
		ip := r.Header.Get(h)
		if ip != "" {
			for _, v := range strings.Split(ip, ",") {
				v = strings.TrimSpace(v)
				if v != "" {
					ips = append(ips, v)
				}
			}
		}
	}

	ips = append(ips, strings.Split(r.RemoteAddr, ":")[0])

	return ips
}
