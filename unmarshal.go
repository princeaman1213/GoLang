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
		var p1 person
	//p2 :=person{"acd","gupta",22,2}
    bs :=[]byte(`{"Fname":"Aman" , "Lname":"Patel" , "wisdom":21}`)
	json.Unmarshal(bs,&p1)

	fmt.Println(p1)



}
