package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
)

type user struct{
	Uname string
	Fname string
	Lname string
}

var t *template.Template

var dbsession=map[string]string{}    //uuid userid
var dbuser=map[string]user{}       //map of uuid to user

func init(){
	t=template.Must(template.ParseFiles("session.gohtml","bar.gohtml"))
}

func main() {
	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter,r *http.Request){
	c,err:=r.Cookie("session1")
	if err!=nil{
		id:=uuid.NewV4()
		c=&http.Cookie{
			Name:"session1",
			Value:id.String(),
		}
		http.SetCookie(w,c)
	}

	var u1 user

	//if user is already there
	if un,ok:=dbsession[c.Value]; ok{
		u1=dbuser[un]
	}

    //if user is not already there create on from the form text fields and retrieve the (email,first,last) to create a user
	if r.Method==http.MethodPost{        //write this in else if eith above if
		un:=r.FormValue("email")
		f:=r.FormValue("first")
		l:=r.FormValue("last")
		u1=user{un,f,l}
		dbsession[c.Value]=un                          //assign uuid to its particular user
		dbuser[un]=u1
	}

	//pass that "already exixting" user or "newly created" user in template
	t.ExecuteTemplate(w,"session.gohtml",u1)

}

func bar(w http.ResponseWriter,r *http.Request){
	c,err:=r.Cookie("session1")
	if err!=nil{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	un,ok:=dbsession[c.Value]
	if !ok{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	u:=dbuser[un]

	t.ExecuteTemplate(w,"bar.gohtml",u)
}
