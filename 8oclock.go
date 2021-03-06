package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	"encoding/json"
	"strings"
	"io/ioutil"
)

const (
	privKeyPath = "app.rsa"
	pubKeyPath = "app.rsa.pub"
)

// User represents a client that can use the API.
type User struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	Profile  UserProfile
}

// UserProfile represents a public part of User information.
type UserProfile struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// UserClaims is a set of JWT claims that contain UserProfile.
type UserClaims struct {
	Profile UserProfile `json:"profile"`
	jwt.StandardClaims
}

/*var signingKey = []byte("signing-key")

func signingKeyFn(*jwt.Token) (interface{}, error) {
	return signingKey, nil
}*/

var VerifyKey, signingKey []byte
func initKeys(){
	var err error

	signingKey, err = ioutil.ReadFile(privKeyPath)
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

func signingKeyFn(*jwt.Token) (interface{}, error) {
	return VerifyKey, nil
}

var sampleUser = User{
	"alex",
	"12345",
	UserProfile{Name: "James Smith", Permissions: []string{"read","write","modify"}},
}

func main() {

	initKeys()

	http.HandleFunc("/login", login)
	http.HandleFunc("/read", read)

	log.Fatalln(http.ListenAndServe(":7777", nil))
}

func login(rw http.ResponseWriter, req *http.Request) {
	/*username := req.FormValue("username")
	password := req.FormValue("password")
*/var user User

	//decode request into UserCredentials struct
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(rw, "Error in request")
		return
	}

	fmt.Println("Username :",user.Username,"\nPassword :",user.Password)

	//validate user credentials
	if strings.ToLower(user.Username) != sampleUser.Username {    //"alex"
		if user.Password != sampleUser.Password {                 //"12345"
			rw.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(rw, "Invalid credentials")
			return
		}
	}
	//if username == sampleUser.Username && password == sampleUser.Password {
		claims := UserClaims{
			UserProfile{Name: "James Smith", Permissions: []string{"read","write","modify"}},
			jwt.StandardClaims{
				Issuer: "administrator007",     //"test-project"
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(err.Error()))
		log.Printf("err: %+v\n", err)
		return
	}

	rw.WriteHeader(200)
	rw.Write([]byte(ss))
	log.Printf("issued token : %v\n", ss)
	return


	rw.WriteHeader(401)
	return
}

func read(rw http.ResponseWriter, req *http.Request) {
	var claims UserClaims
	token, err := jwtreq.ParseFromRequestWithClaims(req, jwtreq.AuthorizationHeaderExtractor, &claims, signingKeyFn)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("Failed to parse token"))
		log.Println("Failed to parse token")
		return
	}

	if !token.Valid {
		rw.WriteHeader(401)
		rw.Write([]byte("Invalid token"))
		log.Println("Invalid token")
		return
	}

	rw.WriteHeader(200)
	claimsString := fmt.Sprintf("claims: %v", claims)
	rw.Write([]byte(claimsString))
	log.Println(claimsString)
}