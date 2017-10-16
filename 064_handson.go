package main

import (
	"net/http"
	"fmt"
	"strconv"
)

func main() {

	http.HandleFunc("/",set)

	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}

func set(w http.ResponseWriter,r *http.Request){
	cookie,err:=r.Cookie("my-cookie")

	if err==http.ErrNoCookie{
		cookie=&http.Cookie{
			Name:"cook",
			Value:"0",
		}
	}

	c,_:=strconv.Atoi(cookie.Value)
	c++
	cookie.Value=strconv.Itoa(c)

	http.SetCookie(w,cookie)

	fmt.Fprintln(w,cookie.Value)
}
