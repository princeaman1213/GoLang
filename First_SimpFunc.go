package main
import "fmt"
func main() {
	var a float64=33
	var b float64=21
	fmt.Println(test(a,b))

}

func test(args ...float64) float64{     //dealing with > 1 input arguments
	var t float64 = 0
	for _, v :=range args{
		t+=v
	}
	return t
}
