package classfile


type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberNum := reader.readUint16()
	members := make([]*MemberInfo, memberNum)
	for i, _ := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (mf *MemberInfo) CodeAttribute() *CodeAttribute {
	for _,attr := range mf.attributes {
		switch attr.(type) {
		case *CodeAttribute :
			return attr.(*CodeAttribute)
		}
	}
	return nil
}

func (mf *MemberInfo) AccessFlags() uint16 {
	return mf.accessFlags
}

func (mf *MemberInfo) Name() string {
	return mf.cp.getUtf8(mf.nameIndex)
}

func (mf *MemberInfo) Description() string {
	return mf.cp.getUtf8(mf.descriptorIndex)
}
