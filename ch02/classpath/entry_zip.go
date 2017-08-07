package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string)*ZipEntry {
	absPath,error := filepath.Abs(path)
	if error != nil{
		panic(error)
	}
	return &ZipEntry{absPath}
}

func (zipEntry *ZipEntry) String() string  {
	return zipEntry.absDir
}

func (zipEntry *ZipEntry) readClass(className string) ([]byte,Entry,error){
	readerCloser,err := zip.OpenReader(zipEntry.absDir)
	if err != nil {
		return nil,nil,err
	}
	defer  readerCloser.Close()
	for _,file := range readerCloser.File{
		//加载返回
		if file.Name == className {
			rs,filErr := file.Open();
			if filErr != nil{
				return nil,nil,filErr
			}
			defer rs.Close()
			data,redErr := ioutil.ReadAll(rs)
			if redErr != nil {
				return nil,nil,filErr
			}
			return data,zipEntry,nil
		}
	}
	return nil,nil,errors.New("class not found" + className)
}