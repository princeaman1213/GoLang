package main
import(
	"fmt"
	"time"
)

func pinger(c chan <- string){            //<- makes it unidirectional (only send , no receive)
	for i :=0; ;i++{
		c <- "ping"
		time.Sleep(time.Second*3)
	}
}

func ponger(c1 chan string){
	for i :=0; ;i++{
		c1 <- "pong"
		time.Sleep(time.Second*5)
	}
}

func printer(c,c1 chan string){
	/*for{
        msg := <- c
		//msg1 := <-d
		fmt.Println("The info is :     ",msg)
		time.Sleep(time.Second*1)
	}*/

	for{
		select{                     //select is same as switch case but only for channels
		case msg1 := <- c:
			fmt.Println(msg1)
		case msg2 := <-c1:
			fmt.Println(msg2)
		case <-time.After(time.Second*2):
			fmt.Println("timeout" , time.After(time.Second*2))

		}
	}

}

func main() {
	var c chan string=make(chan string)               //using a channel between 2 goroutines to sync them
	var c1 chan string=make(chan string)
	//var d chan string=make(chan string)

	go pinger(c)
	go ponger(c1)
	go printer(c,c1)

	var input string
	fmt.Scanln(&input)
}
