package classfile

/*
	SourceFile_attribute{
		u2 attribute_name_index;  /has been read before
		u4 attribute_length;  / has been read before
		u2 sourcefile_index;  / need to get
	}

*/

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (sfa *SourceFileAttribute) readInfo(reader *ClassReader) {
	sfa.sourceFileIndex = reader.readUint16()
}

func (sfa *SourceFileAttribute) SourceFileName() string {
	return sfa.cp.getUtf8(sfa.sourceFileIndex)
}
