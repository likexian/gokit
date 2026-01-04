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
c := xdaemon.Config {
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

## License

Copyright 2012-2026 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
