package main

import (
	    "fmt"
	    //"math"
	    "strings"
)

func main() {
	greetings :=  []string{"Hello","world!"}
	fmt.Println(strings.Join(greetings,"   +    "))
}
