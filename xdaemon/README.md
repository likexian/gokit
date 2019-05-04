# GoKit - xdaemon

Daemon kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xdaemon"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xdaemon)

## Example

### Do deamon

```go
c := daemon.Config {
    Pid:   "/tmp/test.pid", // the pid file name
    Log:   "/tmp/test.log", // the log file name
    User:  "nobody",        // run daemon as user, if set, ROOT is required
    Chdir: "/",             // change working dir
}

err := c.Daemon()
if err != nil {
    panic(err)
}
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
