package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	s:="In the following sections, I'll use the same database table structure for different databases, then create SQL as follows"

	//encodestr:="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+/"   //THIS IS FIXED
	//b64:=base64.NewEncoding(encodestr).EncodeToString([]byte(s))
//above 2 lines or below 1 line

	b64:=base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(b64))
	fmt.Println(s)
	fmt.Println(b64)

	str,_:=base64.StdEncoding.DecodeString(b64)
	fmt.Println(string(str))

}
