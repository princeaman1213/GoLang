package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	dir, err := os.Open(".")    // . open current directory , .. opens prev dir
	if err != nil {
		return
	}
	defer dir.Close()                 // to close at the end

	fileInfos, err := dir.Readdir(5)// read file names and info the dir , same as reading contents in case of a file
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
	//fmt.Println(fileInfos)
}