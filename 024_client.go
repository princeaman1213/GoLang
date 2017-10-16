package main

import (
	"net"
	"fmt"
)

func main() {
	c,err:=net.Dial("tcp","localhost:8080")
	if err !=nil{
		fmt.Println(err)
	}
    defer c.Close()
	fmt.Fprintln(c,"I dialed")
}
