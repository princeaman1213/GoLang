package main

import (
	"fmt"
	//"math/rand"
	"time"
	"math/rand"
)

var a[25]int
var x[25]int
var t int

func main() {


for k :=0;k<=24;k++{
	a[k]=rand.Intn(8)+8
	//fmt.Println(a[k])         //random nos for random delays for each tourist
}

	//x :=[8]int{rand.Intn(25)+1,rand.Intn(25)+1,rand.Intn(25)+1,rand.Intn(25)+1,rand.Intn(25)+1,rand.Intn(25)+1,rand.Intn(25)+1,rand.Intn(25)+1}
	//var x[25]int
	x =[25]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25}
	fmt.Println(x,a)

	for j:=0;j<=7;j++{
		fmt.Println("Tourist",x[j],"is online")
	}

	for i:=8;i<=24;i++{
		fmt.Println("Tourist",x[i],"is waiting")
	}

    queue(x,a)

	var temp int
	//
	fmt.Scanln(&temp)

}

func queue(x [25]int,a [25]int){

	for i:=0;i<=7;i++{
		go waitrandom(x[i],a[i],i)
	}

	//time.Sleep(time.Second*time.Duration(a))
	//fmt.Println("Alarm is armed")

}

func waitrandom(tourist,wait int,i int){
	time.Sleep(time.Second*time.Duration(wait))
	fmt.Println("Tourist ",tourist,"is done having spent",wait,"mins online")
    //go nextset(tourist)
	for i<=9{
		i++
	}
	go online(i)
	//go waitrandom(x[i],a[i],i)
}
/*
func nextset(done_tourist int){
	for i:=8;i<=16;i++{
		go waitrandom1(x[i],a[i],i)
	}
}

func waitrandom1(tourist,wait int,i int){
	time.Sleep(time.Second*time.Duration(wait))
	fmt.Println("Tourist ",tourist,"is done having spent",wait,"mins online")
	go nextset(tourist)
	//go waitrandom(x[i],a[i],i)
}*/

func online(i int){
	fmt.Println("Tourist",x[i],"is online")
	time.Sleep(time.Second*time.Duration(a[i]))
	fmt.Println("Tourist ",x[i],"is done having spent",a[i],"mins online")
}