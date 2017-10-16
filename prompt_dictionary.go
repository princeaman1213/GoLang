package main

import (
	"fmt"
	"strings"
)

func main() {
	db :=map[string]map[int]string{
		"a" : map[int]string{
			1 : "abate",
			2 : "alarm",
			3 : "algorithm",
			5 : "analogy",
		},
		"b" : map[int]string{
			1 : "bad",
			2 : "ball",
			3 : "base",
		},
	}
	//fmt.Println(db[0]["1"])
	var l string
	var index string
	fmt.Println("Enter the first letter :")
	fmt.Scanf("%v",&l)
	l = strings.ToLower(l)

		//fmt.Println(db[0])
		for i,v:=range db[l]{
			fmt.Println(i," : ",v)
		}

	fmt.Println("Enter the index of word to be printed :")
	fmt.Scanf("%v",&index)

		//fmt.Println(db[l][index])

	for j,v1:=range db[l]{
		flag :=strings.HasPrefix(db[l][j],index)
		if flag == true{
			fmt.Println(j," : ",v1)
		}

	}

}
