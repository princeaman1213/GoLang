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
			mux(c, data)
		}
		if data == ""{
			break
		}
		i++
	}
	//defer c.Close()
	fmt.Println("End Of Prog")          //program reaches here when we close the connection i.e close localhost 8080
}

func mux(c net.Conn,data string){

	method:=strings.Fields(data)[0]
	uri:=strings.Fields(data)[1]
	fmt.Println("methos is : ",method,"uri is : ",uri)
	
	//mux
       if method == "GET" && uri =="/"{
       	index(c)
     	}
	if method == "GET" && uri =="/about"{
		about(c)
	}
	if method == "GET" && uri =="/contact"{
		contact(c)
	}
	if method == "GET" && uri == "/apply" {
		apply(c)
	}
	if method == "POST" && uri == "/apply" {
		applyProcess(c)
	}

}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(c net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="POST" action="/apply">
	<input type="text" name="fname" placeholder="enter first name">
	<input type="submit" value="apply">
	</form>
	</body></html>`

	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}

func applyProcess(c net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprint(c, body)
}