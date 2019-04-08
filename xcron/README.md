# GoKit - xcron

Cron kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xcron"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xcron)

## Field of rule

```
Field name   | Mandatory  | Allowed values  | Allowed special characters
------------ | ---------- | --------------- | --------------------------
Seconds      | No         | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , -
Month        | Yes        | 1-12            | * / , -
Day of week  | Yes        | 0-6             | * / , -
```

## Predefined rule

```
Entry                  | Description                                | Equivalent To
---------------------- | ------------------------------------------ | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
```

## Example

### Parse cron rule

```go
// standard cron rule, every hour
rule, err := xcron.Parse("0 * * * *")

// standard cron rule, every half hour
rule, err := xcron.Parse("0,30 * * * *")

// standard cron rule, every half hour
rule, err := xcron.Parse("0/30 * * * *")

// every second, six fields
rule, err := xcron.Parse("* * * * * *")

// every hour
rule, err := xcron.Parse("@every hour")

// every 6 hour
rule, err := xcron.Parse("@every 6 hour")
```

## LICENSE

Copyright 2012-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
