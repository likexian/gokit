# GoKit - xrand

Rand kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xrand"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xrand)

## Example

### Rand int between 0 and 10000

```go
n := xrand.Int(10000)
fmt.Println("rand int between 0 and 10000 is:", n)
```

### Rand int between 1000 and 10000

```go
n := xrand.IntRange(1000, 10000)
fmt.Println("rand int between 1000 and 10000 is:", n)
```

### Rand bytes with length of 10

```go
b, err := xrand.Bytes(10)
if err != nil {
    fmt.Println("rand bytes:", b)
}
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
