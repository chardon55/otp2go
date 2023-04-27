package otp

import "time"

const ensBit = 8

const size = 4

func convCounter(counter uint64) (out []byte) {
	out = make([]byte, 8)
	for i := 7; i >= 0; i-- {
		out[i] = byte(counter & 0xff)
		counter >>= 8
	}

	return
}

func truncate(hash []byte) (code uint32) {
	offset := int(hash[len(hash)-1] & 0xF)

	for i := 0; i < size; i++ {
		code <<= ensBit
		code |= uint32(hash[offset+i])
	}

	code <<= 1
	code >>= 1

	return
}

func calcTimeBasedCounter(interval uint32) (uint64, int) {
	curTime := time.Now().Unix()
	return uint64(curTime) / uint64(interval), int(interval - uint32(uint64(curTime)%uint64(interval)))
}
