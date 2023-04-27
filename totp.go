package otp

import (
	"crypto"
	"encoding/base32"
)

// TOTP interface
type TOTP interface {

	// Get interval
	Interval() uint32

	// Set interval
	SetInterval(uint32)

	// Generate a TOTP password
	Generate(digitCount uint8) (uint32, int)

	// Generate a TOTP password as a string
	GenerateString(digitCount uint8) (string, int)

	// Generate a 4-digit HOTP password
	Generate4() (uint32, int)

	// Generate a 4-digit HOTP password as a string
	GenerateString4() (string, int)

	// Generate a 6-digit HOTP password
	Generate6() (uint32, int)

	// Generate a 6-digit HOTP password as a string
	GenerateString6() (string, int)

	// Generate an 8-digit HOTP password
	Generate8() (uint32, int)

	// Generate an 8-digit HOTP password as a string
	GenerateString8() (string, int)
}

type totpAdapter struct {
	hotp     HOTP
	counter  uint64
	delta    int
	interval uint32
}

func (totp *totpAdapter) updateTime() {
	totp.counter, totp.delta = calcTimeBasedCounter(totp.interval)
	totp.hotp.SetCounter(totp.counter)
}

func (totp *totpAdapter) Interval() uint32 {
	return totp.interval
}

func (totp *totpAdapter) SetInterval(interval uint32) {
	totp.interval = interval
}

func (totp *totpAdapter) Generate(digitCount uint8) (uint32, int) {
	totp.updateTime()
	return totp.hotp.Generate(digitCount), totp.delta
}

func (totp *totpAdapter) GenerateString(digitCount uint8) (string, int) {
	totp.updateTime()
	return totp.hotp.GenerateString(digitCount), totp.delta
}

func (totp *totpAdapter) Generate4() (uint32, int) {
	return totp.Generate(4)
}

func (totp *totpAdapter) GenerateString4() (string, int) {
	return totp.GenerateString(4)
}

func (totp *totpAdapter) Generate6() (uint32, int) {
	return totp.Generate(6)
}

func (totp *totpAdapter) GenerateString6() (string, int) {
	return totp.GenerateString(6)
}

func (totp *totpAdapter) Generate8() (uint32, int) {
	return totp.Generate(8)
}

func (totp *totpAdapter) GenerateString8() (string, int) {
	return totp.GenerateString(8)
}

func NewTOTPBase32(base32Secret string, algorithm crypto.Hash) (TOTP, error) {
	secret, err := base32.StdEncoding.DecodeString(base32Secret)
	if err != nil {
		return nil, err
	}

	return NewTOTP(secret, algorithm), nil
}

func NewTOTP(secret []byte, algorithm crypto.Hash) TOTP {
	return &totpAdapter{
		hotp:     NewHOTP(secret, algorithm),
		interval: 30,
		counter:  0,
		delta:    0,
	}
}
