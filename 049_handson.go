package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func escape(w http.ResponseWriter,r *http.Request){  //(lne1)    //this is the signature of handler interface
	fmt.Fprintln(w,"Drogon is Dead.......!")
}

func d(w http.ResponseWriter,r *http.Request){  //(lne1)    //this is the signature of handler interface
	t,err:=template.ParseFiles("drogon.gohtml")
	if err!=nil{
		http.Error(w,"File not found",404)
	}

	t.ExecuteTemplate(w,"drogon.gohtml",nil)

}

func img(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"aman.jpg")
}

func main() {

	http.HandleFunc("/",escape)
	http.HandleFunc("/dog",d)
	http.HandleFunc("/aman.jpg",img)
	http.ListenAndServe(":8080",nil)        //d here is handler so it is handled by (lne1)
}

//what is the difference in using get or post