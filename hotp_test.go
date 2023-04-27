package otp_test

import (
	"crypto"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"encoding/hex"
	"testing"

	otp "github.com/chardon55/otp2go"
)

func TestHTOP1(t *testing.T) {
	input := "3132333435363738393031323334353637383930"
	counter := 1
	expected := "94287082"

	hexBytes, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	hotp := otp.NewHOTP(hexBytes, crypto.SHA1)
	hotp.SetCounter(uint64(counter))
	result := hotp.GenerateString8()
	if expected != result {
		t.Errorf("Not matched. Expected: %s; Actual: %s", expected, result)
	}
}

func TestHTOP2(t *testing.T) {
	input := "3132333435363738393031323334353637383930313233343536373839303132"
	counter := 0x23523ED
	expected := "67062674"

	hexBytes, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	hotp := otp.NewHOTP(hexBytes, crypto.SHA256)
	hotp.SetCounter(uint64(counter))
	result := hotp.GenerateString8()
	if expected != result {
		t.Errorf("Not matched. Expected: %s; Actual: %s", expected, result)
	}
}

func TestHTOP3(t *testing.T) {
	input := "31323334353637383930313233343536373839303132333435363738393031323334353637383930313233343536373839303132333435363738393031323334"
	counter := 0x23523EC
	expected := "25091201"

	hexBytes, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	hotp := otp.NewHOTP(hexBytes, crypto.SHA512)
	hotp.SetCounter(uint64(counter))
	result := hotp.GenerateString8()
	if expected != result {
		t.Errorf("Not matched. Expected: %s; Actual: %s", expected, result)
	}
}
