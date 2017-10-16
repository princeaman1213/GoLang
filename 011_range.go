package main

import (
	"text/template"
	"os"
	"fmt"
)

var f *template.Template

func init(){
	f=template.Must(template.ParseFiles("passingrange.gohtml"))
}

func main() {

	arr:=[]string{"a","b","c"}
	err:=f.Execute(os.Stdout,arr)
	if err!=nil{
		fmt.Println(err)
	}

}
