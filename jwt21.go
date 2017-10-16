package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
	"log"
)

type User struct{
	Username string   `json:"username"`
	Password string   `json:"password"`
	Prof Profile
}

type Profile struct{
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type Claims struct{
	Prof Profile `json:"profile"`
	jwt.StandardClaims
}

var claims Claims
var SignKey = []byte("password")

func signingKeyFn(*jwt.Token) (interface{}, error) {
	return SignKey, nil
}

var SampleUser = User{
	"alex",
	"12345",
	Profile{Name: "Aman", Permissions: []string{"read","write","modify"}},
}

func main() {

	http.HandleFunc("/login",login)
	http.HandleFunc("/getinfo",getinfo)
	http.ListenAndServe(":8080",nil)
}

func login(w http.ResponseWriter,r *http.Request){

	var user User
	err:=json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	fmt.Println("Username :",user.Username,"\nPassword :",user.Password)

	if strings.ToLower(user.Username) != SampleUser.Username{
		if strings.ToLower(user.Password) != SampleUser.Password{
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}

	claims=Claims{
		Profile{Name:"Aman",Permissions:[]string{"read","write","modify"}},
		jwt.StandardClaims{
			Issuer: "administrator007",
		},
	}


	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	signedtoken,err:=token.SignedString(SignKey)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("err: %+v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(signedtoken))
	log.Printf("issued token : %v\n", signedtoken)
	return

}

func getinfo(w http.ResponseWriter,r *http.Request){

	token,err:=jwtreq.ParseFromRequestWithClaims(r,jwtreq.AuthorizationHeaderExtractor,&claims,signingKeyFn)
    if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Required Token not found !"))
		log.Println("Required Token not found !")
		return
	}

	if !token.Valid {
		w.WriteHeader(401)
		w.Write([]byte("Invalid token"))
		log.Println("Invalid token")
		return
	}

	w.WriteHeader(200)
	claimsString := fmt.Sprintf("claims: %v", claims)
	w.Write([]byte(claimsString))
	log.Println(claimsString)

}
