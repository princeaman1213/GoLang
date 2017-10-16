package main

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	"fmt"
	"context"
)

type key int
const mykey key= 0

type Claims struct {
	Username string    `json:"username"`
	jwt.StandardClaims
}

func main() {
	http.HandleFunc("/settoken", settoken)
	http.HandleFunc("/profile", validate(protectedprofile))
	http.HandleFunc("/logout", validate(logout))
	http.ListenAndServe(":9000", nil)
}

func settoken(w http.ResponseWriter,r *http.Request){

	exptoken:=time.Now().Add(time.Hour).Unix()
	expcookie:=time.Now().Add(time.Hour)

	claims:=Claims{
		"Aman",
		jwt.StandardClaims{
			ExpiresAt:exptoken,
			Issuer:"localhost:9000",
		},
	}

	//Generating Token
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	//signing the token
	signedtoken,_:=token.SignedString([]byte("password"))

	//create a cookie
	cookie:=http.Cookie{Name:"Auth",Value:signedtoken,Expires:expcookie,HttpOnly:true}
	http.SetCookie(w,&cookie)

	http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)   //307

}
// middleware to protect private pages
func validate(pg http.HandlerFunc) http.HandlerFunc{

	a:=func(w http.ResponseWriter,r *http.Request) {
		//check if cookie is present
		c, err := r.Cookie("Auth")
		if err != nil {
			http.Error(w, "cookie not found", http.StatusNotFound)
			return
		}

		//verify whether the signing method and key used is matching with the original ones

		token, err := jwt.ParseWithClaims(c.Value,&Claims{},func(token *jwt.Token) (interface{}, error){
			if _,ok :=token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte("password"),nil
		})
		if err!=nil{
			http.NotFound(w, r)
			return
		}

		if claims,ok:=token.Claims.(*Claims); ok && token.Valid{
			ctx:=context.WithValue(r.Context(),mykey,*claims)
			pg(w,r.WithContext(ctx))
		}else {
			http.NotFound(w, r)
			return
		}

	}
	return a

}

func protectedprofile(w http.ResponseWriter,r *http.Request) {

	cl,ok:=r.Context().Value(mykey).(Claims)
	if !ok{
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w,"Hey %s . How are u !",cl.Username)
}

func logout (w http.ResponseWriter,r *http.Request){
	delcookie:=http.Cookie{Name:"Auth",Value:"none",Expires:time.Now()}
	http.SetCookie(w,&delcookie)
	return
}
