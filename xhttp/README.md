# GoKit - xhttp

HTTP kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xhttp"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xhttp)

## Example

### Do a http request

    req, err := xhttp.New("GET", "https://httpbin.org/get")
    if err != nil {
        panic(err)
    }
    rsp, err := req.Do()
    if err != nil {
        panic(err)
    }
    defer rsp.Close()
    fmt.Println("response status code:", rsp.Response.StatusCode)

### Show http response body

    text, err := rsp.String()
    if err != nil {
        panic(err)
    }
    fmt.Println("response body:", text)

### Save response body to file (file download)

    size, err := rsp.File("get.json")
    if err != nil {
        panic(err)
    }
    fmt.Println("download size:", size)

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
