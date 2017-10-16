package main

import "fmt"

func main() {

	//var a int =0
	//fmt.Scanf("%d",&a)
	for i:=0 ; i<11;i++{
		fmt.Println(fib(i))
	}

	//fib(13)

}

func fib(a int) int{
	//var t int
	if a==0{
		//fmt.Println(a)
		return 0

	} else if a==1{
		//fmt.Println(a)
		return 1

	} else{
	    return fib(a-1)+fib(a-2)

	}



}
