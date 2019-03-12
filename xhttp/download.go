/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// HttpDownload download file from url
func HttpDownload(url, saveName string, tlsSkipVerify bool) (size int64, cost int64, err error) {
	saveName = strings.TrimSpace(saveName)
	if saveName == "" {
		_, saveName = filepath.Split(url)
		if saveName == "" {
			err = fmt.Errorf("File name for saving is invalid")
			return
		}
	} else {
		dir, name := filepath.Split(saveName)
		if name == "" {
			err = fmt.Errorf("File name for saving is invalid")
			return
		}
		if dir != "" && !FileExists(dir) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return
			}
		}
	}

	if FileExists(saveName) {
		err = fmt.Errorf("File name %s is exists", saveName)
		return
	}

	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   time.Duration(10 * time.Second),
			KeepAlive: time.Duration(60 * time.Second),
		}).Dial,
		TLSHandshakeTimeout:   time.Duration(5 * time.Second),
		ResponseHeaderTimeout: time.Duration(5 * time.Second),
		ExpectContinueTimeout: time.Duration(1 * time.Second),
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: tlsSkipVerify},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(60 * time.Second),
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	tStart := time.Now().UnixNano()
	rsp, err := client.Do(req)
	if err != nil {
		return
	}

	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Bad status code: %d", rsp.StatusCode)
		return
	}

	fd, err := os.Create(saveName)
	if err != nil {
		return
	}

	defer fd.Close()
	size, err = io.Copy(fd, rsp.Body)
	cost = (time.Now().UnixNano() - tStart) / int64(time.Millisecond)

	return
}

func FileExists(fpath string) bool {
	_, err := os.Stat(fpath)
	return !os.IsNotExist(err)
}
