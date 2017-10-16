package main
import "fmt"
func main() {
	num :=[]float64 {3,4}
	var y,z float64 = 8,9
	fmt.Println(f1())
	a,b := f2(num)                             //a,b takes the the 2 incoming return variables of f2()
	fmt.Println(a,b)
	fmt.Println(f3("sdfd",y,z))
}

func f1() (a int){                             //1 or more return variables can also be defined here
	a =1
	return
}

func f2(num []float64) (w float64,r float64){   //2 return variables declared here and returned
	w=num[0]+num[1]
	r=num[0]*num[1]
	return
}

func f3(a string,args ...float64) float64{                //args allows to take any no. of i/p arguments of the type defined
	var t float64 =0
	for _, v :=range args{
		t+=v
	}
	fmt.Println(a)
	return t
}
