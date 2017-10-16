package main

import "fmt"

func main() {
	c :=pipe(7,8,56)
	res :=square(c)

	 for r:=range res{
		fmt.Println(r)
	    //fmt.Println(<-res)
	    //fmt.Println(<-res)
     }

}

func pipe(nums ...int) chan int{
	c1:=make(chan int)
	go func(){
		for _,r:=range nums{
			c1<-r
		}
		close(c1)
	}()
	return c1
}

func square(c chan int) chan int{
	out :=make(chan int)
	go func(){
		for r:=range c{
			out<-r*r
			//fmt.Println(sq)
		}
		close(out)
	}()
	return out
}
