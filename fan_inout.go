package main

import (
	"fmt"
	"sync"
)

func main() {

	a:=mul2(4,6)

	//for r:=range a{
	//	fmt.Println(r)
	//}

	//Fan out
	c2:=sq(a)
	c3:=sq(a)

	//Fan in
	c4:=mearge(c2,c3)

	for r1:=range c4{
		fmt.Println(r1)
	}


}

func mul2(num ...int) chan int{
	c:=make(chan int)
	go func(){
		for _,r:=range num{
			c<-r
		}
		close(c)
	}()
	return c
}

func sq(c chan int) chan int{
	c1:=make(chan int)
	go func(){
		for r:=range c{
			c1<-r*r
		}
		close(c1)
	}()
	return c1
}

func mearge(c5 ...chan int) chan int{
	c6:=make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(c5))
	for _,r:=range c5{             //every value (r) of this chan c5 is itself a chan as c5 is a slice of channels
		go func(c9 chan int){
			for r8:=range c9{
				c6<-r8
			}
              wg.Done()
		}(r)
	}

	go func() {
		wg.Wait()
		close(c6)
	}()

	return c6
}