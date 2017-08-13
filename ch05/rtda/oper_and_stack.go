package rtda

import "math"

type OperAndStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperAndStack {
	if maxStack > 0 {
		return &OperAndStack{
			size:  0,
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (oas *OperAndStack) PopSlot() Slot {
	oas.size--
	slot := oas.slots[oas.size]
	return slot
}

func (oas *OperAndStack) PushSlot(slot Slot) {
	oas.slots[oas.size] = slot
	oas.size++
}


func (oas *OperAndStack) PushInt(val int32) {
	oas.slots[oas.size].num = val
	oas.size++
}

func (oas *OperAndStack) PopInt() int32 {
	oas.size--
	return oas.slots[oas.size].num
}

func (oas *OperAndStack) PushFloat(val float32) {
	fToI := math.Float32bits(val)
	oas.slots[oas.size].num = int32(fToI)
	oas.size++
}

func (oas *OperAndStack) PopFloat() float32 {
	oas.size--
	uint32Val := uint32(oas.slots[oas.size].num)
	return math.Float32frombits(uint32Val)
}

func (oas *OperAndStack) PushLong(val int64) {
	oas.slots[oas.size].num = int32(val)
	oas.slots[oas.size+1].num = int32(val >> 32)
	oas.size += 2
}

func (oas *OperAndStack) PopLong() int64 {
	oas.size -= 2
	low := uint32(oas.slots[oas.size].num)
	high := uint32(oas.slots[oas.size+1].num)
	longVal := int64(high) << 32 | int64(low)
	return longVal
}

func (oas *OperAndStack) PushDouble(val float64) {
	fTol := int64(math.Float64bits(val))
	oas.PushLong(fTol)
}

func (oas *OperAndStack) PopDouble() float64 {
	bits := uint64(oas.PopLong())
	return math.Float64frombits(bits)
}

func (oas *OperAndStack) PushRef(val *Object) {
	oas.slots[oas.size].ref = val
	oas.size++
}

func (oas *OperAndStack) PopRef() *Object {
	if oas.size == 0 {
		return nil
	}
	oas.size--
	retRef := oas.slots[oas.size].ref
	oas.slots[oas.size].ref = nil

	return retRef
}
