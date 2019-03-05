# GoKit - assert

Assert kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/assert"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/assert)

## Example

### assert panic

    func willItPanic() {
        panic("failed")
    }
    assert.Panic(t, willItPanic)

### assert err is nil

    fp, err := os.Open("/data/dev/gokit/LICENSE")
    assert.Nil(t, err)

### assert equal

    x := map[string]int{"a": 1, "b": 2}
    y := map[string]int{"a": 1, "b": 2}
    assert.Equal(t, x, y, "x shall equal to y")

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
