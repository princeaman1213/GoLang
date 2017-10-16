package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

var wg sync.WaitGroup
var count int
var mu sync.Mutex

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bofish")
	wg.Wait()
	fmt.Println("Final counter :",count)

}

func increment(s string){
	for i:=0;i<20;i++{

		//x:=count
        //x++
		time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)
		mu.Lock()
		//count=x
		count++
		fmt.Println(s,i,"counter",count)
		mu.Unlock()
	}
	wg.Done()
}
