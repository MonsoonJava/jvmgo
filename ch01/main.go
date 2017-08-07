package main

import(
	"fmt"
	"jvmgo/ch01/struc"
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
	fmt.Printf("classpath:%s class:%s args:%v\n",
	cmd.CpOption,cmd.Class,cmd.Args)
}