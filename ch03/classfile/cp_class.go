package classfile

type ConstantClassInfo struct {
	cp         ConstantPool
	classIndex uint16
}

func (cci *ConstantClassInfo) readInfo(reader *ClassReader) {
	cci.classIndex = reader.readUint16()
}

func (cci *ConstantClassInfo) Name() string {
	return cci.cp.getUtf8(cci.classIndex)
}
