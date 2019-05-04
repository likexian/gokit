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
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// SetUser Set process user
func SetUser(user string) (err error) {
	uid, gid, err := LookupUser(user)
	if err != nil {
		return
	}

	err = setGid(gid)
	if err != nil {
		return
	}

	err = setUid(uid)
	if err != nil {
		return
	}

	return
}

// LookupUser find the user's uid and gid
func LookupUser(name string) (uid, gid int, err error) {
	u, err := user.Lookup(name)
	if err != nil {
		return
	}

	uid, err = strconv.Atoi(u.Uid)
	if err != nil {
		return
	}

	gid, err = strconv.Atoi(u.Gid)
	if err != nil {
		return
	}

	return
}

// setUid set the uid of process
func setUid(uid int) (err error) {
	_, _, errno := syscall.RawSyscall(syscall.SYS_SETUID, uintptr(uid), 0, 0)
	if errno != 0 {
		err = errno
	}

	return
}

// setGid set the gid of process
func setGid(gid int) (err error) {
	_, _, errno := syscall.RawSyscall(syscall.SYS_SETGID, uintptr(gid), 0, 0)
	if errno != 0 {
		err = errno
	}

	return
}

// Write the pid file
func writePid(pid string) (err error) {
	id := fmt.Sprintf("%d\n", os.Getpid())
	err = ioutil.WriteFile(pid, []byte(id), 0644)
	return
}
