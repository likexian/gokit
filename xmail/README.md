# GoKit - xmail

Mail kits for Golang development.

## Features

- Light weight and Easy to use
- Attachment sending support
- Plain text sending support
- TLS sending support

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
// New("smtp server:port", "smtp username", "smtp password", isTLS)
m := New("smtp.likexian.com:465", "i@likexian.com", "8Bd0a7681333214", true)

// Set email from
m.From("i@likexian.com")

// Set send to
m.To("to@likexian.com")

// Set send cc
m.Cc("cc@likexian.com")

// Set send bcc
m.BCc("bcc@likexian.com")

// set mail content type
m.ContentType("text/html")

// Set mail subject
m.Content("Mailer Test", "xmail via github.com/likexian/gokit/xmail.<br /><img src=\"cid:xmail_test.jpg\" />")

// Add attachment
err := m.Attach("xmail_test.jpg")
if err != nil {
    panic(err)
}

err = m.Send()
if err != nil {
    panic(err)
}
```

## License

Copyright 2012-2023 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
