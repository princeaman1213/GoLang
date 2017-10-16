package main

import (
	"net/http"

	"io"
)

func main() {
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8080",nil)

}

func foo(w http.ResponseWriter,r *http.Request){
	val:=r.FormValue("q")
	//fmt.Fprintln(w,"fooooooo"+val)
	io.WriteString(w,"search :"+val)
}

// In browser:  http://localhost:8080/?q=dog