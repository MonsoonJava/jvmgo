package main

import (
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/rtda"
	"fmt"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	codeByte := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)
	defer catchError(frame)
	loop(thread,codeByte)
}

func catchError(frame *rtda.Frame){
	if r := recover(); r != nil {
		fmt.Println("localVals %v",frame.LocalVars())
		fmt.Println("maxStack %v",frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread,codeByte []byte)  {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		fmt.Println(pc)
		thread.SetPc(pc)
		//decode
		reader.Reset(codeByte,pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPc(reader.PC())
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
