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
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// gzPool is gzip writer pool
var gzPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(ioutil.Discard)
	},
}

// gzResponseWriter is gzip ResponseWriter
type gzResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// Write write body byte
func (w *gzResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// WriteString write body string
func (w *gzResponseWriter) WriteString(s string) (int, error) {
	return w.Writer.Write([]byte(s))
}

// WriteHeader write http status code header
func (w *gzResponseWriter) WriteHeader(status int) {
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(status)
}

// GzWrap is http gzip transparent compression middleware
func GzWrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz := gzPool.Get().(*gzip.Writer)
		defer func() {
			gz.Reset(ioutil.Discard)
			gzPool.Put(gz)
		}()

		gz.Reset(w)
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(&gzResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})
}

// SetHeaderWrap is http set header middleware
func SetHeaderWrap(next http.Handler, header Header) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range header {
			w.Header().Set(k, v)
		}
		next.ServeHTTP(w, r)
	})
}
