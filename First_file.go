package main

import(
	"fmt"
	"io/ioutil"
    "errors"
)


func main() {
	var i int
	err :=errors.New("dfdsf fds")
	a,err := ioutil.ReadFile("words.txt")     //This program needs to be understood and learned
	if err !=nil{
		panic("panic panic")
	}

	fmt.Printf("%T\n",string(a))
	for i=0;i<10;i++{
		if string(string(a)[i]) == "\n"{
			fmt.Println("\n")
		}
		fmt.Printf(string(string(a)[i]))

	}


}
