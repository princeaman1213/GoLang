package main

import (
	"sort"
	"fmt"
)

type person []string
type inte []int

//for person, attach these three methods to sort a user defined data
func(a person) Len() int{
	return len(a)
}

func(a person) Less(i,j int) bool{
	return a[i]<a[j]                         //sort in increasing order
}

func(a person) Swap(i,j int){
	a[i],a[j]=a[j],a[i]
}

func main() {

	p1:=person{"Aman","Mukul","Ambuj","Banga","Shubham"}
	p3:=[]string{"Aman","Mukul","Ambuj","Banga","Shubham","Kiran","ganesh"}
	sort.Sort(sort.Reverse(sort.StringSlice(p3)))
	//sort.Strings(p3)
	fmt.Println("p3 : ",p3)
	p2:=inte{5,3,7,8,2}
	p4:=[]int{34,66,22,77,39}
	sort.Sort(p1)
	sort.Ints(p4)      // or use this , sort.Sort(sort.IntSlice(p4))
	fmt.Println("p4 : ",p4)
	fmt.Println("p1 : ",p1)
	sort.Sort(p2)
	fmt.Println("p2 : ",p2)

	//type
	fmt.Printf("%T",sort.Reverse(sort.StringSlice(p1)))
	fmt.Printf("\n%T",sort.StringSlice(p1))

}
//for inte, attach these three methods to sort a user defined data
func(a inte) Len() int{
	return len(a)
}

func(a inte) Less(i,j int) bool{            //sort inreverse order
	return a[i]>a[j]
}

func(a inte) Swap(i,j int){
	a[i],a[j]=a[j],a[i]
}