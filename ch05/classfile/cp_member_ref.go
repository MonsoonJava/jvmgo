package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (cmri *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	cmri.classIndex = reader.readUint16()
	cmri.nameAndTypeIndex = reader.readUint16()
}

func (cmri *ConstantMemberrefInfo) ClassName() string {
	return cmri.cp.getClassName(cmri.classIndex)
}

func (cmri *ConstantMemberrefInfo) NameAndDescription() (string, string) {
	return cmri.cp.getNameAndType(cmri.nameAndTypeIndex)
}
