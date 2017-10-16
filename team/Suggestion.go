package main

import (
	"fmt"

	"strings"
)

func main() {
	dictionary:=map[string]map[int]string{

		"a": map[int]string{
			1: "ability",
			2: "accept",
			3: "admire",
			4: "agency",
			5: "absence",
			6: "absent" ,
			7: "absoadlute",
		},

		"b":map[int]string{
			1: "backhand",
			2: "bankrupt",
			3: "beard",
			4: "behind",
			5: "bold",
			6: "ball",
			7: "band" ,
			8: "bandage",
		},

	}
	var input string
	fmt.Println("Enter the word or alphabet: ")
	fmt.Scan(&input)

	inputFirstAlphabet:= string(input[0])
	subLen:=len(input)

	inputSubString := input[0:subLen]


	fmt.Println("Required Suggestions: ")

	var flag int
	for i,n:=range dictionary[inputFirstAlphabet] {
		if(strings.Contains(n,inputSubString)) {
			fmt.Println(i, ": ", n)
			flag = 1
		}
	}
	if(flag==1) {
		var found int
		fmt.Println("Enter the number corresponding to required word: ")
		fmt.Scan(&found)
		fmt.Println("Required Word: ")
		fmt.Println(dictionary[inputFirstAlphabet][found])
	}else{
		fmt.Println("No Suggestions Found!")
	}
}
