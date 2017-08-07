package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVals
	operAndStack *OperAndStack
}

func NewFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    NewLocalVals(maxLocals),
		operAndStack: newOperandStack(maxStack),
	}
}

func (frame *Frame) LocalVars() LocalVals {
	return frame.localVars
}

func (frame *Frame) OperAndStack() *OperAndStack {
	return frame.operAndStack
}
