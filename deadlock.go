package main

import "fmt"

func main() {
	c :=make(chan int)
	go func(){
		c<-1
	}()
	fmt.Println(<-c)

	c1 :=make(chan int)

	go func(){
		for i:=0;i<10;i++{
			c1<-i
		}
		close(c1)
	}()

	for r:=range c1{
		fmt.Println(r)  //if we print <-c1 directly instead of looping then only 1st value gets printed
	}                   //so we range over c1


}
