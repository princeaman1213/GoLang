package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/",set)
	http.HandleFunc("/get",get)
	http.HandleFunc("/multiple",multiple)

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

func multiple(w http.ResponseWriter,r *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name:"newcookie-1",
		Value:"value1",
	})

	http.SetCookie(w,&http.Cookie{
		Name:"newcookie-2",
		Value:"value2",
	})

	fmt.Fprintln(w,"cookie 1 & 2 are written!...ckeck your browser")
	fmt.Fprintln(w,"in chrome>dev tools>applications>cookies")

}

func get(w http.ResponseWriter,r *http.Request){
	c,err:=r.Cookie("my-cookie")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusNoContent)
	}else {
		fmt.Fprintln(w,"cookie is:",c)
	}

	c,err=r.Cookie("newcookie-1")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusNoContent)
	}else {
		fmt.Fprintln(w,"cookie is:",c)
	}

	c,err=r.Cookie("newcookie-2")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusNoContent)
	}else {
		fmt.Fprintln(w,"cookie is:",c)
	}
}
