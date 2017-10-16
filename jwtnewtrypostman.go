
package main

import (
	"html/template"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	//"github.com/satori/go.uuid"
	"crypto/rsa"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	//"github.com/codegangsta/negroni"
	"log"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"time"
	"strings"
	"database/sql"
	"encoding/json"
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

//for sql database
type UserDetail struct{
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
	Address string `json:"address,omitempty"`
}

var db *sql.DB
var err error
var tpl *template.Template
var dbuser = map[string]user{}      // user ID, user
var dbsession = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseFiles("signup.gohtml","bar1.gohtml","bar2.gohtml","login.gohtml"))
	//bs,_:=bcrypt.GenerateFromPassword([]byte("123"),bcrypt.MinCost)
	dbuser["alex@gmail.com"]=user{"alex@gmail.com",[]byte("123"),"alex","cons"} //sample user
}

var auth bool

func main() {

	initKeys()
	//Connecting Database
	db,err = sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/test1")
	checkError(err)
	defer db.Close()
	err = db.Ping()
	checkError(err)

	router := mux.NewRouter()
	//http.HandleFunc("/", index)
	router.HandleFunc("/signup", signup)

	//PUBLIC ENDPOINTS
	router.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	router.HandleFunc("/resource", ValidateTokenMiddleware)

	router.HandleFunc("/logout",Logout)

	//Routing URLs

	router.HandleFunc("/admin",InsertData).Methods("POST")
	router.HandleFunc("/admin/{id}",FetchData).Methods("GET")
	router.HandleFunc("/admin",UpdateData).Methods("PUT")
	router.HandleFunc("/admin",DeleteData).Methods("DELETE")

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Now listening...")
	http.ListenAndServe(":8000", router)
}

/////////////ENDPOINT HANDLERS////////////
func ProtectedHandler(w http.ResponseWriter, r *http.Request){
	//u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Fprintln(w,"No Access ! Login to continue !")
		return
	}
	fmt.Fprintln(w,"Access to DB Granted !")
    auth =true
	//tpl.ExecuteTemplate(w, "bar2.gohtml", u)
}

/*func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "bar1.gohtml", u)
}*/

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)   //cant login the second time if already logged in
		return
	}

	// process form submission
     var u user
		// get form values
	err:=r.ParseForm()
	if err!=nil{
		fmt.Println("unable to parse form data !")
	}
	name1:=r.PostForm.Get("username")
	pwd1:=r.PostForm.Get("password")
	first:=r.PostForm.Get("first")
	last:=r.PostForm.Get("last")
	fmt.Println(name1,pwd1)

		// username taken?
		if _, ok := dbuser[name1]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		/*c := &http.Cookie{
			Name:  "session",
		}
		http.SetCookie(w, c)*/
		//dbsession[c.Value] = un

		// store user in dbuser
		/*bs,err:=bcrypt.GenerateFromPassword([]byte(pwd1),bcrypt.MinCost)        //encrypt password (we need this password to later)
		if err!=nil{                                                         //(check for first time login)
			http.Error(w,"internal server error",http.StatusInternalServerError)
		}*/

		u = user{name1, []byte(pwd1), first, last}
		dbuser[name1] = u

		fmt.Fprintln(w,"Signup successfull")
		//return

}

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w,r,"/login",http.StatusSeeOther)
	}
	//http.SetCookie(w, c)

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

//for testing purpose (login)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	for _,r:=range dbuser{
		fmt.Println(r.UserName,string(r.Password))
	}

	fmt.Println("size is :",len(dbuser))
	/*var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}*/

	err:=r.ParseForm()
	if err!=nil{
		fmt.Println("unable to parse form data !")
	}
    name1:=r.PostForm.Get("username")
    pwd1:=r.PostForm.Get("password")
    fmt.Println(name1,pwd1)
	//fmt.Println(user.Username, user.Password)
	var flag int=0
	//validate user credentials
	for _,v:=range dbuser{
		if strings.ToLower(name1) == v.UserName && string(v.Password) == pwd1 {
				flag=1
			}
		}

	if flag!=1{
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	//set claims
	claims = UserClaims{
		UserProfile{Name: name1, Permissions: []string{"read","write","modify"}},
		jwt.StandardClaims{
			Issuer: "administrator007",     //"test-project"
		},
	}

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
	dbsession[c.Value] = name1

	go func(){
		time.Sleep(time.Second*300)
		delete(dbsession,c.Value)
	}()

	w.WriteHeader(200)
	w.Write([]byte(ss))
	log.Println("Bearer "+ss)

}

//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(r){
		fmt.Fprintln(w,"Session Expired ! Login again")
		return
	}
	fmt.Println(r.Method)

	//validate token
	token, err := jwtreq.ParseFromRequestWithClaims(r, jwtreq.AuthorizationHeaderExtractor,&claims,func(token *jwt.Token) (interface{}, error){
		return VerifyKey, nil
	})

	if err == nil {

		if token.Valid{
			ProtectedHandler(w,r)
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c,_:=r.Cookie("session")

	delete(dbsession,c.Value)
	c.Value=""
	c.MaxAge = -1
	http.SetCookie(w, c)

	auth=false

	//io.WriteString(w,"You have been logged out !")
	// why does the above line gives this error(http: multiple response.WriteHeader calls)
	fmt.Fprintln(w,"Your session has been logged out")
	//http.Redirect(w, r, "/login", http.StatusSeeOther)
}

//SQL Functions


func InsertData(w http.ResponseWriter, req *http.Request){
    if auth!=true{
    	fmt.Fprintln(w,"Login to proceed !!")
    	return
	}
	req.Header.Set("Content-Type","application/json")
	var user UserDetail
	_ = json.NewDecoder(req.Body).Decode(&user)
	fmt.Println(user)
	_,err:=db.Exec(`INSERT INTO emptable VALUES(?,?,?,?)`,user.ID,user.Name,user.Position,user.Address)
	checkError(err)
}



func DeleteData(w http.ResponseWriter, req *http.Request){
	if auth!=true{
		fmt.Fprintln(w,"Login to proceed !!")
		return
	}
	req.Header.Set("Content-Type","application/json")
	var id string
	_ = json.NewDecoder(req.Body).Decode(&id)
	_,err:= db.Exec(`DELETE FROM emptable WHERE ID = ?`,id)
	checkError(err)

}


func FetchData(w http.ResponseWriter, req *http.Request){
	if auth!=true{
		fmt.Fprintln(w,"Login to proceed !!")
		return
	}
	req.Header.Set("Content-Type","application/json")
	var user UserDetail
	params:=mux.Vars(req)
	res,err:= db.Query(`SELECT * FROM emptable WHERE ID = ? `,params["id"])
	checkError(err)
	for res.Next(){
		err:= res.Scan(&user.ID,&user.Name,&user.Position,&user.Address)
		checkError(err)
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateData(w http.ResponseWriter, req *http.Request){
	if auth!=true{
		fmt.Fprintln(w,"Login to proceed !!")
		return
	}
	req.Header.Set("Content-Type","application/json")
	var user UserDetail
	_ = json.NewDecoder(req.Body).Decode(&user)

	_,err:= db.Exec(`UPDATE emptable SET NAME = ?, POSITION = ?, ADDRESS = ? WHERE ID = ?`,user.Name,user.Position,user.Address,user.ID)
	checkError(err)

}

func checkError(e error){
	if e!=nil{
		log.Fatalln(e)
	}
}
