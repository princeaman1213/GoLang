package main
import "fmt"
func main() {
	var x[5] float64
	var total,avg float64 = 0,0
	//x[0]=3
	//fmt.Println(x[0]*6)

	for i:=0; i<5 ;i++ {
	fmt.Scanf("%f" , &x[i])

	}

	//for j:=0; j<5 ;j++ {
	//	total +=x[j]

	//}
	for _, v :=range x {
		total +=v

	}                               //another way of writing a loop


	avg=total/5
	fmt.Println(avg)

}
