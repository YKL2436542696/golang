package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage: xxx file")
		return
	}

	fileName := list[1]

	// type FileInfo interface {
	//	Name() string       // base name of the file
	//	Size() int64        // length in bytes for regular files; system-dependent for others
	//	Mode() FileMode     // file mode bits
	//	ModTime() time.Time // modification time
	//	IsDir() bool        // abbreviation for Mode().IsDir()
	//	Sys() any           // underlying data source (can return nil)
	//}
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	fmt.Println("fileName = ", info.Name())
	fmt.Println("fileSize = ", info.Size())

}
