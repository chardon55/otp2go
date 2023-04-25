package byteorder

import (
	"encoding/binary"
	"unsafe"
)

func ByteOrder() binary.ByteOrder {
	buffer := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buffer[0])) = uint16(0x00FF)

	if buffer == [2]byte{0xFF, 0x00} {
		return binary.LittleEndian
	}

	return binary.BigEndian
}
