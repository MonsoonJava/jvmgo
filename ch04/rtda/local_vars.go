package rtda

import (
	"math"
)

//局部变量表
type LocalVals []Slot

func NewLocalVals(maxLocals uint) LocalVals {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (lv *LocalVals) SetInt(index uint, val int32) {
	(*lv)[index].num = val
}

func (lv *LocalVals) GetInt(index uint) int32 {
	return (*lv)[index].num
}

func (lv *LocalVals) SetFloat(index uint, val float32) {
	int32Val := math.Float32bits(val)
	(*lv)[index].num = int32(int32Val)
}

func (lv *LocalVals) GetFloat(index uint) float32 {
	int32Val := uint32((*lv)[index].num)
	return math.Float32frombits(int32Val)
}

func (lv *LocalVals) SetLong(index uint, val int64) {
	(*lv)[index].num = int32(val)
	(*lv)[index+1].num = int32(val >> 32)
}

func (lv *LocalVals) GetLong(index uint) int64 {
	low := uint32((*lv)[index].num)
	high := uint32((*lv)[index+1].num)
	return int64(high) << 32 | int64(low)
}

func (lv *LocalVals) SetDouble(index uint, val float64) {
	intVal := math.Float64bits(val)
	(*lv).SetLong(index, int64(intVal))
}

func (lv *LocalVals) GetDouble(index uint) float64 {
	int64Val := (*lv).GetLong(index)
	return math.Float64frombits(uint64(int64Val))
}

func (lv *LocalVals) GetRef(index uint) *Object {
	return (*lv)[index].ref
}

func (lv *LocalVals) SetRef(index uint, ref *Object) {
	(*lv)[index].ref = ref
}
