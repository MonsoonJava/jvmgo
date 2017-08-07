package main

import (
	"fmt"
	"jvmgo/ch04/struc"
	"jvmgo/ch04/rtda"
)

func main() {
	cmd := struc.ParseCmd()
	//fmt.Println(cmd)
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		struc.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *struc.Cmd) {
	frame := rtda.NewFrame(100,100)
	testLocalVals(frame.LocalVars())
	testOperandStack(frame.OperAndStack())
}

func testLocalVals(point rtda.LocalVals){
	point.SetInt(0, 100)
	point.SetInt(1, -100)
	point.SetLong(2, 2997924580)
	point.SetLong(4, -2997924580)
	point.SetFloat(6, 3.1415926)
	point.SetDouble(7, 2.71828182845)
	point.SetRef(9, nil)
	println(point.GetInt(0))
	println(point.GetInt(1))
	println(point.GetLong(2))
	println(point.GetLong(4))
	println(point.GetFloat(6))
	println(point.GetDouble(7))
	println(point.GetRef(9))
}

func testOperandStack(ops *rtda.OperAndStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}


