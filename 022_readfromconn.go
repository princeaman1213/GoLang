package main

import (
	"net"
	"fmt"
	"bufio"
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
		go handleconn(c)
	}
}

func handleconn(c net.Conn){
/*	err:=c.SetDeadline(time.Now().Add(time.Second*5))       //video no. : 023
	if err!=nil{
		fmt.Println(err)
	}
*/
	scanner:=bufio.NewScanner(c)
    //scanner.Split(bufio.ScanWords)       //to print each word in new line
	for scanner.Scan(){
		data:=scanner.Text()
		fmt.Println(data)
		fmt.Fprintln(c,"Received : ",data)         // video no.  023
	}
    defer c.Close()
    fmt.Println("End Of Prog")          //program reaches here when we close the connection i.e close localhost 8080
}
