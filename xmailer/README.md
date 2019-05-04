# mailer.go

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/mailer-go?status.svg)](https://godoc.org/github.com/likexian/mailer-go)
[![Build Status](https://travis-ci.org/likexian/mailer-go.svg?branch=master)](https://travis-ci.org/likexian/mailer-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/mailer-go)](https://goreportcard.com/report/github.com/likexian/mailer-go)

mailer-go is a simple Go module for sending smtp email.

## Overview

Help you sending email in the simple and easy way. Sending attachment is supported.

## Installation

    go get github.com/likexian/mailer-go

## Importing

    import (
        "github.com/likexian/mailer-go"
    )

## Documentation

Init a mailer

    func New(server, username, password string, ishtml bool) (m *Message)

Add attachement

    func (m *Message) Attach(fname string) (err error)

Do sending

    func (m *Message) Send() (err error)

## Example

    // Set the smtp info
    // New("smtp server:port", "smtp username", "smtp password", "is html mail")
    m := mailer.New("smtp.likexian.com:25", "i@likexian.com", "8Bd0a7681333214", true)

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

Copyright 2015-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
