package main
import "fmt"
func main() {
	go f(0)                 //to execute this function in background and proceed with next statements
	var input int
	fmt.Scanf("%d",&input)
	fmt.Println(input)
}

func f(n int){
	for i :=0;i<50;i++{
		fmt.Println(n,":",i)
	}
}
