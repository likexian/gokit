# GoKit - xos

OS kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xos"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xos)

## Example

### Get uid and gid of nobody

    uid, gid, err := xos.LookupUser("nobody")
    if err == nil {
        fmt.Println("uid=", uid, "gid=", gid)
    }

### Set process user to nobody

    err := xos.SetUser("nobody")
    if err != nil {
        fmt.Println("set user failed", err)
    }

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
