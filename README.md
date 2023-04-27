# otp2go

Free and open-source OTP library for Go with simplicity ~~and stupidity~~

License: BSD-3-Clause

## Usage

### At a glance

```go
// ...
import (
    "fmt"
    "crypto"
    _ "crypto/sha1"

    "github.com/chardon55/otp2go"
)
// ...
totp, err := otp.NewTOTPBase32("<Your Base32 secret>", crypto.SHA1)
// check error

password, remainTime := totp.GenerateString6()
fmt.Printf("Password: %s, remaining seconds: %d", password, remainTime)

```
