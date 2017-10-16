package main

import (
	"text/template"
	"os"
	"fmt"
)

type bikes struct{
	Name string
	Model int
}

type cars struct{
	Name string
	Model int
}

type vehicles struct {
	Bike1 []bikes
	Car1 []cars
}

var file *template.Template

func init(){
	file = template.Must(template.ParseFiles("passstruct.gohtml"))
}

func main() {
	b1:=bikes{"Fz",12345}
	b2:=bikes{"Duke",67896}

	c1:=cars{"i10",7476}
	c2:=cars{"santro",5789}

	bike:=[]bikes{b1,b2}
	car:=[]cars{c1,c2}

    v:=vehicles{bike,car}

    //we are passing a slice of structs "v" here , so we range over this slice in the .gohtml file to get individual structs and then print out
    //their value

    err:=file.Execute(os.Stdout,v)
    if err!=nil{
    	fmt.Println(err)
	}

}
