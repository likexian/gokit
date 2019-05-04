# mailer.go

mailer-go is a simple Go module for sending email.

[![Build Status](https://secure.travis-ci.org/likexian/mailer-go.png)](https://secure.travis-ci.org/likexian/mailer-go)

## Overview

Help you sending email in the simple and easy way. Sending attachment is supported.

## Installation

    go get github.com/likexian/mailer-go

## Importing

    import (
        "github.com/likexian/mailer-go"
    )

## Documentation

type Message

    type Message struct {
        From        string
        To          []string
        Cc          []string
        Bcc         []string
        Subject     string
        Body        string
    }

Init a mailer

    func New(server, username, password string, ishtml bool) (m *Message)

Add attachement

    func (m *Message) Attach(fname string) (err error)

Do sending

    func (m *Message) Send() (err error)

## Example

    // Set the smtp info
    // New("smtp server:port", "smtp username", "smtp password", "is html mail")
    m := New("smtp.likexian.com:25", "i@likexian.com", "8Bd0a7681333214", true)

    // Set email from
    m.From = "i@likexian.com"

    // Set send to
    m.To = []string{"i@likexian.com"}

    // Set mail subject
    m.Subject = "Mailer Test"

    // Set mail body
    m.Body = "Hello World. This is mailer via github.com/likexian/mailer-go.<br /><img src=\"cid:mailer_test.jpg\" />"

    // Add attachment
    err := m.Attach("mailer_test.jpg")
    if err != nil {
        panic(err)
    }

    // Do sending
    err = m.Send()
    if err != nil {
        panic(err)
    }

## LICENSE

Copyright 2015-2018, Li Kexian

Apache License, Version 2.0

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)

## About

- [Li Kexian](https://www.likexian.com/)
