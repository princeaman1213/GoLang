package main

import (
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/codegangsta/negroni"
)


//RSA KEYS AND INITIALISATION


const (
	privKeyPath = "app.rsa"
	pubKeyPath = "app.rsa.pub"
)

var VerifyKey, SignKey []byte

var tokenString string
var err1 error
func initKeys(){
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}



//STRUCT DEFINITIONS


type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type User struct {
	ID			int 	`json:"id"`
	Name		string  `json:"name"`
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type Response struct {
	Data	string	`json:"data"`
}

type Token struct {
	Token 	string    `json:"token"`
}

type Claims struct {
	Username string    `json:"username"`
	jwt.StandardClaims
}

var claims Claims
//SERVER ENTRY POINT


func StartServer(){

	//PUBLIC ENDPOINTS
	http.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	http.Handle("/resource/", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))
	log.Println("Now listening...")
	http.ListenAndServe(":8000", nil)
}

func main() {

	initKeys()
	StartServer()
}


//////////////////////////////////////////


/////////////ENDPOINT HANDLERS////////////


/////////////////////////////////////////


func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials

	//decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	fmt.Println(user.Username, user.Password)

	//validate user credentials
	if strings.ToLower(user.Username) != "alex" {
		if user.Password != "12345" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}

//	exptoken:=time.Now().Add(time.Hour).Unix()
	expcookie:=time.Now().Add(time.Hour)

	claims=Claims{
		"Aman",
		jwt.StandardClaims{
			//ExpiresAt:exptoken,
			//Issuer:"localhost:8000",
		},
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	//signing the token
	tokenString,err:=token.SignedString(SignKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}

	//create a cookie
	cookie:=http.Cookie{Name:"Auth",Value:tokenString,Expires:expcookie,HttpOnly:true}
	http.SetCookie(w,&cookie)
	json, _ :=  json.Marshal(tokenString)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	//http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)   //307
	/*//create a token instance using the token string
	response := Token{tokenString}
	JsonResponse(response, w)*/

}



//AUTH TOKEN VALIDATION


// middleware to protect private pages
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//validate token
	token, err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error){
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return VerifyKey,nil
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


func ProtectedHandler(w http.ResponseWriter, r *http.Request){

	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)

}

//HELPER FUNCTIONS


func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err :=  json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
