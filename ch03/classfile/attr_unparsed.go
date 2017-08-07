package classfile

type UnparseAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (ua *UnparseAttribute) readInfo(reader *ClassReader) {
	ua.info = reader.readBytes(ua.length)
}
