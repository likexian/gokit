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

package xdaemon

import (
	"os"
	"syscall"

	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xos"
)

// Config storing config for daemon
type Config struct {
	Pid   string
	Log   string
	User  string
	Chdir string
}

// Version returns package version
func Version() string {
	return "0.7.1"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Daemon start to daemon
func (c *Config) Daemon() (err error) {
	var pid *xos.Pidx

	if c.Pid != "" {
		pid = xos.Pid(c.Pid)
		_, err := pid.Alive()
		if err == nil {
			return xos.ErrPidExists
		}
	}

	err = c.doDaemon()
	if err != nil {
		return
	}

	if c.Pid != "" {
		_, err = pid.Create()
		if err != nil {
			return
		}
	}

	if c.User != "" {
		err = xos.SetUser(c.User)
		if err != nil {
			return
		}
	}

	return
}

// Doing the daemon
func (c *Config) doDaemon() (err error) {
	syscall.Umask(0)

	if c.Chdir != "" {
		err = os.Chdir(c.Chdir)
		if err != nil {
			return
		}
	}

	if syscall.Getppid() == 1 {
		return
	}

	files := make([]*os.File, 3, 6)
	if c.Log != "" {
		fp, err := xfile.Append(c.Log)
		if err != nil {
			return err
		}
		files[0], files[1], files[2] = fp, fp, fp
	} else {
		files[0], files[1], files[2] = os.Stdin, os.Stdout, os.Stderr
	}

	dir, _ := os.Getwd()
	sysattrs := syscall.SysProcAttr{Setsid: true}
	prcattrs := os.ProcAttr{
		Dir:   dir,
		Env:   os.Environ(),
		Files: files,
		Sys:   &sysattrs,
	}

	proc, err := os.StartProcess(os.Args[0], os.Args, &prcattrs)
	if err != nil {
		return
	}

	err = proc.Release()
	if err != nil {
		return
	}
	os.Exit(0)

	return
}
