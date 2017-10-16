package main

import (
	"os"
	"fmt"
)

func main() {
	file,err :=os.Open("name.txt")
	if err!= nil{
		return
	}
	defer file.Close()

	stat,err :=file.Stat()
	if err !=nil{
		return
	}

	bs :=make([]byte,stat.Size())
	data,err :=file.Read(bs)
	if err!=nil{
		return
	}
	fmt.Println(data)
	fmt.Println(string(bs))
}
