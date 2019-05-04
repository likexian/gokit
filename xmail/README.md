# GoKit - xmail

Mail kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xmail"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xmail)

## Example

### Send mail

```go
// Set the smtp info
// New("smtp server:port", "smtp username", "smtp password", "is html mail")
m := xmail.New("smtp.likexian.com:25", "i@likexian.com", "8Bd0a7681333214", true)

// Set email from
m.From = "i@likexian.com"

// Set send to
m.To = []string{"i@likexian.com"}

// Set mail subject
m.Subject = "Mailer Test"

// Set mail body
m.Body = "Hello World. This is xmail via github.com/likexian/gokit/xmail.<br /><img src=\"cid:xmail_test.jpg\" />"

// Add attachment
err := m.Attach("xmail_test.jpg")
if err != nil {
    panic(err)
}

// Do sending
err = m.Send()
if err != nil {
    panic(err)
}
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
