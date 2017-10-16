package main

import (
	"fmt"
	//"time"
)

func main() {
	c :=incrementor()
	c1 :=printchan(c)

	for r:=range c1{
		fmt.Println(r)
	}

}

func incrementor() chan int{
	c :=make(chan int)

	go func(){
		for i:=0;i<10;i++{
			c<-i
			//fmt.Println(c)
		}
		close(c)
	}()
	return c
}

func printchan(c <-chan int) chan int{   //channel direction
     out :=make(chan int)
	go func(){
//var t int
		for r :=range c{
			//fmt.Println(r)
			//t = r
			out<-r
		}
		//out<-t
close(out)
	}()
	return out
}

