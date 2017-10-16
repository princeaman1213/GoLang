package main

import (
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	"crypto/rsa"
	"go/token"
)

//RSA KEYS AND INITIALISATION
const (
	privKeyPath = "app.rsa"
	pubKeyPath = "app.rsa.pub"
)

//var VerifyKey, SignKey []byte
var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)
func initKeys(){
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
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

type UserProfile struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// UserClaims is a set of JWT claims that contain UserProfile.
type UserClaims struct {
	Profile UserProfile `json:"profile"`
	jwt.StandardClaims
}

//SERVER ENTRY POINT
var claims UserClaims
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

/////////////ENDPOINT HANDLERS////////////
func ProtectedHandler(w http.ResponseWriter, r *http.Request){

	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials
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

	//set claims
	claims = UserClaims{
		UserProfile{Name: "James Smith", Permissions: []string{"read","write","modify"}},
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

	http.SetCookie(w,&http.Cookie{
		Name:"Auth",
		Value:ss,
	})

	w.WriteHeader(200)
	w.Write([]byte(ss))
	log.Printf("issued token : %v\n", ss)

}

//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	_,err:=r.Cookie("Auth")
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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