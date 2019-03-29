# GoKit - xstring

String kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xstring"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xstring)

## Example

### Check string is all letter

    s := "abc123"
    ok := xstring.IsLetter(s)
    fmt.Println("IsLetter:", ok)

### Check string is a number

    s := "12345.67"
    ok := xstring.IsNumeric(s)
    fmt.Println("IsNumeric:", ok)

### Expand map value to template string

    t := "i am {name}, i have ${money}."
    m := map[string]interface{}{"name": "Li Kexian", "money": 100}
    s := xstring.Expand(t, m)
    fmt.Println(s)

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
