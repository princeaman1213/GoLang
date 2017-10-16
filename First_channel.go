package main

import (
	"fmt"
	"time"
)

func pinger(c chan int){ //<- makes it unidirectional (only send , no receive)
	//ar :=[3]string{"a","s","f"}
	for i :=0;i<10 ;i++{
         c<-i
	}
}

//func ponger(c chan string){
//	for i :=0; ;i++{
//		c <- "pong"
//		time.Sleep(time.Second*5)
//	}
//}

func printer(c chan int){
	for{
        msg := <- c
		//msg1 := <-d
		fmt.Println("The info is :     ",msg)
		time.Sleep(time.Second)
	}
}

func main() {
	c :=make(chan int)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
