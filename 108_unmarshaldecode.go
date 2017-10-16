package main

import (
	"encoding/json"
	"fmt"
)

type detail struct {
	Name string
	ZIP int
}

func main() {
    var decoded [2]detail
	str:=`[{"Name":"Delhi","ZIP":110078},{"Name":"Noida","ZIP":201301}]`
    err:=json.Unmarshal([]byte(str),&decoded)
    if err!=nil{
    	fmt.Println(err)
    	return
	}

	fmt.Println(decoded)

	for i,v:=range decoded{
		fmt.Println(i,v)
	}

}
