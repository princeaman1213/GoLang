package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type dog int

func (m dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   url.Values
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index1.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index1.gohtml"))
}

func main() {
	var d dog
	http.ListenAndServe(":8080", d)
}