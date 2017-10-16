package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	l,err:=net.Listen("tcp",":8080")
	if err!=nil{
		fmt.Println(err)
	}

	for{
		c,err:=l.Accept()              //here we accept the tcp connection in c and now we can read and write on this connection
		if err!=nil{
			fmt.Println(err)
		}
        n:=3
		//different ways of writing data on a connection
		io.WriteString(c,"First tcp connection....!")
		fmt.Fprintf(c,"%v\n",n)
		fmt.Fprintln(c,"Last line")

		c.Close()
	}

}
