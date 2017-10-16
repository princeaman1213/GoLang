package main

import (
	"fmt"

)
//var wg sync.WaitGroup
var done chan bool
func main() {

    //var num int
	//fmt.Scanln(&num)
	//fmt.Println(math.MaxInt64)
	//n :=10
	//fmt.Printf("%T",n)
   // wg.Add(n)
    c :=gen()
	c1 :=fact(c)

	for r:=range c1{
		fmt.Println(r)
	}
	//fmt.Println(<-fact(6))
	//wg.Wait()
   // time.Sleep(time.Second*22)

}

func gen()chan int{
	c:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			c<-i+1
			//time.Sleep(time.Second)
		}
		close(c)
	}()
	//fmt.Println("dfssadf",<-c)
	//fmt.Println("dfssadf",<-c)
	return c
}

func fact(n chan int) chan int{
	c2 :=make(chan int)

	go func(){
		for r:=range n{
			//f=f*(r+1)
			c2<-fact1(r)
			//fmt.Println("rrrrrr",r)
		}
		//c2<-f
		close(c2)
	}()
	//fmt.Println("fact of", n ," is :",<-c)
    return c2
	//wg.Done()
}

func fact1(n1 int) int{
	var f int=1
	for k:=n1;k>0;k--{
		f*=k
	}
	return f
}
