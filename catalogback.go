package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"github.com/jinzhu/gorm"
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"strings"
)

type Car struct{
	Car string `json:car,omitempty`
	Urlsearch string
}

type Hatchback struct {
	HatchID      uint     `gorm:"primary_key"`
	Pre2URL string          `gorm:"-"`           //ignore this field while making table
	PreURL string          `gorm:"-"`           //ignore this field while making table
	Pre3URL []string          `gorm:"-"`           //ignore this field while making table
	CompanyName  string
	CarName  string
	VariantName  string
	S_ID          uint
	F_ID          uint
}

type Sedan struct {
	SedanID      uint     `gorm:"primary_key"`
	Pre2URL string          `gorm:"-"`           //ignore this field while making table
	PreURL string          `gorm:"-"`           //ignore this field while making table
	Pre3URL []string          `gorm:"-"`           //ignore this field while making table
	CompanyName  string
	CarName  string
	VariantName  string
	S_ID          uint
	F_ID          uint
}

type Spec struct {
	SpecsID      uint     `gorm:"primary_key"`
	Type string
	Length  string
	Width  string
	Height  string
	Wheelbase    string
	GroundClearance  string
	SeatingCapacity  string
	FuelTankCapacity	  string
	Engine_Type	  string
	AlternateFuel    string
	Cylinders  string
	TransmissionType  string
	SuspensionFront	  string
	SuspensionRear	  string
	BrakeType		  string
	Wheels	  string
}

type HatchSpec struct{
	//	PreURL string          `gorm:"-"`           //ignore this field while making table
	Hatchback
	Spec
	Feature
}

type SedanSpec struct{
	//	PreURL string          `gorm:"-"`           //ignore this field while making table
	Sedan
	Spec
	Feature
}

type Feature struct {
	FeatureID      uint     `gorm:"primary_key"`
	Type string
	Airbags  string
	SeatBeltWarning	  string
	AntiLockBrakingSystem	  string
	FourWheelDrive	    string
	DifferentialLock	  string
	CentralLocking	  string
	PowerOutlets12V		  string
	ParkingAssist		  string
	LowFuelLevelWarning	    string
	GPSNavigationSystem	  string
	WarrantyYears	  string
}

var db *gorm.DB
var err error
var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("searchcar.html","catindex.html","routevariant.html","typeofvec.html","routecompany.html","routecar.html","showdata.html"))
}

func main() {
	db,err=gorm.Open("mysql","root:password@/cat_db?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("connected")
	}

	//db.SingularTable(true)
	//db.DropTableIfExists(&Hatchback{},&Sedan{},Spec{},Feature{})
	//db.CreateTable(&Hatchback{},&Sedan{},Spec{},Feature{})

	router:=mux.NewRouter()
	router.HandleFunc("/",catindex)
	router.HandleFunc("/type",typeofvec)
	router.HandleFunc("/type/{type}",carcompany)
	router.HandleFunc("/type/{type}/{carname}",carname)
	router.HandleFunc("/type/{type}/{carname}/{car}",variantname)
	router.HandleFunc("/type/{type}/{carname}/{car}/{variantname}",showdata)
	router.HandleFunc("/searchcar",searchcar)


	router.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",router)
}

func catindex(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	t.ExecuteTemplate(w,"catindex.html",nil)
}

func typeofvec(w http.ResponseWriter,r *http.Request){
	t.ExecuteTemplate(w,"typeofvec.html",nil)
}

func carcompany(w http.ResponseWriter,r *http.Request){
	var route = mux.Vars(r)
	fmt.Println("car company      ",route)

	var s string
	s="/type"+"/"+route["type"]
	//pre="/type"
	var links []string
	links=append(links,route["type"])

	fmt.Println("carcompany url      :",s)

	if route["type"]=="hatchbacks" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT company_name").Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].PreURL = s
			h[i].Pre3URL=links                                          //path links
		}
		t.ExecuteTemplate(w,"routecompany.html",h)

	}else if route["type"]=="sedans"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT company_name").Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].PreURL = s
			h[i].Pre3URL=links                                          //path links
		}
		t.ExecuteTemplate(w,"routecompany.html",h)
	}
}

func carname(w http.ResponseWriter,r *http.Request){
	var route = mux.Vars(r)
	fmt.Println("carname           :",route["type"],route["carname"])

	var s string
	s="/type"+"/"+route["type"]+"/"+route["carname"]
	var links []string
	links=append(links,route["type"])
	links=append(links,route["carname"])

	fmt.Println("carname  url      :",s)

	if route["type"]=="hatchbacks" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT car_name").Where("company_name=?", route["carname"]).Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].PreURL = s
			h[i].Pre3URL=links                                          //path links
		}
		t.ExecuteTemplate(w, "routecar.html", h)
	}

	if route["type"]=="sedans" {
		var s1 []Sedan
		db.Debug().Table("sedan").Select("DISTINCT car_name").Where("company_name=?", route["carname"]).Find(&s1)
		for i:=0;i<len(s1);i++{
			s1[i].PreURL=s
			s1[i].Pre3URL=links                                          //path links
		}
		t.ExecuteTemplate(w, "routecar.html", s1)
	}
}

func variantname(w http.ResponseWriter,r *http.Request){
	var route = mux.Vars(r)
	fmt.Println("url is        ",route)

	var s string
	s="/type"+"/"+route["type"]+"/"+route["carname"]+"/"+route["car"]
    var links []string
	links=append(links,route["type"])
	links=append(links,route["carname"])
	links=append(links,route["car"])

	url:=strings.Split(s,"/")
	fmt.Println(url,len(url))
	fmt.Println(url[0]+url[1]+url[2]+url[3]+url[4])

	fmt.Println("variantname url      :",s)

	if route["type"]=="hatchbacks" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("company_name=? AND car_name=?", route["carname"],route["car"]).Find(&h)
		for i:=0;i<len(h);i++{
			h[i].PreURL = s
			h[i].Pre3URL=links                                          //path links
		}
		t.ExecuteTemplate(w, "routevariant.html", h)
	}

	if route["type"]=="sedans" {
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("company_name=? AND car_name=?", route["carname"], route["car"]).Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].PreURL = s
			h[i].Pre3URL=links                                         //path links
		}
		t.ExecuteTemplate(w, "routevariant.html", h)
	}
}

func showdata(w http.ResponseWriter,r *http.Request){
	var route = mux.Vars(r)
	fmt.Println("showdata url        ",route)

	if route["type"]=="hatchbacks" {
		var t1 HatchSpec
		db.Raw("SELECT spec.*, hatchback.variant_name FROM spec INNER JOIN hatchback ON hatchback.s_id = spec.specs_id Where hatchback.variant_name=? AND hatchback.car_name=?", route["variantname"],route["car"]).Scan(&t1)
		db.Raw("SELECT feature.*, hatchback.variant_name FROM feature INNER JOIN hatchback ON hatchback.f_id = feature.feature_id Where hatchback.variant_name=? AND hatchback.car_name=?", route["variantname"],route["car"]).Scan(&t1)

		t.ExecuteTemplate(w, "showdata.html", t1)
	}else if route["type"]=="sedans" {
		var t1 SedanSpec
		db.Raw("SELECT spec.*, sedan.variant_name FROM spec INNER JOIN sedan ON sedan.s_id = spec.specs_id Where sedan.variant_name=? AND sedan.car_name=?", route["variantname"], route["car"]).Scan(&t1)
		db.Raw("SELECT feature.*, sedan.variant_name FROM feature INNER JOIN sedan ON sedan.f_id = feature.feature_id Where sedan.variant_name=? AND sedan.car_name=?", route["variantname"], route["car"]).Scan(&t1)

		t.ExecuteTemplate(w, "showdata.html", t1)
	}
}

func searchcar(w http.ResponseWriter,r *http.Request) {
	var formcar string
	var data []string

	if r.Method == http.MethodPost {
		formcar = r.FormValue("car")
		var t string
		var h []Hatchback
		var sed []Sedan

		rows, err := db.Raw("select car_name from hatchback where company_name = ?", formcar).Rows()
        if err!=nil{
        	fmt.Println("Error at beginning..!")
		}
		defer rows.Close()

		rows1, err1 := db.Raw("select car_name from sedan where company_name = ?", formcar).Rows()
		if err1!=nil{
			fmt.Println("Error at 2 beginning..!")
		}
		defer rows1.Close()

        if rows.Next(){
			db.Raw("Select DISTINCT car_name FROM hatchback Where company_name=?", formcar).Scan(&h)
			//fmt.Println("1st if")
			fmt.Println("hatchback is      ;", h)
			t = "/type/hatchbacks/" + formcar
			data = append(data, t)
			for _, v := range h {
				data = append(data, v.CarName)
			}
			fmt.Println("data1 is :::::", data)
		}

		if rows1.Next() {
			db.Raw("Select DISTINCT car_name FROM sedan Where company_name=?", formcar).Scan(&sed)
			//fmt.Println("2nd if")
			fmt.Println("sedan is      ;", sed)
			t = "/type/sedans/" + formcar
			data = append(data, t)
			for _, v := range sed {
				data = append(data, v.CarName)
			}
			fmt.Println("data is :::::", data, len(data))
		}
	}
	t.ExecuteTemplate(w,"searchcar.html",data)
}