package main

import (
	"net/http"
	"html/template"
	"io"
)

var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("112_ajaxserver2.html"))
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/foo",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter,r *http.Request){

	t.ExecuteTemplate(w,"112_ajaxserver2.html",nil)

}

func foo(w http.ResponseWriter,r *http.Request){

	io.WriteString(w,"my foo page !")

}