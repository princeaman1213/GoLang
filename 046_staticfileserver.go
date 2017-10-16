package main

import "net/http"

func main() {
	_=http.ListenAndServe(":8080",http.FileServer(http.Dir("."))) //as it returns error and we are not using it

	//OR

	//http.ListenAndServe(":8080",http.FileServer(http.Dir(".")))
}
