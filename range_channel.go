package main

import "fmt"

func main() {
	c :=make(chan int)

	go func(){
		for i:=0;i<10;i++{
			c<-i
		}
		close(c)
	}()

	for  v:=range c{
		fmt.Println(v)
	}
	//for {
		//print(c)
	//	fmt.Println(<-c)
	//}


	//var input string
	//fmt.Scanln(&input)

}
/*
func print(c chan int){
	fmt.Println(<-c)
}
*/