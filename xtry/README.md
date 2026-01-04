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

## License

Copyright 2012-2026 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
