package main
import "fmt"
func main() {
    //x :=[]float64 {43 ,99 ,99 ,97 ,88}
	var xa[5]float64
	for i :=0; i<len(xa) ;i++{                             // Why not able to take input and run the same logic ??
		fmt.Scanf("%f",&xa[i])
	}
    //x :=xa[:]                                        // able to take input , save it in another array and then pass it

	fmt.Println(xa)
	fmt.Println(len(xa))
	fmt.Println(avg(xa))

}

func avg(x [5]float64) float64{
	var t float64=0
	for _, v :=range x{
		t+=v
	}
	return t/float64(len(x))
}
