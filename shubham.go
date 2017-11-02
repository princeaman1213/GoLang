package main

import (
	"html/template"
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

var tpl *template.Template
var db *gorm.DB
var err error


func main() {
	Connect()
	defer db.Close()
	router:=mux.NewRouter()
	router.HandleFunc("/",index)
	router.HandleFunc("/types",types)
	router.HandleFunc("/{type}",listCompanies)
	router.HandleFunc("/{type}/{car}",listCars)
	router.HandleFunc("/{type}/{car}/{variant}",VariantDetails)
	router.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",router)
}



func init(){
	tpl = template.Must(template.ParseGlob("Templates/*.html"))
}



func index(w http.ResponseWriter, req *http.Request){

	tpl.ExecuteTemplate(w,"index.html",nil)
}


func types(w http.ResponseWriter, req *http.Request){
	tpl.ExecuteTemplate(w,"types.html",nil)
}


func listCompanies(w http.ResponseWriter, req *http.Request){

	var route = mux.Vars(req)
	if route["type"] == "hatch"{
		var h []HatchBack
		db.Debug().Table("hatch_backs").Select("DISTINCT company_name").Find(&h)
		var hNames = []string{}
		for _,v:=range h{
			hNames = append(hNames,v.CompanyName)
		}

		tpl.ExecuteTemplate(w, "CompaniesListHatch.html", hNames)
	}
	if route["type"] == "sedan"{
		var s []Sedan
		db.Debug().Table("sedans").Select("DISTINCT company_name").Find(&s)
		var sNames = []string{}
		for _,v:=range s{
			sNames = append(sNames,v.CompanyName)
		}

		tpl.ExecuteTemplate(w, "CompaniesListSedan.html", sNames)
	}
}
func listCars(w http.ResponseWriter, req *http.Request){

	var c = mux.Vars(req)
	if c["type"]=="hatch" {
		var hCars []HatchBack
		if c["car"] == "Maruti" {
			db.Debug().Table("hatch_backs").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "Hyundai" {
			db.Debug().Table("hatch_backs").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "Tata" {
			db.Debug().Table("hatch_backs").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "BMW" {
			db.Debug().Table("hatch_backs").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "Mercedes" {
			db.Debug().Table("hatch_backs").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
	}



	if c["type"]=="sedan" {
		var hCars []HatchBack
		if c["car"] == "Maruti" {
			db.Debug().Table("sedans").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "Hyundai" {
			db.Debug().Table("sedans").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "Honda" {
			db.Debug().Table("sedans").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "BMW" {
			db.Debug().Table("sedans").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
		if c["car"] == "Audi" {
			db.Debug().Table("sedans").Select("DISTINCT car_name").Where("company_name=?", c["car"]).Find(&hCars)
			var carNames= []string{}
			for _, v := range hCars {
				carNames = append(carNames, v.CarName)
			}

			tpl.ExecuteTemplate(w, "ListCarsHatchComp.html", carNames)
		}
	}
}

func Connect() {

	db,err = gorm.Open("mysql","root:password@tcp(127.0.0.1:3306)/product_catalog?charset=utf8" +
		"&parseTime=True&loc=Local")
	if err!=nil {
		log.Fatal(err)
	} else{
		fmt.Println("connected")
	}


	err = db.DB().Ping()
	if err!=nil{
		log.Fatal(err)
	}

	/*db.DropTableIfExists(&Specs{},&Features{})
	db.CreateTable(&Specs{},&Features{})*/

}
