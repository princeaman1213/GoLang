package main

import "fmt"

type link1 struct{
	name string
	age int
	ptr * link1
}

var temp link1

func main() {
	ravi :=link1{"ravi",22,nil}
	paul :=link1{"paul",23,&ravi}
	aman :=link1{"aman",21,&paul}
	head :=link1{"start",0,&aman}
	//temp :=link1{}
	//temp.ptr = head.ptr.ptr

//	fmt.Println(head)
//	fmt.Println(*temp.ptr)

	//print(&head)

	insert(&head,"abc",26)
	insert(&head,"mukul",24)

	print(&head)


}

func print(head *link1) {
    //temp :=link1{}
    //temp = head
	for head.ptr !=nil{

	          fmt.Print(head.ptr.name,"\t",head.ptr.age,"\n")
              head.ptr = head.ptr.ptr


	}
	fmt.Println(2)

}

func insert(head *link1,name string,age int) {
	//temp1 :=link1{}
	//temp1.ptr = head.ptr
	for head.ptr !=nil{

		fmt.Print(head.ptr.name,"\t",head.ptr.age,"\n")
		head.ptr = head.ptr.ptr

	}
	temp =link1{name,age,nil}
	//print(head)
	head.ptr = &temp
	//fmt.Println("  \n \n")
	//fmt.Println("%p",head)
   // fmt.Println(head.ptr)
	//print(head)
}