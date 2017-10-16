package main

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Person struct{
	Username string `json:username,omitempty`
	Password string `json:password,omitempty`
	EmpId string     `json:empid,omitempty`
	Address string   `json:address,omitempty`
	Position string    `json:position,omitempty`
	Personname string  `json:personname,omitempty`

}

var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("112_ajaxserver.html","112_print.html"))
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/foo",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter,r *http.Request){

	t.ExecuteTemplate(w,"112_ajaxserver.html",nil)

}

func foo(w http.ResponseWriter,r *http.Request){

	//io.WriteString(w,"my foo page !")
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	fmt.Fprintln(w, sbs)
	fmt.Println("json is : ",sbs)

	var p Person

	json.Unmarshal(bs,&p)

	fmt.Println("p is ",p)
	/*t.ExecuteTemplate(w,"112_print.html",p)
	fmt.Fprintln(w,p.Username)*/

}