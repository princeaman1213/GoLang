package main

import (
	"net/http"
	"html/template"

	"fmt"
	"io/ioutil"
)

var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("jsontest.html"))
}

func main() {

	http.HandleFunc("/abcd",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8000",nil)

}

func index(w http.ResponseWriter,r *http.Request){

	//w.Header().Set("Content-Type","text/html; charset=utf-8")
    //w.Header().Add("Origin")
	//t.ExecuteTemplate(w,"jsontest.html",nil)

	//fmt.Println("dsf")
//	var u string
	bs,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("Error is : ",err)
	}
	fmt.Println(string(bs))



}
