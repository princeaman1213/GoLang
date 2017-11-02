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
)


type Hatchback struct {
	HatchID      uint     `gorm:"primary_key"`
	Pre2URL string          `gorm:"-"`           //ignore this field while making table
	PreURL string          `gorm:"-"`           //ignore this field while making table
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
	t=template.Must(template.ParseFiles("catindex.html","routevariant.html","typeofvec.html","routecompany.html","routecar.html","showdata.html"))
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
	router.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8000",router)
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

	var s,pre string
	s="/type"+"/"+route["type"]
	pre="/type"
	fmt.Println("temp url      :",s)

	if route["type"]=="hatchbacks" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT company_name").Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].Pre2URL=pre
			h[i].PreURL = s
		}
		t.ExecuteTemplate(w,"routecompany.html",h)

	}else if route["type"]=="sedans"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT company_name").Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].Pre2URL=pre
			h[i].PreURL = s
		}
		t.ExecuteTemplate(w,"routecompany.html",h)
	}
}

func carname(w http.ResponseWriter,r *http.Request){
	var route = mux.Vars(r)
    fmt.Println("carname           :",route["type"],route["carname"])

    var s,pre string
    s="/type"+"/"+route["type"]+"/"+route["carname"]
    pre="/type"+"/"+route["type"]
    fmt.Println("tempcar name  url      :",s)

    if route["type"]=="hatchbacks" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT car_name").Where("company_name=?", route["carname"]).Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].PreURL = s
			h[i].Pre2URL = pre
		}
		t.ExecuteTemplate(w, "routecar.html", h)
	}

	if route["type"]=="sedans" {
			var s1 []Sedan
			db.Debug().Table("sedan").Select("DISTINCT car_name").Where("company_name=?", route["carname"]).Find(&s1)
			for i:=0;i<len(s1);i++{
				s1[i].PreURL=s
				s1[i].Pre2URL = pre
			}
			t.ExecuteTemplate(w, "routecar.html", s1)
		}
}

func variantname(w http.ResponseWriter,r *http.Request){
	var route = mux.Vars(r)
	fmt.Println("url is        ",route)

	var s,pre string
	s="/type"+"/"+route["type"]+"/"+route["carname"]+"/"+route["car"]
	pre="/type"+"/"+route["type"]+"/"+route["carname"]
	fmt.Println("tempera url      :",s)

	if route["type"]=="hatchbacks" {
			var h []Hatchback
				db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("company_name=? AND car_name=?", route["carname"],route["car"]).Find(&h)
				for i:=0;i<len(h);i++{
					h[i].PreURL=s
					h[i].Pre2URL = pre
				}
				t.ExecuteTemplate(w, "routevariant.html", h)
		}

	if route["type"]=="sedans" {
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("company_name=? AND car_name=?", route["carname"], route["car"]).Find(&h)
		for i := 0; i < len(h); i++ {
			h[i].PreURL = s
			h[i].Pre2URL = pre
		}
		t.ExecuteTemplate(w, "routevariant.html", h)
	}
}

func showdata(w http.ResponseWriter,r *http.Request){
	var route = mux.Vars(r)
	fmt.Println("urllastfn is        ",route)

	if route["type"]=="hatchbacks" {
		var t1 HatchSpec
		db.Raw("SELECT spec.*, hatchback.variant_name FROM spec INNER JOIN hatchback ON hatchback.s_id = spec.specs_id Where hatchback.variant_name=? AND hatchback.car_name=?", route["variantname"],route["car"]).Scan(&t1)
		db.Raw("SELECT feature.*, hatchback.variant_name FROM feature INNER JOIN hatchback ON hatchback.f_id = feature.feature_id Where hatchback.variant_name=? AND hatchback.car_name=?", route["variantname"],route["car"]).Scan(&t1)
		//fmt.Println(t1.VariantName, t1.Length)

		t.ExecuteTemplate(w, "showdata.html", t1)
	}else if route["type"]=="sedans" {
		var t1 SedanSpec
		db.Raw("SELECT spec.*, sedan.variant_name FROM spec INNER JOIN sedan ON sedan.s_id = spec.specs_id Where sedan.variant_name=? AND sedan.car_name=?", route["variantname"],route["car"]).Scan(&t1)
		db.Raw("SELECT feature.*, sedan.variant_name FROM feature INNER JOIN sedan ON sedan.f_id = feature.feature_id Where sedan.variant_name=? AND sedan.car_name=?", route["variantname"],route["car"]).Scan(&t1)
		//fmt.Println(t1.VariantName, t1.Length)
          fmt.Println("data is ::",t1)
		t.ExecuteTemplate(w, "showdata.html", t1)
	}
}
