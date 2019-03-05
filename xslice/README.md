# GoKit - xslice

Slice kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xslice"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xslice)

## Example

### check string in array

    ok := xslice.Contains([]string{"a", "b", "c"}, "b")
    if ok {
        fmt.Println("value in array")
    } else {
        fmt.Println("value not in array")
    }

### check string in interface array

    ok := xslice.Contains([]interface{}{0, "1", 2}, "1")
    if ok {
        fmt.Println("value in array")
    } else {
        fmt.Println("value not in array")
    }

### check object in struct array

    ok := xslice.Contains([]A{A{0, 1}, A{1, 2}, A{1, 3}}, A{1, 2})
    if ok {
        fmt.Println("value in array")
    } else {
        fmt.Println("value not in array")
    }

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
