package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	mionorVersion uint16
	majorVersion  uint16
	constantPool  ConstantPool
	accessFlags   uint16
	thisClass     uint16
	superClass    uint16
	interfaces    []uint16
	fields        []*MemberInfo
	methods       []*MemberInfo
	attributes    []AttributeInfo
}

func Prase(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if v := recover(); v != nil {
			var ok bool
			err, ok = v.(error)
			if ok {
				fmt.Errorf("%v", err)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)

}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.mionorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.mionorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportCalssVersionError!")
}

func (cf *ClassFile) MinorVersion() uint16 { // getter
	return cf.mionorVersion
}

func (cf *ClassFile) MajorVersion() uint16 { // getter
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool { // getter
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 { // getter
	return cf.accessFlags
}

func (cf *ClassFile) Fields() []*MemberInfo { // getter
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo { // getter
	return cf.methods
}
func (cf *ClassFile) ClassName() string { // getter
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string { // getter
	if cf.superClass > 0{
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""  //只有java.lang.Object类的superClass = 0，没有超类

}

func (cf *ClassFile) InterfaceNames() []string { // getter
	interfacesName := make([]string,len(cf.interfaces))
	for i,interIndex := range cf.interfaces{
		interfacesName[i] = cf.constantPool.getClassName(interIndex)
	}
	return interfacesName
}
