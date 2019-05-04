/*
 * Go module for doing daemon
 * https://www.likexian.com/
 *
 * Copyright 2015-2016, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */


package daemon


import(
    "errors"
    "strings"
    "strconv"
    "syscall"
    "io/ioutil"
)


func LookupUser(name string) (uid, gid int, err error) {
    text, err := ioutil.ReadFile("/etc/passwd")
    if err != nil {
        return
    }
    result := string(text)

    s_uid := ""
    s_gid := ""
    lines := strings.Split(result, "\n")
    for _, v := range lines {
        ls := strings.Split(v, ":")
        if ls[0] == name {
            s_uid = ls[2]
            s_gid = ls[3]
        }
    }

    if s_uid == "" || s_gid == "" {
        err = errors.New("User not exits")
        return
    }

    gid, err = strconv.Atoi(s_gid)
    if err != nil {
        return
    }

    uid, err = strconv.Atoi(s_uid)
    if err != nil {
        return
    }

    return
}


func Setuid(uid int) (err error) {
    _, _, errno := syscall.RawSyscall(syscall.SYS_SETUID, uintptr(uid), 0, 0)
    if errno != 0 {
        err = errno
    }

    return
}


func Setgid(gid int) (err error) {
    _, _, errno := syscall.RawSyscall(syscall.SYS_SETGID, uintptr(gid), 0, 0)
    if errno != 0 {
        err = errno
    }

    return
}
