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

package xversion

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
	"github.com/likexian/gokit/xhttp"
	"github.com/likexian/gokit/xjson"
)

var (
	check_cache_listen = ""
	check_cache_file   = "check.cache"

	check_cache_req = &CheckUpdateRequest{
		Product:       "test",
		Current:       "1.0.0",
		CacheFile:     check_cache_file,
		CacheDuration: 1 * time.Hour,
		CheckPoint:    "",
	}

	check_cache_rsp = &CheckUpdateResponse{
		Product:   "test",
		Current:   "1.0.0",
		Latest:    "1.0.1",
		Outdated:  true,
		Emergency: true,
	}
)

func init() {
	check_cache_listen = ServerForTesting()
}

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestCheckUpdate(t *testing.T) {
	defer os.Remove(check_cache_file)

	ctx := context.Background()
	_, err := check_cache_req.Run(ctx)
	assert.NotNil(t, err)

	check_cache_req.CheckPoint = fmt.Sprintf("http://%s/todo/check", "%s")
	_, err = check_cache_req.Run(ctx)
	assert.NotNil(t, err)

	check_cache_req.CheckPoint = fmt.Sprintf("http://%s/todo/check", check_cache_listen)
	_, err = check_cache_req.Run(ctx)
	assert.NotNil(t, err)

	check_cache_req.CheckPoint = fmt.Sprintf("http://%s/update/nofound", check_cache_listen)
	_, err = check_cache_req.Run(ctx)
	assert.NotNil(t, err)

	check_cache_req.CheckPoint = fmt.Sprintf("http://%s/update/check", check_cache_listen)
	rsp, err := check_cache_req.Run(ctx)
	assert.Nil(t, err)
	assert.Equal(t, rsp, check_cache_rsp)

	rsp, err = check_cache_req.Run(ctx)
	assert.Nil(t, err)
	assert.Equal(t, rsp, check_cache_rsp)
}

func ServerForTesting() string {
	listen := "127.0.0.1:8080"

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
		http.HandleFunc("/update/nofound",
			func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNotFound) })
		http.HandleFunc("/update/check", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			text, _ := xjson.Dumps(check_cache_rsp)
			fmt.Fprint(w, text)
		})
		_ = http.ListenAndServe(listen, http.DefaultServeMux)
	}()

	req := xhttp.New()
	for {
		_, err := req.Do(context.Background(), "GET", fmt.Sprintf("http://%s/", listen))
		if err == nil {
			break
		}
	}

	return listen
}
