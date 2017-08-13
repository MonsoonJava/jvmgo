package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLoacls      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributeInfo  []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}


func (ca *CodeAttribute) MaxLocals() uint{
	return uint(ca.maxLoacls)
}


func (ca *CodeAttribute) MaxStack() uint{
	return uint(ca.maxStack)
}


func (ca *CodeAttribute) Code() []byte{
	return ca.code
}

func (ca *CodeAttribute) readInfo(reader *ClassReader) {
	ca.maxStack = reader.readUint16()
	ca.maxLoacls = reader.readUint16()
	codeLen := reader.readUint32()
	ca.code = reader.readBytes(codeLen)
	ca.exceptionTable = readExceptionTable(reader)
	ca.attributeInfo = readAttributes(reader, ca.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionLen := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionLen)
	for i, _ := range exceptionTable {
		exceptionTable[i] = readExceptionEntry(reader)
	}
	return exceptionTable
}

func readExceptionEntry(reader *ClassReader) *ExceptionTableEntry {
	exceptionEntry := ExceptionTableEntry{}
	exceptionEntry.startPc = reader.readUint16()
	exceptionEntry.endPc = reader.readUint16()
	exceptionEntry.handlerPc = reader.readUint16()
	exceptionEntry.catchType = reader.readUint16()
	return &exceptionEntry
}
