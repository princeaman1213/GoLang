package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"io"
)

func main() {

	a:=getcode("example")
	fmt.Println(a)
	b:=getcode("example1")
	fmt.Println(b)
}

func getcode(str string) string{

	h:=hmac.New(sha256.New,[]byte("passkey"))
	io.WriteString(h,str)             //write our string into the hash

	return fmt.Sprintf("%x",h.Sum(nil))     //this line has no explanation but its necessary

}
