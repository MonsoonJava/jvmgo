package main

import(
	"fmt"
	"jvmgo/ch02/struc"
	"jvmgo/ch02/classpath"
	"strings"
)

func main()  {
	cmd := struc.ParseCmd()
	//fmt.Println(cmd)
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	}else if cmd.HelpFlag || cmd.Class == "" {
		struc.PrintUsage()
	}else {
		startJVM(cmd)
	}
}

func startJVM(cmd *struc.Cmd)  {
	fmt.Println("pathSeprate ( " + classpath.PathListSepartor + " )")
	cl := classpath.Parse(cmd.XjreOption,cmd.CpOption)
	fmt.Println("classpath:%v class:%v args:%v\n",cmd.CpOption,cmd.Class,cmd.Args)
	className := strings.Replace(cmd.Class,".","/",-1)
	data,_,err := cl.ReadClass(className)
	if err != nil{
		fmt.Printf("can not find class or load class" + cmd.Class)
		return
	}
	fmt.Printf("class data:%v\n",data)
}