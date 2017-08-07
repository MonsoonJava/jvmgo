package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/

type ConstantMethodHandleInvokeDynamicInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (cmhi *ConstantMethodHandleInvokeDynamicInfo) readInfo(reader *ClassReader) {
	cmhi.referenceKind = reader.readUint8()
	cmhi.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/

type ConstantMethodTypeInfo struct {
	descriptionIndex uint16
}

func (cmti *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	cmti.descriptionIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (cmidi *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	cmidi.bootstrapMethodAttrIndex = reader.readUint16()
	cmidi.nameAndTypeIndex = reader.readUint16()
}
