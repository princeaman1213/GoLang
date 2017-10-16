package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"strings"
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
	c=appenddata(w,c)
	sl:=strings.Split(c.Value,"|")
	t.ExecuteTemplate(w,"96files.gohtml",sl)
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

func appenddata(w http.ResponseWriter,c *http.Cookie) *http.Cookie{
	p1:="img1.jpg"
	p2:="img2.jpg"
	p3:="img3.jpg"

	if !strings.Contains(c.Value,p1){
		c.Value+="|"+p1
	}
	if !strings.Contains(c.Value,p2){
		c.Value+="|"+p2
	}
	if !strings.Contains(c.Value,p3){
		c.Value+="|"+p3
	}

	http.SetCookie(w,c)
    return c
}

