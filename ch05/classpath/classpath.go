package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption string,cpOption string) *Classpath  {
	cp := Classpath{}
	cp.parseJreWithJreOption(jreOption)
	cp.parseClassPathWithCpOption(cpOption)
	return &cp
}

func(cls *Classpath)parseJreWithJreOption(jreOption string){
	jreDir := getJreDir(jreOption)
	// 去jre/lib/* 下加载
	jreLibPath := filepath.Join(jreDir,"lib","*")
	cls.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir,"lib","ext","*")
	cls.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string  {
	//jre is real exist
	if jreOption != "" || existsDir(jreOption){
		return jreOption
	}
	//当前目录是否存在jre
	if existsDir("./jre"){
		return "./jre"
	}
	//寻找java_home下是否存在jre
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome,"/jre")
	}
	panic("can not find jre floder")
}

func existsDir(fileDir string) bool {
	_,err := os.Stat(fileDir)
	if os.IsNotExist(err){
		return false
	}
	return true
}

func(cls *Classpath)parseClassPathWithCpOption(cpOption string){
	if cpOption == "" {
		cpOption = "."
	}
	cls.userClasspath = newEntry(cpOption)
}


//从类路径，扩展类路径，用户路径来搜索class类
func (cls *Classpath) ReadClass(className string) ([]byte,Entry,error) {
	clazzName := className + ".class"
	if data,entry,err := cls.bootClasspath.readClass(clazzName); err == nil{
		return data,entry,err;
	}
	if data,entry,err := cls.extClasspath.readClass(clazzName); err == nil{
		return data,entry,err;
	}
	return cls.userClasspath.readClass(clazzName)
}

func (cls *Classpath) String() string{
	return cls.userClasspath.String()
}
