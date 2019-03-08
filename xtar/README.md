# GoKit - xtar

Tar kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xtar"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xtar)

## Example

### create a tar with gzip compress

    err := xtar.Create("likexian.tar.gz", "xtar.go", "xtar_test.go")
    if err != nil {
        fmt.Println("Create tar error:", err)
    }

### Extract a tar with gzip compress

    err := xtar.Extract("likexian.tar.gz", "tmp")
    if err != nil {
        fmt.Println("Extract tar error:", err)
    }

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
