package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberEntry
}

type LineNumberEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (lnta *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	lnta.lineNumberTable = make([]*LineNumberEntry, lineNumberTableLength)
	for i, _ := range lnta.lineNumberTable {
		lnta.lineNumberTable[i] = &LineNumberEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}

}
