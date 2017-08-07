package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) *CompositeEntry{
	rightPath := path[:len(path) - 1]
	entrys := []Entry{}
	walkFu := func(path string,info os.FileInfo,err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != rightPath{
			return filepath.SkipDir
		}
		if strings.Contains(path,".jar") || strings.Contains(path,".JAR"){
			jarEntry := newZipEntry(path)
			entrys = append(entrys,jarEntry)
		}
		return nil
	}
	filepath.Walk(rightPath,walkFu)
	compositeEntry := CompositeEntry(entrys)
	return &compositeEntry

}
