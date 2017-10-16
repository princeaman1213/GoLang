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
			4 : "analogy",
			5 : "admire",
			6 : "alarity",
			7 : "abscond",
			8 : "apathy",
		},
		"b" : map[int]string{
			1 : "bad",
			2 : "ball",
			3 : "base",
			4 : "bonquet",
			5 : "bully",
			6 : "bunk",
		},
	}
	var l string
	var index int
	fmt.Println("Enter some(1 or more) letters of your word :")
	fmt.Scanf("%v",&l)
	t:=string(l[0])
	fmt.Println("Suggestions are listed as follows...")
	for i,v:=range db[t]{
		if strings.HasPrefix(v,l){
			fmt.Println(i," : ",v)
		}
	}

	fmt.Println("Enter the index of the word to be printed :")
	fmt.Scanf("%v",&index)

	if index<=len(db[t]){
		if strings.HasPrefix(db[t][index],l){
			fmt.Println("Selected word is...\n",db[t][index])
		}else{
			fmt.Println("Please enter the correct index")
		}
	}


}
