package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

//类路径
func newDirectEntry(path string) *DirEntry {
	absPath,error := filepath.Abs(path)
	if error != nil{
		panic(error)
	}
	return &DirEntry{absPath}
}


func (de *DirEntry)readClass(className string) ([]byte,Entry,error)  {
	fileName := filepath.Join(de.absDir,className)
	fileDatas,error := ioutil.ReadFile(fileName)
	return fileDatas,de,error
}

func (de *DirEntry) String() string{
	return de.absDir
}