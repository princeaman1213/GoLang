package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter,r *http.Request){

	ctx:=r.Context()
	fmt.Fprintln(w,ctx)

}
