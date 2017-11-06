package main

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"

)

var t *template.Template

func init(){

	t=template.Must(template.ParseFiles("catindex.html"))

}

func main() {
	router:=mux.NewRouter()

	router.Handle("/background.jpg",http.StripPrefix("/",http.FileServer(http.Dir("./"))))
	router.HandleFunc("/index",catindex)


	router.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8000",router)
}

func catindex(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf-8")

	t.ExecuteTemplate(w,"catindex.html",nil)
}