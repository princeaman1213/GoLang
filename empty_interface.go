package main

import "fmt"

type shapes interface {}

//type shape struct {
//	a int
//}


type circle struct{
	//shape
	x,y,r float64
}

type rect struct{
	//shape
	a,b float64
}

func main() {
	c :=circle{0,0,4}
	c1 :=circle{}
	r :=rect{4,5.26}

	cir :=[]circle{c,c1,circle{0,0,54}}

	for _,v :=range cir {
		fmt.Println(v)
	}

	//An empty interface is created is created so that we can store different types of shapes , i.e circle,rect

		shapes :=[]shapes{c,c1,r}  // or shapes :=[]interface{}{c,c1,r} and no need to define the shape interface at the top
    fmt.Println("\n")
	for _,v :=range shapes{
		fmt.Println(v)
	}



}



