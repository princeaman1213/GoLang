package main

import (
	"net/http"
	"fmt"
)

type dog int

func(d dog) ServeHTTP(w http.ResponseWriter,r *http.Request){      //this is the signature of handler interface
	fmt.Fprintln(w,"Any code")
}

func main() {
	var d dog
	http.ListenAndServe(":8080",d)
}
