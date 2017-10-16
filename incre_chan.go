package main

import (
	"fmt"

	"time"
)
//var wg sync.WaitGroup

func main(){
	c1:=text("one")
	c2:=text("two")

	//fan in
	cn:=mearge(c1,c2)
	for i:=0;i<10;i++{     //why is it going in infinite loop if i use range keyword here instead of i
		fmt.Println(<-cn)
	}
}

func text(s string)chan string{
	c:=make(chan string)
	//for i:=0;i<5;i++{
		go func(){
			for i:=0;i<5;i++{
				c<-fmt.Sprintln("the process alive is",s," : ",i)
				time.Sleep(time.Second)
			}
			close(c)
		}()

	return c
}

func mearge(ch,chh chan string) chan string{
	cn:=make(chan string)
	go func(){
		for{
			cn<-<-ch
		}
	}()

	go func(){
		for{
			cn<-<-chh
		}
	}()
	return cn
}

//without channel
/*
func main() {

	wg.Add(2)
	go inc("1")
	time.Sleep(time.Microsecond)
	go inc("2")

	wg.Wait()

}

func inc(s string){
	//c:=make(chan string)

	go func(){
		for i:=0;i<5;i++{
			fmt.Println("Process",s," : ",i)
		}
		wg.Done()
	}()
}

*/
