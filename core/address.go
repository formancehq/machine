package core

import "encoding/binary"

// Represents an adress in the machine's resources, which include
// constants (literals) and variables passed to the program
type Address uint16

func NewDataAddress(x uint16) Address {
	return Address(x)
}

func NewVarAddress(x uint16) Address {
	return Address((1 << 15) + x)
}

func (a Address) ToBytes() []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(a))
	return bytes
}

func (a Address) ToIdx() int {
	return int(a) & 0x7FFF
}
