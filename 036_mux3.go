package main

import (
	"net/http"
	"fmt"
)


func d(w http.ResponseWriter,r *http.Request){      //this is the signature of handler interface
	//switch r.URL.Path {
	fmt.Fprintln(w,"barks")
	//case "/cat":fmt.Fprintln(w,"meows")
	//}
}

func c(w http.ResponseWriter,r *http.Request){      //this is the signature of handler interface
	fmt.Fprintln(w,"mews")
}

func main() {
    //Handle func needs a function with the following signature -- (w http.ResponseWriter,r *http.Request)
	http.HandleFunc("/dog/",d)           //c,d are functions with the above signature so they implement the handler interface
	http.HandleFunc("/cat",c)

	http.ListenAndServe(":8080",nil)
}
