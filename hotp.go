package otp

import (
	"crypto"
	"crypto/hmac"
	"encoding/base32"
	"fmt"
)

var digitPow = []uint32{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}

// HOTP interface
type HOTP interface {

	// Set counter
	SetCounter(uint64)

	// Generate an HOTP password
	Generate(digitCount uint8) uint32

	// Generate an HOTP password as a string
	GenerateString(digitCount uint8) string

	// Generate a 4-digit HOTP password
	Generate4() uint32

	// Generate a 4-digit HOTP password as a string
	GenerateString4() string

	// Generate a 6-digit HOTP password
	Generate6() uint32

	// Generate a 6-digit HOTP password as a string
	GenerateString6() string

	// Generate an 8-digit HOTP password
	Generate8() uint32

	// Generate an 8-digit HOTP password as a string
	GenerateString8() string
}

// HOTP implementation
type hotpImpl struct {
	secret    []byte
	algorithm crypto.Hash
	counter   uint64
}

func (hotp *hotpImpl) SetCounter(counter uint64) {
	hotp.counter = counter
}

func (hotp *hotpImpl) Generate(digitCount uint8) (result uint32) {
	counterHead := convCounter(hotp.counter)

	encoder := hmac.New(hotp.algorithm.New, hotp.secret)
	encoder.Write(counterHead)

	result = truncate(encoder.Sum(nil))

	if digitCount < 10 {
		result %= digitPow[digitCount]
	}

	return
}

func (hotp *hotpImpl) GenerateString(digitCount uint8) string {
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", digitCount), hotp.Generate(digitCount))
}

func (hotp *hotpImpl) Generate4() uint32 {
	return hotp.Generate(4)
}

func (hotp *hotpImpl) GenerateString4() string {
	return hotp.GenerateString(4)
}

func (hotp *hotpImpl) Generate6() uint32 {
	return hotp.Generate(6)
}

func (hotp *hotpImpl) GenerateString6() string {
	return hotp.GenerateString(6)
}

func (hotp *hotpImpl) Generate8() uint32 {
	return hotp.Generate(8)
}

func (hotp *hotpImpl) GenerateString8() string {
	return hotp.GenerateString(8)
}

func NewHOTPBase32(base32Secret string, algorithm crypto.Hash) (HOTP, error) {
	secret, err := base32.StdEncoding.DecodeString(base32Secret)
	if err != nil {
		return nil, err
	}

	return NewHOTP(secret, algorithm), nil
}

func NewHOTP(secret []byte, algorithm crypto.Hash) HOTP {
	return &hotpImpl{
		secret:    secret,
		algorithm: algorithm,
		counter:   0,
	}
}
