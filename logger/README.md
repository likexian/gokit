# GoKit - logger

Log kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/logger"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/logger)

## Example

### Do logging to stderr

```go
log := logger.New(os.Stderr, logger.INFO)
log.Info("This is Info")
log.SetLevel(logger.DEBUG)
log.Debug("This is Debug")
```

### Do logging to a file

```go
flog, err := logger.File("test.log", logger.DEBUG)
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
