# GoKit - xcron

Cron kits for Golang development.

## Features

- Thread safe and Easy to use
- Fractional precision to Seconds
- Compatible with Standard cron expression
- Support Nonstandard macros definitions
- Extense Support @every N duration
- Dynamic add、update、remove、empty cron job

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
Month        | Yes        | 1–12 or JAN–DEC | * / , -
Day of week  | Yes        | 0–6 or SUN–SAT  | * / , -
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

### Cron service

```go
// start a cron service
service := xcron.New()

// add a job to service, specify the rule and loop func
id, err := service.Add("@every second", func(){fmt.Println("add a echo")})

// update exists job by job id
err = service.Set(id, "@every second", func(){fmt.Println("set a echo")})

// delete exists job from service, job will stop
service.Del(id)

// clear all jobs, jobs will stop
service.Empty()

// wait for all job exit
service.Wait()
```

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

## License

Copyright 2012-2021 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
