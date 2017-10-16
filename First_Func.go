package main
import "fmt"
func main() {
    //x :=[]float64 {43 ,99 ,99 ,97 ,88}
	var xa[5]float64
	for i :=0; i<5 ;i++{                             // Why not able to take input and run the same logic ??
		fmt.Scanf("%f",&xa[i])
	}
    x :=xa[2:4]                                        // able to take input , save it in another array and then pass it
	//fmt.Println(x)
	//fmt.Println(len(x))
	fmt.Println(avg(x))

}

func avg(x []float64) float64{
	var t float64=0
	fmt.Println(x)
	for _, v :=range x{
		t+=v
	}
	fmt.Println(t)
	return t/float64(len(x))
}
