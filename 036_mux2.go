 package main

import (
	"net/http"
	"fmt"
)

type dog int
type cat int

func(d dog) ServeHTTP(w http.ResponseWriter,r *http.Request){      //this is the signature of handler interface
	//switch r.URL.Path {
	fmt.Fprintln(w,"barks")
	//case "/cat":fmt.Fprintln(w,"meows")
	//}
}

func(c cat) ServeHTTP(w http.ResponseWriter,r *http.Request){      //this is the signature of handler interface
	fmt.Fprintln(w,"mews")
}

func main() {
	var d dog
	var c cat

	//Handle need a handler
	http.Handle("/dog/",d)              //no servemux() is created it uses the default serve mux
	http.Handle("/cat",c)

	http.ListenAndServe(":8080",nil)
}
