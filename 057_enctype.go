package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("enctype.gohtml"))
}

type detil struct {
	Fname string
	Lname string
	Sub bool
}

func main() {
	http.HandleFunc("/",f)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func f(w http.ResponseWriter,r *http.Request){
	bs:=make([]byte,r.ContentLength)
    //fmt.Println(r.ContentLength)
	r.Body.Read(bs)                                //reads the content from gohtml file
	bdy:=string(bs)

//	v:=r.FormValue("first")

	err:=t.ExecuteTemplate(w,"enctype.gohtml",bdy)
	if err!=nil{
		fmt.Println(err)
	}


	/*var s string
	fmt.Println(r.Method)
	if r.Method==http.MethodPost {
		f, h, err := r.FormFile("q")
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

	w.Header().Set("Content-type","text/html; charset=utf-8")
	fmt.Fprintln(w,s)*/

}

//try all 3 enctypes in thegohtml file for this program
