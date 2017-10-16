package main

import (
	"text/template"
	"os"
	"fmt"
)

func main() {
	f:=template.Must(template.ParseFiles("passing.gohtml","passingvar.gohtml"))

	err:=f.ExecuteTemplate(os.Stdout,"passing.gohtml","hey 22")
	if err!=nil{
		fmt.Println(err)
	}

	err=f.ExecuteTemplate(os.Stdout,"passingvar.gohtml",`variable pass`)
	if err!=nil{
		fmt.Println(err)
	}

}
