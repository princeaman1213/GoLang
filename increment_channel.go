package main

import (
	"fmt"
	//"time"
)

func main() {
	c :=incrementor("foo")
	c1 :=incrementor("foolish")
	c2 :=printchan(c)
	c3 :=printchan(c1)

	//for r:=range c2{
		fmt.Println(<-c2)
	//}

	//fmt.Println(<-c2)
	for r1:=range c3{
		fmt.Println(r1)
	}

}

func incrementor(s string) chan int{
	c :=make(chan int)

	go func(){
		for i:=0;i<10;i++{
			c<-1
			//fmt.Println(s,i)
			//fmt.Println(c)
		}
		close(c)
	}()
	return c
}

func printchan(c <-chan int) chan int{   //channel direction
	out :=make(chan int)
	go func(){
		var t int
		for r :=range c{
			//fmt.Println(r)
			t += r                      //we take c<-1 therefore we are making a incrementor here
			//out<-r
		}
		out<-t
		close(out)
	}()
	return out
}

