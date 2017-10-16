package main

import (
	"net/http"
	"io"
)


func d(w http.ResponseWriter,r *http.Request){  //(lne1)    //this is the signature of handler interface
	//w.Header().Set("Key","from me")
	w.Header().Set("Content-type","text/html ; charset=utf-8")   //this defines the header info and is necessary to identify this as a html

	//can also use fmt.Fprintln below
	io.WriteString(w,`
	<img src="/aman.jpg">
	`)                                                                                                            //"aman.jpg"
	//http.Handle("/",http.FileServer(http.Dir(".")))
}


func main() {

	http.HandleFunc("/dog",d)
	http.Handle("/",http.FileServer(http.Dir(".")))   //it searches all files in current directory

	http.ListenAndServe(":8080",nil)        //d here is handler so it is handled by (lne1)
}

//what is the difference in using get or post