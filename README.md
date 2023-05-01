# otp2go

![Version](https://img.shields.io/github/v/release/chardon55/otp2go?include_prereleases)
![License](https://img.shields.io/github/license/chardon55/otp2go)

Simple OTP library for Go

License: BSD-3-Clause

## Usage

### At a glance

```go
// ...
import (
    "fmt"
    "crypto"
    _ "crypto/sha1"

    otp "github.com/chardon55/otp2go"
)
// ...
totp, err := otp.NewTOTPBase32("<Your Base32 secret>", crypto.SHA1)
// check error

password, remainTime := totp.GenerateString6()
fmt.Printf("Password: %s, remaining seconds: %d", password, remainTime)

```
