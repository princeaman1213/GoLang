package main
import "fmt"
func main() {
	x :=make(map[string]int)
	q :=make([]int,3)
	x["a"]=3
	x["c"]=4
	q[0]=7
	q[1]=4
	q[2]=3
	fmt.Println(x)
	fmt.Println(q)
}
