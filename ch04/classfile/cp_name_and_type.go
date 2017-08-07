package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex        uint16
	descriptionIndex uint16
}

func (cnati *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cnati.nameIndex = reader.readUint16()
	cnati.descriptionIndex = reader.readUint16()
}
