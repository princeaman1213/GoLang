package main
import(
	"fmt"
	"time"
	"math/rand"
)

func pinger(a,b int,c chan string){            //<- makes it unidirectional (only send , no receive)

	fmt.Println("Tourist",a,"is online")
	time.Sleep(time.Second*time.Duration(b))
	fmt.Println("Tourist ",a,"is done having spent",b,"mins online")
		//c <- "ping"

}


func printer(c chan string){
	for{
		//msg := <- c
		//msg1 := <-d
		//fmt.Println("The info is :     ",msg)
		time.Sleep(time.Second*1)
	}
}

func main() {
	var a[25]int
	var x[25]int
	var c chan string=make(chan string)               //using a channel between 2 goroutines to sync them

	x =[25]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25}
	fmt.Println(x)

	for k :=0;k<=24;k++{
		a[k]=rand.Intn(8)+8
		//fmt.Println(a[k])         //random nos for random delays for each tourist
	}

	for i:=0;i<8;i++{
		go pinger(x[i],a[i],c)
	}

	//go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
