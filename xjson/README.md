# GoKit - xjson

Json kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xjson"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xjson)

Visit simplejson docs on [GoDoc](https://godoc.org/github.com/likexian/simplejson-go)

## Example

### Dump struct data to json string

    type Status struct {
        Code    int64  `json:"code"`
        Message string `json:"message"`
    }

    status := Status{1, "Success"}
    text, err := xjson.Encode(status)
    if err == nil {
        fmt.Println(text)
    }

### Load string to json object

    text := `{"Code":1,"Message":"Success","Student":["Li Kexian"]}`
    json, err := xjson.Decode(text)
    if err == nil {
        fmt.Println(json.Get("Code").MustInt())
        fmt.Println(json.Get("Message").MustString())
        fmt.Println(json.Get("Student.0").MustString())
    }

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
