package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		//每有一个常量类型为Constant_Long_info或Constant_Double_info类型，常量数量减1
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			{
				i++
			}
		}
	}
	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	cnat := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(cnat.nameIndex)
	description := cp.getUtf8(cnat.descriptionIndex)
	return name,description
}

func (cp ConstantPool) getClassName(index uint16) string {
	cn := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(cn.classIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.val
}
