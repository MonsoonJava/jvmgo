package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func (cr *ClassReader) readUint8() uint8 { //u1
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}
func (cr *ClassReader) readUint16() uint16 { //u2
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

func (cr *ClassReader) readUint32() uint32 { //u4
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

func (cr *ClassReader) readUint16s() []uint16 {
	count := cr.readUint16()
	table := make([]uint16, count)
	for i, _ := range table {
		table[i] = cr.readUint16()
	}
	return table

}

func (cr *ClassReader) readBytes(length uint32) []byte {
	val := cr.data[:length]
	cr.data = cr.data[length:]
	return val
}
