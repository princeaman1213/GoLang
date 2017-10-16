package main

import (
	"fmt"
	//"time"
	"time"
)

func main() {
	c :=make(chan int)
	done :=make(chan bool)
	n:=2

		go func() {
			for i := 0; i < 5; i++ {
				c <- i
				time.Sleep(time.Second)
			}
			close(c)
		}()


	for i:=0;i<n;i++{
		go func() {
			for n:=range c{
				fmt.Println(n)
			}
			done<-true
		}()
	}

	for i:=0;i<n;i++{
		<-done
	}

}
