# GoKit - xjson

JSON kits for Golang development.

## Features

- Easy load to json and dump to string
- Load and dump with file is supported
- Modify the json data is simple
- One line retrieval with MustXXX
- Get by dot notation key is supported

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xjson"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xjson)

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

// Dump status to json string
j := xjson.New(status)
s, err := j.Dumps()
if err == nil {
    fmt.Println("JSON text is:", s)
}

// OR dumps using the easy way
s, err := xjson.Dumps(status)
if err == nil {
    fmt.Println("JSON text is:", s)
}
```

### Dump the map data to JSON string

```go
// Init a map data
data := map[string]interface{}{
    "code": 1,
    "message": "success",
    "result": {
        "Name": "Li Kexian"
    }
}

// Dump to string in the easy way
s, err := xjson.Dumps(status)
if err == nil {
    fmt.Println("JSON text is:", s)
}
```

### Load the JSON string

```go
// JSON strig
text := `{"Code": 1, "Message": "Success", "Result": {"Student": [{"Name": "Li Kexian"}]}}`

// Load json string
j, err := xjson.Loads(text)
if err == nil {
    fmt.Println("Code is:", j.Get("Code").MustInt(0))
    fmt.Println("Message is:", j.Get("Message").MustString(""))
    fmt.Println("First Student name is:", j.Get("Result.Student.0.Name").MustString("-"))
}
```

## License

Copyright 2012-2021 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
