package main

import (
	"fmt"

)

type errortest struct {
	lat,long string
	err error
}

func (rec *errortest) Error() string{        // Any type with this signature implicitly implements the error interface
	return fmt.Sprintf("error occured %v %v %v",rec.lat,rec.long,rec.err)
}

func main() {
	_,err:=sqrt(-12)
	if err!=nil{
		fmt.Println(err)
	}

	//fmt.Printf("%T \n %T",fmt.Errorf("sq root of neg num. %v"),fmt.Sprintf("sq root of neg num. %v"))
	//fmt.Println("\n 	new update ")
}

func sqrt(n int) (int , error){
	if n<0{
		err1:=fmt.Errorf("cannot calculate sq root of neg num. %v",n)

		return 0,err1
		//return 0,&errortest{"30 N","43 E",err1}
		//return 0,errors.New("cannot calculate sq root of neg num. %v")
	}else {
		return 0,nil
	}

}
