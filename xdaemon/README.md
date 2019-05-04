# daemon.go

daemon-go is a Go module for doing daemon.

[![Build Status](https://secure.travis-ci.org/likexian/daemon-go.png)](https://secure.travis-ci.org/likexian/daemon-go)

## Overview

A perfect golang module for doing daemon.

*Work for *nix ONLY*

## Installation

    go get github.com/likexian/daemon-go

## Importing

    import (
        "github.com/likexian/daemon-go"
    )

## Documentation

type Config

    type Config struct {
        Pid   string
        Log   string
        User  string
        Chdir string
    }

Do daemon

    func (c *Config) Daemon() (err error)

## Example

    c := daemon.Config {
        Pid:   "/tmp/test.pid", // the pid file name
        Log:   "/tmp/test.log", // the log file name
        User:  "nobody",        // run daemon as user, if set, ROOT is required
        Chdir: "/",             // change working dir
    }

    err := c.Daemon()
    if err != nil {
        panic(err)
    }

## LICENSE

Copyright 2015-2016, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)
