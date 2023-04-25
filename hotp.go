package otp

import (
	"crypto"
	"crypto/hmac"
	"strconv"

	"github.com/chardon55/otp2go/byteorder"
)

const counterLength = 8

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

func (hotp *hotpImpl) Generate(digitCount uint8) uint32 {
	counterHead := make([]byte, counterLength)
	byteorder.ByteOrder().PutUint64(counterHead, hotp.counter)

	encoder := hmac.New(hotp.algorithm.New, hotp.secret)
	encoder.Write(counterHead)

	return truncate(encoder.Sum(nil)) % digitPow[digitCount]
}

func (hotp *hotpImpl) GenerateString(digitCount uint8) string {
	return strconv.FormatUint(uint64(hotp.Generate(digitCount)), 10)
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

func NewHOTP(secret []byte, algorithm crypto.Hash) HOTP {
	return &hotpImpl{
		secret:    secret,
		algorithm: algorithm,
		counter:   0,
	}
}
