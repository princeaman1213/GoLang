package main

import (
	"net"
	"fmt"
	"encoding/gob"
)

func main() {
	go server()
	go client()

	var ip int
	fmt.Scanln(&ip)

}

func server(){
	l,err :=net.Listen("tcp",":9999")  //listen on port
	if err != nil{
		fmt.Println(err)
		return
	}

	for{
		c,err :=l.Accept()
		if err!= nil{
			fmt.Println(err)
			continue
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn){
	var msg string
	err :=gob.NewDecoder(c).Decode(&msg)  //receive msg
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Received",msg)
	}
	c.Close()
}

func client(){
	c,err :=net.Dial("tcp","127.0.0.1:9999")
	if err !=nil{
		fmt.Println(err)
		return
	}
	msg :="Hello World"
	fmt.Println("sending",msg)
	err =gob.NewEncoder(c).Encode(msg)
	if err!= nil{
		fmt.Println(err)
	}
	c.Close()

}