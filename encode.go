package main

import (
	"encoding/json"

	"os"
	"strings"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Age int               //start with caps to export and thus print data after using marshal
}

func main() {
	p1 :=person{"AMan","Patel",21}
	//p2 :=person{"acd","gupta",22,2}

	json.NewEncoder(os.Stdout).Encode(p1)       //encode
var i int
	//json.NewEncoder(os.Stdout).Encode(i)
	var p2 person
	t :=strings.NewReader(`{"Fname":"AMan","Lname":"Patel","Age":21}`)
	json.NewDecoder(t).Decode(&p2)               //decode
	fmt.Println(p2)

	t1 :=strings.NewReader(`2`)
	json.NewDecoder(t1).Decode(&i)       //decode
	fmt.Println(i)


}
