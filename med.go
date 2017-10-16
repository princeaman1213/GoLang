package main

import "fmt"

type med struct {
	name string
	tablets float64
	mgpertab float64
}


func main() {

	//var obj []med
	obj:=[]med{
		{"med1", 1,5.5},
		{"med2", 2,4.6},
		{"med3", 3,9.0},
		{"med4", 4,2.5},
	}
	fmt.Println(obj[0])

	leng:=len(obj)
	for i:=0;i<leng;i++{
		if (obj[i].tablets*obj[i].mgpertab<10) {
			fmt.Println("Tablet ", obj[i].name, " is safe to use")
		}else {
			fmt.Println("Tablet ", obj[i].name, " is not safe to use")
		}
	}

}
