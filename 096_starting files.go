package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
)

var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("96files.gohtml"))
}

func main() {
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8080",nil)

}

func index(w http.ResponseWriter,r *http.Request){
	c:=getcookie(w,r)
	t.ExecuteTemplate(w,"96files.gohtml",c.Value)
}

func getcookie(w http.ResponseWriter,r *http.Request) *http.Cookie{

	c,err:=r.Cookie("photo-cookie")
	if err!=nil{
		id:=uuid.NewV4()
		c=&http.Cookie{
			Name:"photo-cookie",
			Value:id.String(),
		}
		http.SetCookie(w,c)
	}

	return c

}

