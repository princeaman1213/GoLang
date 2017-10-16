package main

import (
	"net/http"
	"fmt"
	"context"
)

func main() {

	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter,r *http.Request){

	ctx:=r.Context()
	ctx=context.WithValue(ctx,"uid",1234)
	ctx=context.WithValue(ctx,"uname","aman")    //we can some values in context

	res:=access(ctx)
    fmt.Fprintln(w,res)

}

func access(ctx context.Context) int{
	uid:=ctx.Value("uid").(int)
	return uid
}