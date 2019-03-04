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

### assert err is nil

    fp, err := os.Open("/data/dev/gokit/LICENSE")
    assert.Nil(t, err)

### assert equal

    x := map[string]int{"a": 1, "b": 2}
    y := map[string]int{"a": 1, "b": 2}
    assert.Equal(t, x, y, "x shall equal to y")

### assert not equal

    x := map[string]interface{}{"a": 1, "b": 1}
    y := map[string]interface{}{"a": 1, "b": "1"}
    assert.NotEqual(t, x, y, "x shall not equal to y")

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
