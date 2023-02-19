# GoKit - xaes

AES kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xaes"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xaes)

## Example

```go
// Encrypt
ciphertext, err := xaes.CBCEncrypt(plaintext, key, iv)
if err != nil {
    panic(err)
}
fmt.Printf("Encrypted: %x", ciphertext)

// Decrypt
plaintext, err := CBCDecrypt(ciphertext, key, iv)
if err != nil {
    panic(err)
}
fmt.Printf("Decrypted: %s", plaintext)
```

## License

Copyright 2012-2023 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
