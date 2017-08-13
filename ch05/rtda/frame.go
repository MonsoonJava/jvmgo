package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVals
	operAndStack *OperAndStack
	thread 		 *Thread
	nextPC 		 int    //the next instruction after call
}

func NewFrame(maxLocals uint, maxStack uint,thread *Thread) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    NewLocalVals(maxLocals),
		operAndStack: newOperandStack(maxStack),
	}
}

func (frame *Frame) LocalVars() LocalVals {
	return frame.localVars
}

func (frame *Frame) Thread() *Thread {
	return frame.thread
}

func (frame *Frame) NextPC() int {
	return frame.nextPC
}

func (frame *Frame) OperandStack() *OperAndStack {
	return frame.operAndStack
}

func (frame *Frame) SetNextPc(nextPc int) {
	frame.nextPC = nextPc
}