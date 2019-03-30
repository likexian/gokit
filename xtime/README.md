# GoKit - xtime

Time kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xtime"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xtime)

## Example

### Get current timestamp

    // in second: 1552314204
    fmt.Println(xtime.S())

    // in millisecond: 1552314204000
    fmt.Println(xtime.Ms())

    // in string: 2019-03-11T22:23:24
    fmt.Println(xtime.String())

### Time string to timestamp

    // print 1552314204
    n, err := StrToTime("2019-03-11 22:23:24")
    if err != nil {
        fmt.Println(n)
    }

### Timestamp to time string

    // print 2019-03-11 22:23:24
    s := TimeToStr(1552314204)
    if err != nil {
        fmt.Println(n)
    }

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
