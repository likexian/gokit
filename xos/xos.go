/*
 * A toolkit for Golang development
 * https://www.likexian.com/
 *
 * Copyright 2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package xos

import (
	"fmt"
	"github.com/likexian/gokit/xfile"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// Version returns package version
func Version() string {
	return "0.1.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
}

// SetUser Set process user
func SetUser(user string) (err error) {
	uid, gid, err := LookupUser(user)
	if err != nil {
		return
	}

	err = SetGid(gid)
	if err == nil {
		err = SetUid(uid)
	}

	return
}

// LookupUser returns the uid and gid of user
func LookupUser(name string) (uid, gid int, err error) {
	u, err := user.Lookup(name)
	if err != nil {
		return
	}

	uid, err = strconv.Atoi(u.Uid)
	if err == nil {
		gid, err = strconv.Atoi(u.Gid)
	}

	return
}

// SetUid set the uid of process
func SetUid(uid int) (err error) {
	_, _, errno := syscall.RawSyscall(syscall.SYS_SETUID, uintptr(uid), 0, 0)
	if errno != 0 {
		err = errno
	}

	return
}

// SetGid set the gid of process
func SetGid(gid int) (err error) {
	_, _, errno := syscall.RawSyscall(syscall.SYS_SETGID, uintptr(gid), 0, 0)
	if errno != 0 {
		err = errno
	}

	return
}

// WritePid write pid to file path
func WritePid(path string) error {
	return xfile.WriteText(path, fmt.Sprintf("%d\n", os.Getpid()))
}
