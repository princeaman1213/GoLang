
package main

import (
	"html/template"

	"net/http"
	//"github.com/satori/go.uuid"
	"crypto/rsa"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/codegangsta/negroni"
	"log"
	jwtreq "github.com/dgrijalva/jwt-go/request"

	"time"
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
	//,_:=bcrypt.GenerateFromPassword([]byte("123"),bcrypt.MinCost)
	dbuser["alex@gmail.com"]=user{"alex@gmail.com",[]byte("123"),"alex","cons"} //sample user
}


/*func StartServer(){

	//PUBLIC ENDPOINTS
	http.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	http.Handle("/resource/", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	http.HandleFunc("/logout",Logout)

}*/

func main() {

	initKeys()

	http.HandleFunc("/", index)
	//	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	//http.HandleFunc("/login", login)
	//http.HandleFunc("/logout", logout)

	//PUBLIC ENDPOINTS
	http.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	http.Handle("/resource/", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	http.HandleFunc("/logout",Logout)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Now listening...")
	http.ListenAndServe(":8000", nil)
	//http.ListenAndServe(":8080", nil)
}

/////////////ENDPOINT HANDLERS////////////
func ProtectedHandler(w http.ResponseWriter, r *http.Request){
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar2.gohtml", u)
	/*response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)*/
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "bar1.gohtml", u)
}

/*func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar2.gohtml", u)
}*/

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
			//Value: sID.String(),
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
		/*for i,r:=range dbuser{
			fmt.Println(i,r.UserName,string(r.Password))
		}*/

		return
		// redirect
		//http.Redirect(w, r, "/login", http.StatusSeeOther)
		//return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		/*sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}*/
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
			if un == v.UserName && p == string(v.Password){
				//fmt.Println(user.Username,"\n",v.UserName,"\n",user.Password,"\n",string(v.Password),"\n")
				/*w.WriteHeader(http.StatusForbidden)
				fmt.Println("Error logging in")
				fmt.Fprint(w, "Invalid credentials")
				return*/
				flag=1
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
			UserProfile{Name: un, Permissions: []string{"read","write","modify"}},
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

		//w.WriteHeader(200)
		//w.Write([]byte(ss))
		log.Printf("issued token : %v\n", ss)

		//http.Redirect(w, r, "/", http.StatusSeeOther)     //can also redirect to /resource now
		//return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}


//AUTH TOKEN VALIDATION
/*func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if !alreadyLoggedIn(r){
		fmt.Fprintln(w,"Session Expired")
		return
	}
	fmt.Println(r.Method)

	//sending req
	bearer:="Bearer "
	c,err:=r.Cookie("session")
	if err==nil{
		bearer+=c.Value
		w.Header().Set("Content-Type","application/json")
		w.Header().Add("Authorization",bearer)
		fmt.Println(bearer)
		*//*http.Redirect(w,r,"/resource",http.StatusSeeOther)
		return*//*
	}


	//validate token at server
	token, err := jwtreq.ParseFromRequestWithClaims(r, jwtreq.AuthorizationHeaderExtractor,&claims,func(token *jwt.Token) (interface{}, error){
		return VerifyKey, nil
	})

	if err == nil {

		if token.Valid{
			var token1 string
			// Get token from the Authorization header
			// format: Authorization: Bearer
			tokens, ok := r.Header["Authorization"]
			if ok && len(tokens) >= 1 {
				token1 = tokens[0]
				token1 = strings.TrimPrefix(token1, "Bearer ")
				fmt.Println(token1)
				if token1 != c.Value{
					fmt.Println("Wrong Token sent for authorization")
					return
				}
			}
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}

}*/

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if !alreadyLoggedIn(r){
		fmt.Fprintln(w,"Session Expired")
		return
	}
	fmt.Println(r.Method)

	//sending req
	bearer:="Bearer "
	c,err:=r.Cookie("session")
	if err==nil{
		bearer+=c.Value
		w.Header().Set("Content-Type","application/json")
		w.Header().Set("Authorization",bearer)
		fmt.Println(bearer)
		/*http.Redirect(w,r,"/resource",http.StatusSeeOther)
		return*/
	}

	fmt.Println(jwtreq.AuthorizationHeaderExtractor)
    fmt.Println(r.Header.Get("Authorization"))
	//validate token
	token, err := jwtreq.ParseFromRequestWithClaims(r, jwtreq.AuthorizationHeaderExtractor,&claims,func(token *jwt.Token) (interface{}, error){
		return VerifyKey, nil
	})

	if err == nil {

		if token.Valid{
			next(w, r)
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