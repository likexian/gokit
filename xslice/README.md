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

### Get unique of string array

```go
array := xslice.Unique([]string{"a", "a", "b", "b", "b", "c"})
fmt.Println("new array:", array)
```

### Get unique of int array

```go
array := xslice.Unique([]int{0, 0, 1, 1, 1, 2, 2, 3})
fmt.Println("new array:", array)
```

## LICENSE

Copyright 2012-2020 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
