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

### Define a struct first

```go
// Define Staff struct
type Staff struct {
    Id int64 `json:"id"`
    Name string  `json:"name"`
    Enabled bool  `json:"enabled"`
}

// Init staff struct
staff := Staff{1, "likexian", true}
```

### Use as global functions

```go
// ["Id", "Name", "Enabled"]
names, _ := xstruct.Names(staff)

// [1, "likexian", true]
values, _ := xstruct.Values(staff)

// list all field as [*Field]
fields, _ := xstruct.Fields(staff)

// get struct field value
value, _ := xstruct.Field(staff, "Name").Value()

// set struct field value
xstruct.Set(staff, "Name", "kexian.li")
```

### Use as Interactive mode

```go
// create a xstruct object
s, err := xstruct.New(staff)
if err != nil {
    panic(err)
}

// ["Id", "Name", "Enabled"]
names := s.Names()

// [1, "likexian", true]
values := s.Values()

// list all field as [*Field]
fields := s.Fields()

// get struct field value
value := s.Field("Name").Value()

// set struct field value
s.Set("Name", "kexian.li")
```

## License

Copyright 2012-2023 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
