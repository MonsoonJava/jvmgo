package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (th *Thread) Pc() int {
	return th.pc
}

func (th *Thread) SetPc(pc int) {
	th.pc = pc
}

func (th *Thread) PushFrame(frame *Frame) {
	th.stack.PushFrame(frame)
}

func (th *Thread) PopFrame() *Frame {
	return th.stack.PopFrame()
}

func (th *Thread) CurrentFrame() *Frame {
	return th.stack.TopFrame()
}
