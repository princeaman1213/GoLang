package main

import (
	"html/template"
	"os"
	"log"
	"strings"
)

var tpl *template.Template

type structure struct{
	Name string
	Age int
}

var fm = template.FuncMap{
	"trim" : first_3,
}

	func first_3(str string) string{
		str = strings.TrimSpace(str)
		str =str[:3]
		return str
}

func init(){
	tpl=template.Must(template.New("").Funcs(fm).ParseFiles("passfn.gohtml"))
	//tpl = template.Must(template.ParseFiles("fns.gohtml"))

}

func main() {

	st :=structure{Name: "Aman", Age:21}                          //composit type 3
	st1 :=structure{Name: "Ambuj", Age:23}

	slice_st :=[]structure{st,st1}

	err :=tpl.ExecuteTemplate(os.Stdout,"passfn.gohtml",slice_st)
	if err!=nil{
		log.Fatalln(err)
	}

}
