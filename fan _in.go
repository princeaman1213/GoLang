package main

import "fmt"

func main() {
	c :=fanin(str("ice"),str("fire"))  //c:=one(c,c)

	for i:=0;i<13;i++{
		fmt.Println(<-c)
	}

}

func str(s string) chan string{
	c :=make(chan string)
	go func(){
		for i:=0;i<8;i++{
			c<-fmt.Sprintln(s,"is good",i)
		}
	}()
	return c
}

func fanin(c1,c2 chan string) chan string{
	c3 :=make(chan string)
	go func(){
		for {
			c3<-<-c1
		}

	}()

	go func(){
		for {
			c3<-<-c2
		}
	}()

	return c3

}