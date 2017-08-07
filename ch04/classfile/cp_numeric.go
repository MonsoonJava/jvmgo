package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct {
	val float64
}

func (cii *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	cii.val = int32(reader.readUint32())
}

func (cfi *ConstantFloatInfo) readInfo(reader *ClassReader) {
	cfi.val = math.Float32frombits(reader.readUint32())
}

func (cli *ConstantLongInfo) readInfo(reader *ClassReader) {
	cli.val = int64(reader.readUint64())
}

func (cdi *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	cdi.val = math.Float64frombits(reader.readUint64())
}
