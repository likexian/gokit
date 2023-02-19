# GoKit - xhuman

Human kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xhuman"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xhuman)

## Example

### Get human string for bytes size

```go
// print 1024 * 1024 as 1MB
stringSize := xhuman.FormatByteSize(1024 * 1024)
fmt.Println("formated bytes size is:", stringSize)
```

### Get bytes size from human string

```go
// get 1024 * 1024 from 1MB
byteSize, err := xhuman.ParseByteSize("1MB")
if err != nil {
    fmt.Println("original bytes size is:", byteSize)
}
```

### Get comma split string for number

```go
// print 123456789123456 as "123,456,789,123,456"
s := xhuman.Comma(float64(123456789123456), 0)
if err != nil {
    fmt.Println("comma number:", s)
}
```

## License

Copyright 2012-2023 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
