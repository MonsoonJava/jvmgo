package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) *CompositeEntry {
	lists := strings.Split(pathList,PathListSepartor)
	var entrys CompositeEntry = make([]Entry,0,10)
	for _,path := range lists{
		entry := newEntry(path)
		entrys = append(entrys,entry)
	}
	return &entrys
}

func(compositeEntry *CompositeEntry) readClass(className string)([]byte,Entry,error){
	for _,entry := range *compositeEntry{
		data,from,err :=entry.readClass(className)
		if err == nil {
			return data,from,err
		}
	}
	return nil,nil,errors.New("class not found: "+className)
}

func(compositeEntry *CompositeEntry) String() string {
	strs := make([]string,len(*compositeEntry))
	for i,ele := range *compositeEntry{
		strs[i] = ele.String()
	}
	return strings.Join(strs,PathListSepartor)
}