package base

import "jvmgo/ch05/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().Pc()
	nextPC := pc + offset
	frame.SetNextPc(nextPC)
}
