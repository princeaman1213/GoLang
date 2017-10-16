package main

import "fmt"

func main() {
	var p *int
	a :=32
	p =&a
	fmt.Println(*p,p)
}
