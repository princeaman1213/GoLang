package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/",set)
	http.HandleFunc("/get",get)

	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func set(w http.ResponseWriter,r *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name:"my-cookie",
		Value:"value",
	})

	fmt.Fprintln(w,"cookie written!...ckeck your browser")
	fmt.Fprintln(w,"in chrome>dev tools>applications>cookies")

}

func get(w http.ResponseWriter,r *http.Request){
    c,err:=r.Cookie("my-cookie")
    if err!=nil{
    	http.Error(w,err.Error(),http.StatusNoContent)
	}

	fmt.Fprintln(w,"cookie is:",c)
}
