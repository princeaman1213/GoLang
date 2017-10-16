package main

import "fmt"

func main() {
	//var a string = "A"
	//fmt.Println(int(a))

	char := 'a' // rune, not string
	//fmt.Printf("%T",char)
	ascii := int(char)
	fmt.Println(string(char), " : ", ascii)
	//fmt.Printf("%T",char)

	//var s int = 67
	//fmt.Println("       " ,string(s))

	var str rune = 'z'
	fmt.Println(int(str))

}
