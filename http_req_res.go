package main

import (
	"net"
	"fmt"
	"bufio"

	"strings"
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
			continue
		}
		go handleconn(c)
	}
}

func handleconn(c net.Conn){
    defer c.Close()
	request(c)

	respond(c)
}

func request(c net.Conn){
    var i int
	scanner:=bufio.NewScanner(c)
	//scanner.Split(bufio.ScanWords)       //to print each word in new line
	for scanner.Scan(){
		data:=scanner.Text()
		fmt.Println(data)
		//fmt.Fprintln(c,"Received : ",data)         // video no.  023
        if i==0 {
			fmt.Println("Method : ",strings.Fields(data)[i],"URI :",strings.Fields(data)[i+1])

		}
		if data==""{
			break
		}
		i++
	}
	//defer c.Close()
	fmt.Println("End Of Prog")          //program reaches here when we close the connection i.e close localhost 8080
}

func respond(c net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello GoLang Prog</strong></body></html>`

	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")                          //to respond with ok status
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))       //optional
	fmt.Fprint(c, "Content-Type: text/html\r\n")                  //optional
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)                                              //any html response
}