package main

import (
	"net/http"
	"fmt"
	"context"
	"time"
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

	res,err:=access(ctx)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusRequestTimeout)
		return
	}
	fmt.Fprintln(w,res)

}

func access(ctx context.Context) (int,error){

	ctx,cancel:=context.WithTimeout(ctx,time.Second*2)
	defer cancel()

	ch:=make(chan int)

	go func(){
		uid:=ctx.Value("uid").(int)
		time.Sleep(time.Second*10)

		if ctx.Err!=nil{
			return
		}
       ch <-uid
	}()

	 select{
	 case <-ctx.Done(): return 0,ctx.Err()
	 case c:=<-ch: return c,nil
	 }

}