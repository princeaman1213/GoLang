package main

import (
	"text/template"
	"os"
	"fmt"
)

var f *template.Template

func init(){
	f=template.Must(template.ParseFiles("define.gohtml","new.gohtml"))
}

func main() {

	err:=f.ExecuteTemplate(os.Stdout,"define.gohtml",33)
	if err!=nil{
		fmt.Println(err)
	}

}
