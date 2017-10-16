package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)

	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func foo(w http.ResponseWriter,r *http.Request){
	fmt.Println("method(foo) is:",r.Method)
}

func bar(w http.ResponseWriter,r *http.Request){
	fmt.Println("method (bar) is:",r.Method)

	http.Redirect(w,r,"/",301)      //will not run now untill we clear the browser cache
}
