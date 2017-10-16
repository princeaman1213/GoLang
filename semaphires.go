package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	//var wg sync.WaitGroup
	//wg.Add(2)
	go func() {
		//wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		done<-true
	}()

	go func() {
		//wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		done<-true
	}()



	go func(){
		<-done
		<-done
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	//wg.Wait()
	//close(c)



}