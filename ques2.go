package main
import(
	"fmt"
	"time"
	"math/rand"
)

var a[25]int
var x[25]int
var t int

func main() {
	x =[25]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25}

	for k :=0;k<=24;k++{
		a[k]=rand.Intn(105)+15
	}

	for i :=0;i<8;i++ {
		go cafe(i)
	}
    time.Sleep(time.Millisecond)
	for v :=8;v<25;v++ {
		fmt.Println("Tourist",v+1,"is waiting for turn")
	}
	var input string
	fmt.Scanln(&input)
}

func cafe(i int){
	fmt.Println("Tourist",x[i],"is online")
	time.Sleep(time.Second*time.Duration(a[i]))
	fmt.Println("Tourist ",x[i],"is done , having spent",a[i],"mins online")
    t++
		if i>=0 && i<17 {
			cafe(i+8)
		}
	if t==25{
		fmt.Println("\nThe place is empty, let's close up and go to the beach!")
	}
}