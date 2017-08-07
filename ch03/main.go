package main

import (
	"fmt"
	"jvmgo/ch03/classfile"
	"jvmgo/ch03/classpath"
	"jvmgo/ch03/struc"
	"strings"
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
	cl := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cl)
	printClassInfo(cf)
}

func loadClass(className string, cl *classpath.Classpath) *classfile.ClassFile {
	data, _, err := cl.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, erss := classfile.Prase(data)
	if erss != nil {
		panic(erss)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
