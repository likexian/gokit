# GoKit - xversion

Version kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xversion"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xversion)

## Example

```go
req := &CheckUpdateRequest{
    Product:       "test",
    Current:       "1.0.0",
    CacheFile:     "check_cache_file",
    CacheDuration: 1 * time.Hour,
    CheckPoint:    "https://check_url/",
}

ctx := context.Background()
rsp, err := req.Run(ctx)
if err != nil {
    panic(err)
} else {
    fmt.Println(rsp.Outdated)
}
```

## License

Copyright 2012-2023 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
