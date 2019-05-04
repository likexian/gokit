# daemon.go

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/daemon-go?status.svg)](https://godoc.org/github.com/likexian/daemon-go)
[![Build Status](https://travis-ci.org/likexian/daemon-go.svg?branch=master)](https://travis-ci.org/likexian/daemon-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/daemon-go)](https://goreportcard.com/report/github.com/likexian/daemon-go)

daemon-go is a Go module for doing daemon.

## Overview

A simple golang module for doing daemon.

*Work for \*nix ONLY*

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

Copyright 2015-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
