# logger.go

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/logger-go?status.svg)](https://godoc.org/github.com/likexian/logger-go)
[![Build Status](https://travis-ci.org/likexian/logger-go.svg?branch=master)](https://travis-ci.org/likexian/logger-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/logger-go)](https://goreportcard.com/report/github.com/likexian/logger-go)

logger-go is a simple Go module for doing logging.

## Overview

It will do logging in the simple and easy way.

## Installation

    go get github.com/likexian/logger-go

## Importing

    import (
        "github.com/likexian/logger-go"
    )

## Documentation

Init a logger

    func New(w io.Writer, level Level) *Logger

Init a file logger

    func File(fname string, level Level) (*Logger, error)

Setting log level

    func (l *Logger) SetLevelString(level string) error

Do a DEBUG logging

    func (l *Logger) Debug(msg string, args ...interface{}) error

Do a INFO logging

    func (l *Logger) Info(msg string, args ...interface{}) error

Do a NOTICE logging

    func (l *Logger) Notice(msg string, args ...interface{}) error

Do a WARNING logging

    func (l *Logger) Warning(msg string, args ...interface{}) error

Do a ERROR logging

    func (l *Logger) Error(msg string, args ...interface{}) error

Do a CRITICAL logging

    func (l *Logger) Critical(msg string, args ...interface{}) error

## Example

### Do logging to stderr

    log := logger.New(os.Stderr, logger.INFO)
    log.Info("This is Info")
    log.SetLevel(logger.DEBUG)
    log.Debug("This is Debug")

### Do logging to a file

    flog, err := logger.File("test.log", logger.DEBUG)
    if err != nil {
        panic(err)
    }
    flog.Debug("This is Debug")
    flog.Info("This is Info")

## LICENSE

Copyright 2015-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
