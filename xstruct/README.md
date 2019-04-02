# GoKit - xstruct

Struct kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xstruct"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xstruct)

## Example

```go
// Staff struct
type Staff struct {
    Id int64 `json:"id"`
    Name string  `json:"name"`
    Enabled bool  `json:"enabled"`
}

// New a object
s, err := xstruct.New(Staff{1, "likexian", true})
if err != nil {
    panic(err)
}

// ["Id", "Name", "Enabled"]
fmt.Println(s.Names())
// [1, "likexian", true]
fmt.Println(s.Values())
// [*Field] list
fmt.Println(s.Fields())

// set struct value
s.Set("Name", "kexian.li")
// get a field object
f, _ := s.Field("Name")
// set a field, same as s.Set(k, v)
f.Set("likexian")
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
