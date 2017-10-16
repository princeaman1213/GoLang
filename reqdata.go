package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"html/template"
)

type user struct {
	Type string `json:"type,omitempty"`
	Model int `json:"model,omitempty"`
	Color string `json:"color,omitempty"`
}

//var obj = {empid:$('#empid').val(), name:$('#name'), mobile:$('#mobile'),address:$('#address'),position:$('#position')};
var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("reqdata.gohtml"))
}


func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8000",nil)


}

func index(w http.ResponseWriter,r *http.Request){

	var u user

//	var k string

	json.NewDecoder(r.Body).Decode(&u)
	//json.NewDecoder(r.Body).Decode(&k)


	fmt.Println("user" , u)
	//fmt.Println(k)
	//fmt.Println(u)

	t.ExecuteTemplate(w,"reqdata.gohtml",nil)

}
