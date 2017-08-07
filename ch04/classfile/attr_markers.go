package classfile

/*
Deprecated_attribute{
	u2 attribute_name_index
	u4 attribute_length === 0
}


Synthetic_attribute{
	u2 attribute_name_index
	u4 attribute_length === 0
}

*/
type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}
type MarkerAttribute struct {
}

//becauese of Deprecated and Sybthetic is just marker ,readInfo()
//method dont need to read anything
func (ma *MarkerAttribute) readInfo(reader *ClassReader) {

}
