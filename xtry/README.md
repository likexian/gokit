# GoKit - xtry

Retry kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xtry"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xtry)

## Example

```go
c := Config{
    Timeout: 5 * time.Minute,
    RetryDelay: 2 * time.Second,
}

ctx := context.Background()
err := c.Run(ctx, func(context.Context) error {
    return doSomething()
})
if err != nil {
    panic(err)
}
```

## LICENSE

Copyright 2012-2020 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
