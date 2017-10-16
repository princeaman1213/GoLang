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

	scanner:=bufio.NewScanner(c)
	//scanner.Split(bufio.ScanWords)       //to print each word in new line
	for scanner.Scan(){
		data:=scanner.Text()
		fmt.Println("Word is : ",data)
		fmt.Println("Encrypted word is : ",rot13(data))
	}
	defer c.Close()
	fmt.Println("End Of Prog")          //program reaches here when we close the connection i.e close localhost 8080
}

func rot13(data string) string{
	bs:= []byte(data)
	var rot=make([]byte,len(bs))
	for i,v:=range bs{
		if v<109{
			v+=13
		}else{
			v-=13
		}
		rot[i]=v
	}
   return string(rot)
}