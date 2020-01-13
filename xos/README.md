# GoKit - xos

OS kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xos"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xos)

## Example

### Get uid and gid of nobody

```go
uid, gid, err := xos.LookupUser("nobody")
if err == nil {
    fmt.Println("uid=", uid, "gid=", gid)
}
```

### Set process user to nobody

```go
err := xos.SetUser("nobody")
if err != nil {
    fmt.Println("set user failed", err)
}
```

## LICENSE

Copyright 2012-2020 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
