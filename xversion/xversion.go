/*
 * Copyright 2012-2026 Li Kexian
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

package xversion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"

	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xhttp"
)

// CheckUpdateRequest is check update request
type CheckUpdateRequest struct {
	Product       string        `json:"product"`
	Current       string        `json:"current"`
	Arch          string        `json:"arch"`
	OS            string        `json:"os"`
	CacheFile     string        `json:"-"`
	CacheDuration time.Duration `json:"-"`
	CheckPoint    string        `json:"-"`
}

// CheckUpdateResponse is check update response
type CheckUpdateResponse struct {
	Product     string `json:"product"`
	Current     string `json:"current"`
	Latest      string `json:"latest"`
	Outdated    bool   `json:"outdated"`
	Emergency   bool   `json:"emergency"`
	DownloadURL string `json:"download_url"`
	ProductURL  string `json:"product_url"`
}

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
	return "Licensed under the Apache License 2.0"
}

// Run do check update
func (req *CheckUpdateRequest) Run(ctx context.Context) (rsp *CheckUpdateResponse, err error) {
	rsp = &CheckUpdateResponse{}

	if req.CacheDuration > 0 && req.CacheFile != "" {
		if s, e := os.Stat(req.CacheFile); e == nil {
			if s.ModTime().Add(req.CacheDuration).Unix() > time.Now().Unix() {
				data, e := xfile.Read(req.CacheFile)
				if e == nil {
					err = json.Unmarshal(data, rsp)
					if err == nil {
						return rsp, nil
					}
				}
			}
		}
	}

	if req.Arch == "" {
		req.Arch = runtime.GOARCH
	}

	if req.OS == "" {
		req.OS = runtime.GOOS
	}

	u, err := url.Parse(req.CheckPoint)
	if err != nil {
		return
	}

	q := u.Query()
	q.Set("product", req.Product)
	q.Set("current", req.Current)
	q.Set("arch", req.Arch)
	q.Set("os", req.OS)
	u.RawQuery = q.Encode()

	httpRsp, err := xhttp.Get(ctx, u.String(), xhttp.Header{"Accept": "application/json"})
	if err != nil {
		return
	}

	if httpRsp.StatusCode != http.StatusOK {
		return rsp, fmt.Errorf("xversion: bad status code: %d", httpRsp.StatusCode)
	}

	data, err := httpRsp.Bytes()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, rsp)
	if err != nil {
		return
	}

	_ = xfile.Write(req.CacheFile, data)

	return rsp, nil
}
