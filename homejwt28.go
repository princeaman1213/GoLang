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
	"time"
	"strings"
	"golang.org/x/crypto/bcrypt"

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

func init() {
	tpl = template.Must(template.ParseFiles("signup.gohtml","bar1.gohtml","bar2.gohtml","login.gohtml"))
	bs,_:=bcrypt.GenerateFromPassword([]byte("123"),bcrypt.MinCost)
	dbuser["alex@gmail.com"]=user{"alex@gmail.com",bs,"alex","cons"} //sample user
}

func main() {

	initKeys()

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)

	//PUBLIC ENDPOINTS
	http.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	http.Handle("/resource/",http.HandlerFunc(ValidateTokenMiddleware))
	http.HandleFunc("/logout",Logout)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Now listening...")
	http.ListenAndServe(":8000", nil)
}

/////////////ENDPOINT HANDLERS////////////
func ProtectedHandler(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type","text/html; charset=utf-8")
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar2.gohtml", u)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "bar1.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)   //cant login the second time if already logged in
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
		//sID := uuid.NewV4()
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
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
		//return
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
	/*if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}*/
	for _,v:=range dbuser{
		fmt.Println("fdx  ",v.UserName,string(v.Password))
	}
	fmt.Println(len(dbuser))

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

		//http.Redirect(w, r, "/", http.StatusSeeOther)     //can also redirect to /resource now
		//return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(r){
		fmt.Fprintln(w,"Session Expired")
		return
	}
	//r.Method="POST"
	//fmt.Println(r.Method)
//var req *http.Request
	//sending req
	//bearer:="Bearer "

	c,_:=r.Cookie("session")
	req,err:=http.NewRequest("GET","/resource",nil)
	if err!=nil{
		return
	}
	req.Header.Set("Authorization","Bearer "+c.Value)

	tokenstr:=req.Header.Get("Authorization")
	tokenslice:=strings.Split(tokenstr," ")
	fmt.Println(tokenslice,len(tokenslice))
	//fmt.Println(jwtreq.AuthorizationHeaderExtractor)

	//fmt.Println(strings.Split(r.Header.Get("Authorization")," ")[1])
	//validate token
	token, err := jwt.ParseWithClaims(tokenslice[1],&claims,func(token *jwt.Token) (interface{}, error){
		return VerifyKey, nil
	})

	if err == nil {

		if token.Valid{
			ProtectedHandler(w, r)
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

	/*claims = UserClaims{
		UserProfile{Name: "", Permissions: []string{}},
		jwt.StandardClaims{
			ExpiresAt:time.Now().Unix(),     //"test-project"
		},
	}*/

	//io.WriteString(w,"You have been logged out !")
	// why does the above line gives this error(http: multiple response.WriteHeader calls)
	fmt.Fprintln(w,"Your session has been logged out")
	//http.Redirect(w, r, "/login", http.StatusSeeOther)
}