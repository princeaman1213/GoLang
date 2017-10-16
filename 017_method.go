package main

import (
	"html/template"
	"os"
	"fmt"
)

type per struct{
	Name string
	Age float64
}

var f *template.Template

func init(){
	f=template.Must(template.ParseFiles("passmethod.gohtml"))
}

func(p *per) Agedbl() float64{
	p.Age*=2
	return p.Age
}

func(p *per) Square(x float64) float64{
	return x*x
}

func main() {

	p1 :=per{"Aman",21}
	p2:=per{"Ambuj",22}
	p3:=per{"banga",24}

	p:=[]per{p1,p2,p3}

	err:=f.ExecuteTemplate(os.Stdout,"passmethod.gohtml",p)
	if err!=nil{
		fmt.Println(err)
	}


}
