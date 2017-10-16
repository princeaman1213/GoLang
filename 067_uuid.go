package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

func main() {

	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8080",nil)

}

func foo(w http.ResponseWriter,r *http.Request){
	c,err:=r.Cookie("session")
	if err!=nil{
		//http.Error(w,err.Error(),http.StatusNotFound)

	    id:=uuid.NewV4()

	    c=&http.Cookie{
	    	Name: "session",
	    	Value:id.String(),
	    	//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w,c)
	}

	fmt.Fprintln(w,"cookie is : ",c)

}

