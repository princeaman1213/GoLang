package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
)


func d(w http.ResponseWriter,r *http.Request){  //(lne1)    //this is the signature of handler interface
	//w.Header().Set("Key","from me")
	w.Header().Set("Content-type","text/html ; charset=utf-8")
	io.WriteString(w,`
	<img src="/aman.jpg">
	`)                         //calls this signature in main which calls the img function            //"aman.jpg"
	fmt.Fprintln(w,"<h1>Any code here</h1>")
}

func img(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","text/html ; charset=utf-8")
	f,err:=os.Open("aman.jpg")
	if err!=nil{
		fmt.Println(err)
	}
	defer f.Close()
	io.Copy(w,f)
}

func main() {

	http.HandleFunc("/dog",d)
	http.HandleFunc("/aman.jpg",img)
	http.ListenAndServe(":8080",nil)        //d here is handler so it is handled by (lne1)
}

//what is the difference in using get or post