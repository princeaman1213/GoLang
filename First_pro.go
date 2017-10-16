package main

import "fmt"

func main() {
	f := "Aman Patel"
	var input int
	var t,u string = "first" , "ffh"     //multiple variables can be declared in one go
	fmt.Println("name is",f)
	fmt.Println(len(f))
	fmt.Println("Hello"[0])    //prints ascii value of that character
	fmt.Println(321325*424521)  //simple multiplication
	fmt.Println(t)
	t+=" second"
	fmt.Println(t)
	fmt.Println(u)

	fmt.Scanf("%d",&input)
	fmt.Println(input*2)
}

