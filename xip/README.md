# GoKit - xip

IP kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xip"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xip)

## Example

### Check string is a valid ip

```go
ok := xip.IsIP("1.1.1.1")
fmt.Println("1.1.1.1 is a ip:", ok)
```

### IPv4 ip2long

```go
i, err := IPv4ToLong("1.1.1.1")
if err == nil {
    fmt.Println("1.1.1.1 ip2long is:", i)
}
```

### IPv4 long2ip

```go
ip := LongToIPv4(16843009)
fmt.Println("16843009 long2ip is:", ip)
```

## License

Copyright 2012-2023 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
