# GoKit - xtar

Tar kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xtar"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xtar)

## Example

### create a tar with gzip compress

```go
err := xtar.Create("likexian.tar.gz", "xtar.go", "xtar_test.go")
if err != nil {
    fmt.Println("Create tar error:", err)
}
```

### Extract a tar with gzip compress

```go
err := xtar.Extract("likexian.tar.gz", "tmp")
if err != nil {
    fmt.Println("Extract tar error:", err)
}
```

## License

Copyright 2012-2022 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
