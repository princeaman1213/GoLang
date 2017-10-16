package main

import "fmt"

type link1 struct{
	name string
	age int
	ptr * link1
}

var head link1

//var temp link1


func insert(head *link1) {
	for head !=nil{

		//fmt.Print(head.name,"\t",head.age,"\n")
		head = head.ptr
	}

	temp :=link1{"fdgffg",29,nil}
	head = &temp
	//print(head)
}



func print(head *link1) {
	//temp :=link1{}
	//temp = head
	for head !=nil{

		fmt.Print(head.name,"\t",head.age,"\n")
		head = head.ptr


	}
	//fmt.Println(2)

}

func main() {


	print(&head)
	fmt.Println("\n gap \n")

	insert(&head)

	print(&head)


}






