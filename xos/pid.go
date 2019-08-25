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

package xos

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/likexian/gokit/xfile"
)

// Pidx storing pid info
type Pidx struct {
	file string
}

var (
	// ErrPidExists is pid file exists
	ErrPidExists = errors.New("xos: process is running")
	// ErrPidLockFailed is pid lock failed
	ErrPidLockFailed = errors.New("xos: lock pid file failed")
)

// Pid init a new pid
func Pid(f string) *Pidx {
	return &Pidx{f}
}

// Create check create a new pid file
func (p *Pidx) Create() (int, error) {
	if xfile.Exists(p.file) {
		pid, err := p.Alive()
		if err == nil {
			return pid, ErrPidExists
		}
	}

	fd, err := xfile.New(p.file)
	if err != nil {
		return 0, err
	}

	defer fd.Close()
	pid := os.Getpid()
	_, err = fd.Write([]byte(strconv.Itoa(pid)))
	if err != nil {
		return pid, err
	}

	err = syscall.Flock(int(fd.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err == syscall.EWOULDBLOCK {
		return pid, ErrPidLockFailed
	}

	return pid, err
}

// Alive returns pid is alive
func (p *Pidx) Alive() (int, error) {
	pid, err := p.Value()
	if err != nil {
		return 0, err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return pid, err
	}

	err = process.Signal(os.Signal(syscall.Signal(0)))
	if err != nil {
		return pid, err
	}

	return pid, nil
}

// Value returns pid value
func (p *Pidx) Value() (int, error) {
	text, err := xfile.ReadText(p.file)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(text))
}
