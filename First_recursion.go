package main
import "fmt"
func main() {
	var num int=0
	fmt.Println("Enter the num :")
	fmt.Scanf("%d",&num)
	fmt.Println(fact(num))

}

func fact(num int) int{
	var f int
	if num==0{
		return 1
	}
	f = num*fact(num-1)
	//fmt.Println(f)
	return f                                      //factorial by recursion

}