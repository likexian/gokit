/*
 * Go module for doing daemon
 * https://www.likexian.com/
 *
 * Copyright 2015-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package daemon

import (
	"os"
	"syscall"
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
	return "0.5.1"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// Daemon start to daemon
func (c *Config) Daemon() (err error) {
	err = c.doDaemon()
	if err != nil {
		return
	}

	if c.Pid != "" {
		err = writePid(c.Pid)
		if err != nil {
			return
		}
	}

	if c.User != "" {
		err = SetUser(c.User)
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
		os.Chdir(c.Chdir)
	}

	if syscall.Getppid() == 1 {
		return
	}

	files := make([]*os.File, 3, 6)
	if c.Log != "" {
		fp, err := os.OpenFile(c.Log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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

	proc.Release()
	os.Exit(0)

	return
}
