package main

import "fmt"
import "jvmgo/ch05/struc"
import "strings"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/classpath"

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
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	fmt.Println(className)
	cf := loadClass(className, cp)
	fmt.Println(cf)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.Class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Prase(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		fmt.Println(m.Name())
		if m.Name() == "main" && m.Description() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}