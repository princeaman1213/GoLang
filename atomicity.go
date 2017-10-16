package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
	"sync/atomic"
)

var wg sync.WaitGroup
var count int64
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
        atomic.AddInt64(&count,1)
		fmt.Println(s,i,"counter",count)

	}
	wg.Done()
}
