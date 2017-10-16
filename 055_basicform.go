package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("basicform.gohtml"))
}

type detail struct {
	Fname string
	Lname string
	Sub bool
}

func main() {
	http.HandleFunc("/",f)

	//t=template.Must(template.ParseFiles("basicform.gohtml"))   //this also works in place of init() function

	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func f(w http.ResponseWriter,r *http.Request){
	v1:=r.FormValue("first")
//	v2:=r.FormValue("last")
//	v3:=r.FormValue("sub") == "on"

	fmt.Println("v1:",v1)

	err:=t.ExecuteTemplate(w,"basicform.gohtml",nil)
	if err!=nil{
		fmt.Println(err)
	}
}
