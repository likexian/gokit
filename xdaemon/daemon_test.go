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
    "testing"
)


func TestDaemon(t *testing.T) {
    c := Config {
        Pid:   "/tmp/test.pid",
        Log:   "/tmp/test.log",
        User:  "nobody",
        Chdir: "/",
    }

    err := c.Daemon()
    if err != nil {
        panic(err)
    }

    for {}
}
