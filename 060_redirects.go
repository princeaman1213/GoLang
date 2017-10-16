package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var t*template.Template

func init(){
	t=template.Must(template.ParseFiles("redirect1.gohtml"))
}

func main() {

	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.HandleFunc("/barred",barred)

	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func foo(w http.ResponseWriter,r *http.Request){
	fmt.Println("method(foo) is:",r.Method)
}

func bar(w http.ResponseWriter,r *http.Request){
	fmt.Println("method (bar) is:",r.Method)

	//http.Redirect(w,r,"/",303)
	//(above 1 line) OR (below 2 lines)
    w.Header().Set("Location","/")  //redirect to foo
    w.WriteHeader(http.StatusSeeOther)        //redirects to foo with method GET due to code 303(see other)
                                              //redirects without changing method if we use code 307 instead
}

func barred(w http.ResponseWriter,r *http.Request){
	fmt.Println("method (barred) is:",r.Method)
    t.ExecuteTemplate(w,"redirect1.gohtml",nil)

}
