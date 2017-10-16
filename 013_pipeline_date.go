package main

import (
	"html/template"
	"os"
	"log"
	"time"
)

var tpl *template.Template



type structure struct{
	Name string
	Age int
}

var fm = template.FuncMap{
	"dmy" : daymonyear,
}

func daymonyear(t time.Time) string{
	return t.Format("01-02-2006") // t.Format(time.Kitchen)
}

func init(){
	tpl=template.Must(template.New("new one").Funcs(fm).ParseFiles("passtime.gohtml"))
	//tpl = template.Must(template.ParseFiles("fns.gohtml"))

}

func main() {

	err :=tpl.ExecuteTemplate(os.Stdout,"passtime.gohtml",time.Now())
	if err!=nil{
		log.Fatalln(err)
	}

}
