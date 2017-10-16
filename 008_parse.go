package main

import (
	"text/template"
	"fmt"
	"os"
)

func main() {
	f,err :=template.ParseFiles("tpl.gohtml") 
	if err!=nil{
		fmt.Println(err)
	}
/*
	file,err1:=os.Create("newfile.html")
	if err!=nil{
		fmt.Println(err1)
	}
*/
	err=f.Execute(os.Stdout,nil)    //  use file name in place of os.stdout to write directly in file
	//err=f.ExecuteTemplate(os.Stdout,"tpl.gohtml",nil)         //to execute a single file if more than 1 files are parsed in f above
	if err!=nil{
		fmt.Println(err)
	}


}
