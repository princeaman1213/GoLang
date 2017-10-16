package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/",index)
	http.HandleFunc("/set",set)
	http.HandleFunc("/get",get)
	http.HandleFunc("/del",del)

	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func index(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","text/html; charset=utf-8")
	fmt.Fprintln(w,`<a href="/set">Set Cookie</a>`)
}

func set(w http.ResponseWriter,r *http.Request){
	http.SetCookie(w,&http.Cookie{
		Name:"Cookie-1",
		Value:"Val-1",
	})


	fmt.Fprintln(w,`<a href="/get">Get Cookie</a>`)

}

func get(w http.ResponseWriter,r *http.Request){
	c,err:=r.Cookie("Cookie-1")
	if err!=nil{
		http.Redirect(w,r,"/set",http.StatusSeeOther)  //if cookie not found go to set again
		//http.Error(w,err.Error(),http.StatusNoContent)
		return
	}
		//fmt.Fprintln(w,"Your Cookie is : ",c)


	fmt.Fprintln(w,`<h1>Your Cookie : %v </h1>`,c)
	fmt.Fprintln(w,`<br><a href="/del">Delete Cookie</a>`)

}

func del(w http.ResponseWriter,r *http.Request){
	c,err:=r.Cookie("Cookie-1")
	if err!=nil{
		http.Redirect(w,r,"/set",http.StatusSeeOther)    //if cookie not found go to set again
		//http.Error(w,err.Error(),http.StatusNoContent)
		return
	}else{
		c.MaxAge=-1
	}

	http.SetCookie(w,c)

	fmt.Fprintln(w,`<a href="/">Index (Home)</a>`)

}