package main

import (
	"net/http"
	"encoding/json"
)

type person struct{
	Name string
	Age int
	Hobies []string
}

func main() {

	http.HandleFunc("/",index)
	http.HandleFunc("/marshal",marshal)
	http.HandleFunc("/encode",encode)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func index(w http.ResponseWriter,r *http.Request){

	s:=`<!DOCTYPE HTML>
	<html lang="en">
	<head>
	<meta charset="utf-8">
	<title>INDEX</title>
	</head>
	<body>
	<h1> At Index</h1>
	</body>
	</html>`

	w.Write([]byte(s))

}

func marshal(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-type","application/json")
	p1:=person{
		Name:"Aman",
		Age:21,
		Hobies:[]string{"Athletics","Music"},
	}

	bs,err:=json.Marshal(p1)
	if err!=nil{
		http.Error(w,"no json data",http.StatusNoContent)
		return
	}

	w.Write(bs)
}

func encode(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-type","application/json")
	p1:=person{
		Name:"Aman",
		Age:21,
		Hobies:[]string{"Athletics","Music"},
	}

	err:=json.NewEncoder(w).Encode(p1)
	if err!=nil{
		http.Error(w,"no json data",http.StatusNoContent)
		return
	}

}