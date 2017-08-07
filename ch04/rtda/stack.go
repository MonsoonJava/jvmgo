package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: 1024,
		size:    0,
		_top:    nil,
	}
}

func (stack *Stack) PushFrame(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("error:java.lang.StackOverFlowError")
	}
	if stack._top == nil {
		stack._top = frame
	} else {
		frame.lower = stack._top
		stack._top = frame
	}
	stack.size++

}

func (stack *Stack) PopFrame() *Frame {
	if stack._top == nil {
		panic("the stack is empty!")
	}
	topFrame := stack._top
	stack._top = stack._top.lower
	topFrame.lower = nil
	stack.size--
	return topFrame
}

func (stack *Stack) TopFrame() *Frame {
	if stack._top == nil {
		panic("the stack is empty!")
	}
	return stack._top
}
