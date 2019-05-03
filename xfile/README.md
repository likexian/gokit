# GoKit - xfile

File kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xfile"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xfile)

## Example

### check file is exists

```go
exists := xfile.Exists("/data/dev/gokit/LICENSE")
if exists {
    fmt.Println("file is exists")
} else {
    fmt.Println("file not exists")
}
```

### get file size

```go
size, err := xfile.Size("/data/dev/gokit/LICENSE")
if err != nil {
    panic(err)
} else {
    fmt.Println("file size is", size)
}
```

### write text to file

```go
err := xfile.WriteText("/tmp/not-exists-dir/LICENSE", "Copyright 2012-2019 Li Kexian\n")
if err != nil {
    panic(err)
} else {
    fmt.Println("write to file successful")
}
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
