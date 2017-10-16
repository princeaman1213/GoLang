package main

import (
	"sort"
	"fmt"
)

type person struct{
	name string
	age int
}

type byname []person

func (this byname) Len() int{
	//fmt.Println(len(this))          // no of records of struct
	//fmt.Println(this)
	return len(this)
}

func (this byname) Less(i,j int) bool{
	return this[i].name < this[j].name        // change .name to .age to sort by age
}

func (this byname) Swap(i,j int){
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids :=[]person{

		{"Paul",22},
		{"Aman",21},
	}
	sort.Sort(byname(kids))
	fmt.Println(kids)

}
