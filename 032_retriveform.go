package main

import (
	"net/http"
	"fmt"
	"html/template"
)

type dog int
var t *template.Template
func(d dog) ServeHTTP(w http.ResponseWriter,r *http.Request){  //(lne1)    //this is the signature of handler interface
	   err:=r.ParseForm()
	   if err!=nil{
	   	fmt.Println(err)
	   }
	   t.ExecuteTemplate(w,"index.gohtml",r.Form)      //use r.postform to only get form data in the webpage
       //fmt.Fprintln(w,"Any code")
}



func init(){
	t=template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d dog
	http.ListenAndServe(":8080",d)        //d here is handler so it is handled by (lne1)
}

//what is the difference in using get or post