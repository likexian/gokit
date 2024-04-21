# GoKit - xcache

Cache kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xcache"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xcache)

## Example

### Use memory cache

```go
// init memory cache
c := xcache.New(xcache.MemoryCache)

// set gc param, gc every 60s, once clean max 100
c.SetGC(60, 100)

// set key value cache with no expire
c.Set("key", "value", 0)

// set key value cache with ttl, expire after 30s
c.Set("key", "value", 30)

// check key exists
c.Has("key")

// get value
c.Get("key")

// remove key
c.Del("key")

// get multiple once
c.MGet("k1", "k2", "k3")

// do not forget stop the service
c.Close()
```

## License

Copyright 2012-2024 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
