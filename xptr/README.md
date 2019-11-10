# GoKit - xptr

Pointer kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xptr"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xptr)

## Example

### Get pointer of int

```go
fmt.Println("&int:", xptr.Int(1))
```

### Get pointer of string

```go
fmt.Println("&string:", xptr.String("test"))
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
