# GoKit - xhttp

HTTP kits for Golang development.

## Features

- Light weight and Easy to use
- Cookies and Proxy are support
- Easy use with friendly JSON api
- Upload and Download file support
- Debug and Trace info are open
- Retry request is possible
- Cache request by method

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xhttp"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xhttp)

## Example

### The Most easy way

```go
rsp, err := xhttp.Get(context.Background(), "https://www.likexian.com/")
if err != nil {
    panic(err)
}

defer rsp.Close()
text, err := rsp.String()
if err == nil {
    fmt.Println("http status code:", rsp.StatusCode)
    fmt.Println("http response body:", text)
}
```

### Do a Post with form and files

```go
// xhttp.FormParam is form, xhttp.FormFile is file
rsp, err := xhttp.Post(context.Background(), "https://www.likexian.com/",
    xhttp.FormParam{"name": "likexian", "age": 18}, xhttp.FormFile{"file": "README.md"})
if err != nil {
    panic(err)
}

defer rsp.Close()
json, err := rsp.JSON()
if err == nil {
    // http response {"status": {"code": 1, "message": "ok"}}
    code, _ := json.Get("status.code").Int()
    fmt.Println("json status code:", code)
}
```

### Use as Interactive mode

```go
req := xhttp.New()

// set ua and referer
req.SetUA("the new ua")
req.SetReferer("http://the-referer-url.com")

// set tcp connect timeout and client total timeout
req.SetConnectTimeout(3)
req.SetClientTimeout(30)

// not follow 302 and use cookies
req.FollowRedirect(false)
req.EnableCookie(true)

// will send get to https://www.likexian.com/?v=1.0.0
rsp, err := req.Get(context.Background(), "https://www.likexian.com/", xhttp.QueryParam{"v", "1.0.0"})
if err != nil {
    panic(err)
}

// save file as index.html
defer rsp.Close()
_, err := rsp.File("index.html")
if err == nil {
    fmt.Println("Url download as index.html")
}

// use the request param as above
rsp, err := req.Get(context.Background(), "https://www.likexian.com/page/")
if err != nil {
    panic(err)
}

defer rsp.Close()
...
```

### xhttp.Request not thread-safe

This version of xhttp.Request is not thread-safe, please New every thread when doing concurrent

```go
for i := 0; i < 100; i++ {
    go func() {
        // always New one
        req := New()
        rsp, err := req.Do(context.Background(), "GET", LOCALURL)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer rsp.Close()
        str, err := rsp.String()
        if err == nil {
            fmt.Println(str)
        }
    }()
}
```

## License

Copyright 2012-2021 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
