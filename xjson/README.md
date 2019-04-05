# GoKit - xjson

Json kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xjson"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xjson)

Visit simplejson docs on [GoDoc](https://godoc.org/github.com/likexian/simplejson-go)

## Example

### Dump the struct data to JSON string

```go
// Define Status struct
type Status struct {
    Code    int64  `json:"code"`
    Message string `json:"message"`
}

// Init status
status := Status{1, "Success"}

// Encode as json string
s, err := xjson.Encode(status)
if err == nil {
    fmt.Println("Json text is:", s)
}
```

### Load the JSON string

```go
// Json strig
text := `{"Code": 1, "Message": "Success", "Result": {"Student": [{"Name": "Li Kexian"}]}}`

// Decode json string
j, err := xjson.Decode(text)
if err == nil {
    fmt.Println("Code is:", j.Get("Code").MustInt(0))
    fmt.Println("Message is:", j.Get("Message").MustString(""))
    fmt.Println("First Student name is:", j.Get("Result.Student.0.Name").MustString("-"))
}
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
