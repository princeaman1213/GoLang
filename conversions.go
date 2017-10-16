package main

import (
	"fmt"
	//"strconv"
)

func main() {
	//a :=strconv.Itoa(44)
	//fmt.Printf("%T",a)

	//fmt.Println(strconv.Itoa(22))


	var name interface{}=33
	var name1 interface{}="Aman"

	fmt.Println(name1.(string))
	fmt.Println(name.(int))       //assertion
	fmt.Printf("%T %T \n",name1,name)

	fmt.Println(name.(int)+12)



}
