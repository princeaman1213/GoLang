package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age int  `json:"wisdom"` //tags               //start with caps to export and thus print data after using marshal
}

func main() {
	p1 :=person{"AMan","Patel",21}
    //p2 :=person{"acd","gupta",22,2}

	bs,_ :=json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T",bs)
	fmt.Println("\n",string(bs))


}
