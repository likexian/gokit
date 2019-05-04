# GoKit - xlog

Log kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xlog"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xlog)

## Example

### Do logging to stderr

```go
log := xlog.New(os.Stderr, xlog.INFO)
log.Info("This is Info")
log.SetLevel(xlog.DEBUG)
log.Debug("This is Debug")
```

### Do logging to a file

```go
flog, err := xlog.File("test.log", xlog.DEBUG)
if err != nil {
    panic(err)
}
flog.Debug("This is Debug")
flog.Info("This is Info")
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
