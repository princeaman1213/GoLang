package main

import (
	"html/template"
	"net/http"
	"crypto/rsa"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	"log"
	//	jwtreq "github.com/dgrijalva/jwt-go/request"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io"
)

const (
	privKeyPath = "app.rsa"
	pubKeyPath = "app.rsa.pub"
)

//var VerifyKey, SignKey []byte
var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type UserProfile struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// UserClaims is a set of JWT claims that contain UserProfile.
type UserClaims struct {
	Profile UserProfile `json:"profile"`
	jwt.StandardClaims
}

var claims UserClaims

func initKeys(){
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)
	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err !=nil{
		fmt.Println("key not read")
		return
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err !=nil{
		fmt.Println("key not read")
		return
	}
}

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbuser = map[string]user{}      // user ID, user
var dbsession = map[string]string{} // session ID, user ID


var auth bool
var t *template.Template
var err error
var db *sql.DB

func init() {
	tpl = template.Must(template.ParseFiles("signup.gohtml","bar1.gohtml","bar2.gohtml","newlogin.gohtml"))
	t=template.Must(template.ParseFiles("insertapi.gohtml","getpersonapi.gohtml","getall.gohtml","updateapi.gohtml","getallapi.gohtml"))
	bs,_:=bcrypt.GenerateFromPassword([]byte("123"),bcrypt.MinCost)
	dbuser["alex@gmail.com"]=user{"alex@gmail.com",bs,"alex","cons"} //sample user
}

func main() {

	initKeys()
	//initialize sql
	db,err=sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/test1")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)

	//PUBLIC ENDPOINTS
	http.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	http.Handle("/resource/",http.HandlerFunc(ValidateTokenMiddleware))
	http.HandleFunc("/logout",Logout)

	//api funcs
	http.HandleFunc("/gotoindex",index)
	//http.HandleFunc("/getall",getall)
	http.HandleFunc("/getperson",getperson)      //get person by id
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/deleterec",deleterec)  //.Methods("DELETE") (how to use this along in browser and not postman)

	http.HandleFunc("/update",update)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Now listening...")
	http.ListenAndServe(":8000", nil)
}

/////////////ENDPOINT HANDLERS////////////
func ProtectedHandler(w http.ResponseWriter, r *http.Request){

	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		//http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("Current User is : \n",u.UserName,u.First,u.Last)
	auth=true
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/gotoindex", http.StatusSeeOther)   //cant login the second time if already logged in
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

		// username taken?
		if _, ok := dbuser[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		c := &http.Cookie{
			Name:  "session",
		}
		http.SetCookie(w, c)
		//dbsession[c.Value] = un

		// store user in dbuser
		/*bs,err:=bcrypt.GenerateFromPassword([]byte(p),bcrypt.MinCost)        //encrypt password (we need this password to later)
		if err!=nil{                                                         //(check for first time login)
			http.Error(w,"internal server error",http.StatusInternalServerError)
		}*/

		u = user{un, []byte(p), f, l}
		dbuser[un] = u

		fmt.Fprintln(w,"Signup successfull")

		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w,r,"/login",http.StatusSeeOther)
	}

	// if the user exists already, get user
	var u user
	if un, ok := dbsession[c.Value]; ok {
		u = dbuser[un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbsession[c.Value]
	_, ok := dbuser[un]
	return ok
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	for _,v:=range dbuser{
		fmt.Println("dbuser database\n",v.UserName,string(v.Password))
	}
	fmt.Println(len(dbuser))

	w.Header().Set("Content-Type","text/html; charset=utf-8")

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("email")
		p := r.FormValue("password")

		var flag int=0

		//validate user credentials
		for _,v:=range dbuser{
			if un == v.UserName {
				err:=bcrypt.CompareHashAndPassword([]byte(v.Password),[]byte(p))
				if err ==nil{
					flag=1
				}
			}
		}

		if flag!=1{
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}

		fmt.Println("now issuing token.....")

		//set claims
		claims = UserClaims{
			UserProfile{Name: un, Permissions: []string{"read","write","modify","abcd"}},
			jwt.StandardClaims{
				Issuer: "administrator007",     //"test-project"
				//ExpiresAt:,
			},
		}

		//create and sign the token
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		ss, err := token.SignedString(SignKey)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			log.Printf("err: %+v\n", err)
			return
		}
		// create session
		c,err:=r.Cookie("session")
		if err!=nil{
			c = &http.Cookie{
				Name:  "session",
				Value: ss,
				Expires:time.Now().Add(time.Second*300),
			}
		}else{
			c.Value=ss
		}

		http.SetCookie(w, c)
		dbsession[c.Value] = un

		log.Printf("issued token : %v\n", ss)

		go func(){                                 //delete session at expiry
			time.Sleep(time.Second*300)
			delete(dbsession,c.Value)
			auth=false
		}()
		http.Redirect(w,r,"/resource",http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "newlogin.gohtml", nil)
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(r){
		fmt.Fprintln(w,"Session Expired")
		return
	}
	c,_:=r.Cookie("session")
	req,err:=http.NewRequest("GET","/resource",nil)
	if err!=nil{
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Header().Add("Authorization","Bearer "+c.Value)
	fmt.Println("Bearer "+c.Value)

	req.Header=w.Header()
	fmt.Printf("tokennn type : %T\n",r.Header.Get("Authorization"))

	tokenstr:=req.Header.Get("Authorization")
	tokenslice:=strings.Split(tokenstr," ")
	fmt.Println(tokenslice,len(tokenslice))

	//validate token
	token, err := jwt.Parse(tokenslice[1],func(token *jwt.Token) (interface{}, error){
		return VerifyKey, nil
	})

	if err == nil {

		if token.Valid{
			ProtectedHandler(w, r)
			http.Redirect(w,r,"/gotoindex",http.StatusSeeOther)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}
}

func Logout(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	c,_:=r.Cookie("session")

	delete(dbsession,c.Value)
	c.Value=""
	c.MaxAge = -1
	http.SetCookie(w, c)
    auth=false

	fmt.Fprintln(w,"Your session has been logged out")
}

func index(w http.ResponseWriter,r *http.Request){
	//checksession(w,r)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	if auth==false{
		fmt.Fprintln(w,"Not Logged In")
		io.WriteString(w,`<a href="/login">Go to Login Page</a>`)
		return
	}
	//t.ExecuteTemplate(w,"getall.gohtml",nil)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`<a href="/getperson">SEARCH BY EMP ID</a><br>
	<a href="/insert">INSERT RECORD</a><br>
	<a href="/deleterec">DELETE RECORD</a><br>
	<a href="/update">UPDATE RECORD</a><br>`)
}

func getall(w http.ResponseWriter,r *http.Request){
	//checksession(w,r)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	if auth==false{
		fmt.Fprintln(w,"Not Logged In")
		io.WriteString(w,`<a href="/login">Go to Login Page</a>`)
		return
	}
	rows, err := db.Query(`SELECT * FROM table2;`)
	check(err)
	defer rows.Close()
	// data to be used in query
	var s,str1,str2,str3,str4,str5 string
	fmt.Fprintln(w,"RETRIEVED RECORDS:\n")

	var data []string

	// query
	for rows.Next() {
		err=rows.Scan(&str1,&str2,&str3,&str4,&str5)
		check(err)
		s+=str1+"  "+str2+"  "+str3+"  "+str4+"  "+str5+"\t"+"\n"
		data=append(data,s)

	}
     fmt.Fprintln(w,data)
	//t.ExecuteTemplate(w,"getallapi.gohtml",data)
	//w.Header().Set("Content-Type","text/html; charset=utf-8")
	//fmt.Fprintln(w,s+`<h1><a href="/">Go to index</a></h1>`)
}

func getperson(w http.ResponseWriter,r *http.Request){            //by EmpId
	//checksession(w,r)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	if auth==false{
		fmt.Fprintln(w,"Not Logged In")
		io.WriteString(w,`<a href="/login">Go to Login Page</a>`)
		return
	}
	v1:=r.FormValue("empid")
	//params:=mux.Vars(r)
	rows,err:=db.Query(`SELECT * FROM table2 WHERE EmpId = ?;`,v1)
	check(err)

	var s,str1,str2,str3,str4,str5 string

	for rows.Next(){
		s="Selected Record from customer Table :\n"
		err=rows.Scan(&str1,&str2,&str3,&str4,&str5)
		check(err)
		s+=str1+"  "+str2+"  "+str3+"  "+str4+"  "+str5+"\n"
	}

	t.ExecuteTemplate(w,"getpersonapi.gohtml",s)
	//fmt.Fprintln(w,s)
}

func update(w http.ResponseWriter,r *http.Request){            //by EmpId
	//checksession(w,r)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	if auth==false{
		fmt.Fprintln(w,"Not Logged In")
		io.WriteString(w,`<a href="/login">Go to Login Page</a>`)
		return
	}
	v1:=r.FormValue("empid")
	v2:=r.FormValue("name")
	/*if v1==""{
		fmt.Fprintln(w,"No Record to update")
		return
	}*/
	//params:=mux.Vars(r)
	_,err:=db.Exec(`UPDATE table2 SET Name = ? WHERE EmpId = ?;`,v2,v1)
	check(err)

	t.ExecuteTemplate(w,"updateapi.gohtml",nil)
	//fmt.Fprintln(w,s)
}

func insert(w http.ResponseWriter, r *http.Request) {
	//checksession(w,r)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	if auth==false{
		fmt.Fprintln(w,"Not Logged In")
		io.WriteString(w,`<a href="/login">Go to Login Page</a>`)
		return
	}

	v1:=r.FormValue("empid")
	v2:=r.FormValue("name")
	v3:=r.FormValue("mobile")
	v4:=r.FormValue("address")
	v5:=r.FormValue("position")
	fmt.Println(v1,v2,v3,v4,v5)               //checking
	s:="INSERTED RECORD"
	//params:=mux.Vars(r)

	rows1,err1:=db.Query(`SELECT * FROM table2 WHERE EmpId = ?;`,v1)
	check(err1)
	if rows1.Next(){
		fmt.Fprintln(w,"Duplicate EmpId")
		return
	}

	query:="INSERT INTO table2(EmpId,Name,Mobile,Address,Position) VALUES(?,?,?,?,?);"
	stmt,err:=db.Prepare(query)
	check(err)

	exe, err := stmt.Exec(v1,v2, v3,v4,v5)
	check(err)

	t.ExecuteTemplate(w,"insertapi.gohtml",nil)

	n, err := exe.RowsAffected()
	check(err)

	fmt.Fprintln(w,s,n)

}

func deleterec(w http.ResponseWriter,r *http.Request){
	//checksession(w,r)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	if auth==false{
		fmt.Fprintln(w,"Not Logged In")
		io.WriteString(w,`<a href="/login">Go to Login Page</a>`)
		return
	}

	v1:=r.FormValue("empid")

	_, err := db.Query(`DELETE FROM table2 WHERE EmpId=?;`,v1)
	check(err)
	//defer stmt.Close()

	//r1, err := stmt.Exec()
	//check(err)

	//n, err := r1.RowsAffected()
	//check(err)
	t.ExecuteTemplate(w,"getpersonapi.gohtml",nil)
	fmt.Fprintln(w, "DELETED RECORD")

}

func check(err error){
	if err!=nil{
		log.Fatal(err)
	}
}