package main

import "net/http"

func main() {
	http.Handle("/car.jpeg", http.FileServer(http.Dir("./")))
	http.Handle("favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)

}
