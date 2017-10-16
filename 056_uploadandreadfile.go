package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
)

func main() {
	http.HandleFunc("/",f)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func f(w http.ResponseWriter,r *http.Request){
	var s string
	fmt.Println(r.Method)
	if r.Method==http.MethodPost {
		f, h, err := r.FormFile("q")        //for file
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile :", f, "\nheader :", h, "\nerr :", err)

		bs, err1 := ioutil.ReadAll(f)
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)
	}
	//the below code should not be in if block
		w.Header().Set("Content-type","text/html; charset=utf-8")
		io.WriteString(w,`
		<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="text" name="var">
		<input type="submit">
		</form><br>`+s)

		//fmt.Fprintln(w,"\n\n")
		fmt.Fprintln(w,r.FormValue("var"))

}


