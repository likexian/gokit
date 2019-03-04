# GoKit - file

File kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/file"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/file)

## Example

check file is exists

    exists := file.FileExists("/data/dev/gokit/LICENSE")
    if exists {
        fmt.Println("file is exists")
    } else {
        fmt.Println("file not exists")
    }

get file size

    size, err := file.FileSize("/data/dev/gokit/LICENSE")
    if err != nil {
        panic(err)
    } else {
        fmt.Println("file size is", size)
    }

write text to file

    err := WriteText("/tmp/not-exists-dir/LICENSE", "Copyright 2019, Li Kexian\n")
    if err != nil {
        panic(err)
    } else {
        fmt.Println("write to text successful")
    }

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
