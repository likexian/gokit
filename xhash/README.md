# GoKit - xhash

Hash kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xhash"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xhash)

## Example

### Get md5 of string

```go
h := xhash.Md5("12345678")
fmt.Println(h.Hex())
fmt.Println(h.B64())
```

## Get Hmac Md5 of string

```go
h := xhash.HmacMd5("key", "12345678")
fmt.Println(h.Hex())
fmt.Println(h.B64())
```

### Get md5 of file

```go
h, err := xhash.FileMd5("xhash.go")
if err != nil {
    panic(err)
}
fmt.Println(h.Hex())
fmt.Println(h.B64())
```

## LICENSE

Copyright 2012-2020 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
