package otp

import "time"

const ensBit = 8

const size = 4

func truncate(hash []byte) uint32 {
	offset := hash[len(hash)-1] & 0xF
	var code uint32

	for i := 0; i < size; i++ {
		code <<= ensBit
		code |= uint32(hash[offset+1])
	}

	return code << 1 >> 1
}

func calcTimeBasedCounter(interval uint32) (uint64, int) {
	curTime := time.Now().Unix()
	return uint64(curTime) / uint64(interval), int(interval - uint32(uint64(curTime)%uint64(interval)))
}
