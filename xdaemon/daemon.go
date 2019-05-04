/*
 * Go module for doing daemon
 * https://www.likexian.com/
 *
 * Copyright 2015-2016, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */


package daemon


import (
    "fmt"
    "os"
    "syscall"
    "io/ioutil"
)


type Config struct {
    Pid   string
    Log   string
    User  string
    Chdir string
}


func (c *Config) Daemon() (err error) {
    err = DoDaemon(c.Log, c.Chdir)
    if err != nil {
        return
    }

    if c.Pid != "" {
        err = WritePid(c.Pid)
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


func DoDaemon(log, chdir string) (err error) {
    syscall.Umask(0)

    if chdir != "" {
        os.Chdir(chdir)
    }

    if syscall.Getppid() == 1 {
        return
    }

    files := make([]*os.File, 3, 6)
    if log != "" {
        fp, err := os.OpenFile(log, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
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
        Dir: dir,
        Env: os.Environ(),
        Files: files,
        Sys: &sysattrs,
    }

    proc, err := os.StartProcess(os.Args[0], os.Args, &prcattrs)
    if err != nil {
        return
    }

    proc.Release()
    os.Exit(0)

    return
}


func WritePid(pid string) (err error) {
    id := fmt.Sprintf("%d\n", os.Getpid())
    err = ioutil.WriteFile(pid, []byte(id), 0644)
    return
}


func SetUser(user string) (err error) {
    uid, gid, err := LookupUser(user)
    if err != nil {
        return
    }

    err = Setgid(gid)
    if err != nil {
        return
    }

    err = Setuid(uid)
    if err != nil {
        return
    }

    return
}
