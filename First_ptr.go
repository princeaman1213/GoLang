package main
import "fmt"
func main() {
     x :=4
	 fun(&x)

	fmt.Println(x)
	fmt.Println(&x)

}

func fun(xptr *int){            //passing by reference
	*xptr=2
}