package main

import (
	"html/template"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/satori/go.uuid"
	"time"
	"fmt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	UserName string
	lastact time.Time
}

var tpl *template.Template

var dbuser = map[string]user{}      // user ID, user
var dbsession = map[string]session{} // session ID, session

var dbsessioncleaned time.Time      //takes the value of time at the start of the program
const sessionlen int = 30

func init() {
	tpl = template.Must(template.ParseFiles("signupperm.gohtml","bar1.gohtml","bar2.gohtml","login.gohtml"))
	//bs,_:=bcrypt.GenerateFromPassword([]byte("123"),bcrypt.MinCost)
	//dbuser["test1@gmail.com"]=user{"test1@gmail.com",bs,"Aman","Patel","admin"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

	//go cleansession()
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	fmt.Println("fron index")
	showsession() //for understanding
	tpl.ExecuteTemplate(w, "bar1.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(w,r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if u.Role!="bond"{
		http.Error(w,"Entry not allowed , you are not BOND!!",http.StatusForbidden)
		return
	}
//showsession() //for demo
	tpl.ExecuteTemplate(w, "bar2.gohtml", u)

}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w,r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)   //cant login the second time
		return
	}

	var u user
	// process form submission
	if r.Method == http.MethodPost {

		// get form values
		un := r.FormValue("email")
		p := r.FormValue("password")
		f := r.FormValue("first")
		l := r.FormValue("last")
		r1 := r.FormValue("role")

		// username taken?
		if _, ok := dbuser[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbsession[c.Value] = session{un ,time.Now()}

		// store user in dbuser
		bs,err:=bcrypt.GenerateFromPassword([]byte(p),bcrypt.MinCost)        //encrypt password
		if err!=nil{
			http.Error(w,"internal server error",http.StatusInternalServerError)
		}

		u = user{un, bs, f, l,r1}
		dbuser[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("from sign up")
     showsession() //demo
	tpl.ExecuteTemplate(w, "signupperm.gohtml", u)
}

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge=sessionlen
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if un, ok := dbsession[c.Value]; ok {
		un.lastact=time.Now()
		dbsession[c.Value]=un
		u = dbuser[un.UserName]
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter,r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	s,ok:=dbsession[c.Value]
	if ok{
		s.lastact=time.Now()
		dbsession[c.Value]=s
	}

	_, ok = dbuser[s.UserName]
	c.MaxAge=sessionlen
	http.SetCookie(w,c)
	return ok
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w,r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("email")
		p := r.FormValue("password")
		// is there a username?
		u, ok := dbuser[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match?!", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge=sessionlen
		http.SetCookie(w, c)
		dbsession[c.Value] = session{un ,time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("from login")
    showsession()
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(w,r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c,_:=r.Cookie("session")

	delete(dbsession,c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)

	/*if time.Now().Sub(dbsessioncleaned) > (time.Second*30){ //if time now - time at start of prog is greater than 30 sec then clean session
		go cleansession()
	}*/


	//io.WriteString(w,"You have been logged out !")
	// why does the above line gives this error(http: multiple response.WriteHeader calls)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func cleansession(){
	fmt.Println("before clean")
    showsession()
	for i,v:=range dbsession{
		if time.Now().Sub(v.lastact)>(time.Second*30){
			delete(dbsession,i)
		}
	}
	dbsessioncleaned=time.Now()        //Why ??
	fmt.Println("after clean")
	showsession()
}

func showsession(){
	for i,v:=range dbsession{
		fmt.Println("\n",i,v.UserName)

	}
}