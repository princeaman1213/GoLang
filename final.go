package main

import (
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	"crypto/rsa"
	"time"
	"database/sql"
	"github.com/gorilla/mux"

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

func StartServer(){

	//PUBLIC ENDPOINTS
	http.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	http.Handle("/resource/", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	http.HandleFunc("/logout",Logout)


	//API
	//Connecting Database
	db,err = sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/sample_db")
	checkError(err)
	defer db.Close()
	err = db.Ping()
	checkError(err)

	//Routing URLs
	router := mux.NewRouter()
	router.HandleFunc("/admin",InsertData).Methods("POST")
	router.HandleFunc("/admin/{id}",FetchData).Methods("GET")
	router.HandleFunc("/admin",UpdateData).Methods("PUT")
	router.HandleFunc("/admin",DeleteData).Methods("DELETE")


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
	claims := UserClaims{
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


	// create session
	c,err:=r.Cookie("session")
	if err!=nil{
		c = &http.Cookie{
			Name:  "session",
			Value: ss,
			Expires:time.Now().Add(time.Second*30),
		}
	}else{
		c.Value=ss
	}

	http.SetCookie(w, c)

	go func(){                                 //delete session at expiry
		time.Sleep(time.Second*30)
		ss=""
//		auth=false
	}()

	w.WriteHeader(200)
	w.Write([]byte(ss))
	log.Printf("issued token : %v\n", ss)

}

//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {


	//validate token
	token, err := jwtreq.ParseFromRequest(r, jwtreq.AuthorizationHeaderExtractor,func(token *jwt.Token) (interface{}, error){
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

func Logout (w http.ResponseWriter,r *http.Request){

	c,err:=r.Cookie("Auth")
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.MaxAge=-1
	http.SetCookie(w,c)

}





func api(){

}


var db *sql.DB
var err error

type UserDetail struct{
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
	Address string `json:"address,omitempty"`
}

func InsertData(w http.ResponseWriter, req *http.Request){
	req.Header.Set("Content-Type","application/json")
	var user UserDetail
	_ = json.NewDecoder(req.Body).Decode(&user)
	fmt.Println(user)
	_,err:=db.Exec(`INSERT INTO emptable VALUES(?,?,?,?)`,user.ID,user.Name,user.Position,user.Address)
	checkError(err)
}

func DeleteData(w http.ResponseWriter, req *http.Request){
	req.Header.Set("Content-Type","application/json")
	var id string
	_ = json.NewDecoder(req.Body).Decode(&id)
	_,err:= db.Exec(`DELETE FROM emptable WHERE ID = ?`,id)
	checkError(err)

}

func FetchData(w http.ResponseWriter, req *http.Request){
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



