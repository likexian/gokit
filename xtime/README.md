# GoKit - xtime

Time kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xtime"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xtime)

## Example

### Get current timestamp

```go
// print as int64 unix timestamp, example: 1552314204
fmt.Println(xtime.S())

// print as int64 unix timestamp of millisecond, example: 1552314204000
fmt.Println(xtime.Ms())

// print as YYYY-MM-DD HH:II:SS
fmt.Println(xtime.String())
```

### Time string to timestamp

```go
// print as int64 unix timestamp
n, err := xtime.StrToTime("2019-03-11 22:23:24")
if err != nil {
    fmt.Println(n)
}
```

### Timestamp to time string

```go
// print as YYYY-MM-DD HH:II:SS
s := xtime.TimeToStr(1552314204)
if err != nil {
    fmt.Println(n)
}
```

## License

Copyright 2012-2024 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
