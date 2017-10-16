package main
import "fmt"
func main() {
	arr :=[]float64 {2,3,4,2,5}
	x :=make([]float64,3)
	x[0]=1
	x[1]=2
	x[2]=5
	//x[3]=7
	//x[4]=1
	//a :=arr[0:5]
	fmt.Println(arr)
	fmt.Println(x)

	slice()

}

func slice(){
	a :=[]float64{1,3,2,4,3}
	x :=a[1:3]
	fmt.Println(x)
}
