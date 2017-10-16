package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

var fm = template.FuncMap{
	"sq" : square,
	"div2" : divide,
}

func square(n float64) float64{
	return n*n
}

func divide(n float64) float64{
	return n/2
}

func init(){
	tpl=template.Must(template.New("new one").Funcs(fm).ParseFiles("pipeline.gohtml"))
	//tpl = template.Must(template.ParseFiles("fns.gohtml"))

}

func main() {

	err :=tpl.ExecuteTemplate(os.Stdout,"pipeline.gohtml",5.0)
	if err!=nil{
		log.Fatalln(err)
	}

}
