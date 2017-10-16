package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var a,b,c,d int
	fmt.Println("Let's go for a walk")
	fmt.Println("Jon started getting ready")
	fmt.Println("Jack started getting ready")

	a =rand.Intn(30)+60
	b =rand.Intn(30)+60
	c =rand.Intn(10)+35
	d =rand.Intn(10)+35

	if a>=b {
		jackready(a)
		go jonready(b)
	}else if b>a{
		go jackready(a)
		jonready(b)
	}

	fmt.Println("Arming Alarm")

	go func(){
			time.Sleep(time.Second*60)
			fmt.Println("Alarm is armed")
	}()

	fmt.Println("Jon started putting on shoes")
	fmt.Println("Alarm is counting down")
	fmt.Println("Jack started putting on shoes")

	if c>=d{
		jackshoes(c)
		go jonshoes(d)
	}else if d>c{
		go jackshoes(c)
		jonshoes(d)
	}

	fmt.Println("Exiting and locking the door")
    var temp int
	fmt.Scanln(&temp)

}

func jackready(a int){
	time.Sleep(time.Second*time.Duration(a))
	fmt.Println("Jack spent",a,"seconds getting ready")
}

func jonready(b int){
	time.Sleep(time.Second*time.Duration(b))
	fmt.Println("Jon spent",b,"seconds getting ready")
}

func jackshoes(c int){
	time.Sleep(time.Second*time.Duration(c))
	fmt.Println("Jack spent",c,"seconds putting on shoes")
}

func jonshoes(d int){
	time.Sleep(time.Second*time.Duration(d))
	fmt.Println("Jon spent",d,"seconds putting on shoes")
}
