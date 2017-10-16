package main

import (
	"net/http"
	"fmt"
	"strings"
)

func main() {

	//http.HandleFunc("/",index)
	http.HandleFunc("/",setheader)
	http.ListenAndServe(":8000",nil)

}

func index(w http.ResponseWriter,r *http.Request){
	setheader(w,r)
	var request []string

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	fmt.Println(request)

	fmt.Printf("tokennn type : %T\n",r.Header.Get("Authorization"))

	tokenstr:=r.Header.Get("Authorization")
	fmt.Printf("token is :%v",tokenstr)
	fmt.Println(len(tokenstr))

	tokenslice:=strings.Split(tokenstr," ")
	fmt.Println(tokenslice,len(tokenslice))

}

func setheader(w http.ResponseWriter,r *http.Request){
	//r.Method="POST"
//	w.Header().Set("Content-Type","application/json")
//	w.Header().Add("Authorization","Bearer "+"abcd")
	//fmt.Println("Bearer "+"abcd")

     req,err:=http.NewRequest("GET","/",nil)
     if err!=nil{
     	return
	 }
	client := &http.Client{}
	 req.Header.Set("Authorization","Bearer "+"abcd")
//url.ParseRequestURI("/")
	 fmt.Println(req.Header.Get("Authorization"))
	res, _ := client.Do(req)

	fmt.Fprintln(w,res)


}