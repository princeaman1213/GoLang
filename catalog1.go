package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"github.com/jinzhu/gorm"
	"fmt"
	"net/http"
	"html/template"
	"strings"
)


type Hatchback struct {
	HatchID      uint     `gorm:"primary_key"`
	CompanyName  string
	CarName  string
	VariantName  string
	S_ID          uint
	F_ID          uint
}

type Sedan struct {
	SedanID      uint     `gorm:"primary_key"`
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
	Hatchback
	Spec
	Feature
}

type SedanSpec struct{
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
	t=template.Must(template.ParseFiles("catindex.html","typeofvec.html","routecompany.html","routecar.html","showdata.html"))
	//thv=template.Must(template.ParseFiles("datsun.html","tata.html","chevrolet.html",""))
	//tsv=template.Must(template.ParseFiles("mahindraverito.html","nissansunny.html","fiatlinea.html","hyundaiverna.html","skodarapid.html"))

}

func main() {
	db,err=gorm.Open("mysql","root:password@/cat_db?charset=utf8&parseTime=True&loc=Local")
	//db,err:=gorm.Open("postgres","user=aman password=password dbname=test1 sslmode=disable")

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


    http.HandleFunc("/",catindex)
	//http.Handle("/car.jpeg",http.StripPrefix("/",http.FileServer(http.Dir("./"))))

	http.HandleFunc("/type",typeofvec)
	//http.Handle("/hatchback.jpg",http.StripPrefix("/",http.FileServer(http.Dir("./"))))
	//http.Handle("/sedan.png",http.StripPrefix("/",http.FileServer(http.Dir("./"))))

	http.HandleFunc("/hatchbacks",carcompany)
	http.HandleFunc("/sedans",carcompany)

	//http.HandleFunc("/getdata",getdata)

	//hatchbags
	http.HandleFunc("/Datsun",carname)
	http.HandleFunc("/Tata",carname)
	http.HandleFunc("/Chevrolet",carname)
	http.HandleFunc("/Maruti Suzuki",carname)

	//sedans
	http.HandleFunc("/Audi",carname)
	http.HandleFunc("/Hyundai",carname)
	http.HandleFunc("/Skoda",carname)
	http.HandleFunc("/Jaguar",carname)

	//Datsun
	http.HandleFunc("/Redi-Go",variantname)
	http.HandleFunc("/Sport",showdata)
	http.HandleFunc("/Gold Limited Edition",showdata)

	http.HandleFunc("/Go",variantname)
	http.HandleFunc("/Anniversary Edition",showdata)
	http.HandleFunc("/A EPS",showdata)

	//Tata
	http.HandleFunc("/Tiago",variantname)
	http.HandleFunc("/Revotron XB",showdata)
	http.HandleFunc("/Revotron XE",showdata)

	http.HandleFunc("/Nano",variantname)
	http.HandleFunc("/XM",showdata)

	//Chevrolet
	http.HandleFunc("/Beat",variantname)
	http.HandleFunc("/PS Petrol",showdata)
	http.HandleFunc("/LS Diesel",showdata)

	http.HandleFunc("/Sail",variantname)
	http.HandleFunc("/1.2 Base",showdata)
	http.HandleFunc("/1.2 LS",showdata)

	//Maruti Suzuki
	http.HandleFunc("/Alto K10",variantname)
	http.HandleFunc("/LX",showdata)
	http.HandleFunc("/VXi",showdata)

	http.HandleFunc("/Celerio",variantname)
	http.HandleFunc("/Zxi",showdata)
	http.HandleFunc("/Vxi AMT",showdata)

	//Audi
	http.HandleFunc("/A3",variantname)
	http.HandleFunc("/35 TFSI Premium Plus",showdata1)
	http.HandleFunc("/35 TDI Premium Plus",showdata1)

	http.HandleFunc("/A4",variantname)
	http.HandleFunc("/35 TDI",showdata1)
	http.HandleFunc("/30 TFSI Technology Pack",showdata1)

	//Hyundai
	http.HandleFunc("/Verna",variantname)
	http.HandleFunc("/1.6 VTVE",showdata1)
	http.HandleFunc("/1.6 CRDI E",showdata1)

	http.HandleFunc("/Elantra",variantname)
	http.HandleFunc("/2.0 S MT",showdata1)
	http.HandleFunc("/2.0 SX MT",showdata1)

	//Skoda
	http.HandleFunc("/Rapid",variantname)
	http.HandleFunc("/Active 1.6 MPI",showdata1)
	http.HandleFunc("/Monte Carlo 1.5 TDI MT",showdata1)

	http.HandleFunc("/Octavia",variantname)
	http.HandleFunc("/Ambition 1.6 MPI",showdata1)
	http.HandleFunc("/1.4 TSI Ambition",showdata1)

	//Jaguar
	http.HandleFunc("/XE",variantname)
	http.HandleFunc("/Pure",showdata1)
	http.HandleFunc("/Portfolio",showdata1)

	http.HandleFunc("/XF",variantname)
	http.HandleFunc("/Prestige Petrol",showdata1)
	http.HandleFunc("/Prestige Diesel",showdata1)

	http.Handle("/favicon.ico",http.NotFoundHandler())

	http.ListenAndServe(":8000",nil)

}

func catindex(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	t.ExecuteTemplate(w,"catindex.html",nil)
}

func typeofvec(w http.ResponseWriter,r *http.Request){
	t.ExecuteTemplate(w,"typeofvec.html",nil)
}

func carcompany(w http.ResponseWriter,r *http.Request){
	c:=fmt.Sprint(r.URL)
	var route string
	//temp:="still in process...!"
	route=string(c)[1:]
    fmt.Println(route)

    if strings.ToLower(route)=="hatchbacks" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT company_name").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.CompanyName)
		}
		t.ExecuteTemplate(w,"routecompany.html",names)

	}else if strings.ToLower(route)=="sedans"{
		var s []Sedan
		db.Debug().Table("sedan").Select("DISTINCT company_name").Find(&s)

		var names = []string{}
		for _,v:=range s{
			names = append(names,v.CompanyName)
		}
		t.ExecuteTemplate(w,"routecompany.html",names)
	}

}

func carname(w http.ResponseWriter,r *http.Request){

	c:=fmt.Sprint(r.URL)
	var route string
	//temp:="still in process...!"
	route=string(c)[1:]
	fmt.Println(route)

	if strings.ToLower(route)=="datsun" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT car_name").Where("company_name=?","Datsun").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)

	}else if strings.ToLower(route)=="tata"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT car_name").Where("company_name=?","Tata").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if strings.ToLower(route)=="chevrolet"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT car_name").Where("company_name=?","Chevrolet").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if strings.ToLower(route)=="maruti%20suzuki"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT car_name").Where("company_name=?","Maruti Suzuki").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if strings.ToLower(route)=="audi"{
		var s []Sedan
		db.Debug().Table("sedan").Select("DISTINCT car_name").Where("company_name=?","Audi").Find(&s)

		var names = []string{}
		for _,v:=range s{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if strings.ToLower(route)=="hyundai"{
		var s []Sedan
		db.Debug().Table("sedan").Select("DISTINCT car_name").Where("company_name=?","Hyundai").Find(&s)

		var names = []string{}
		for _,v:=range s{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if strings.ToLower(route)=="skoda"{
		var s []Sedan
		db.Debug().Table("sedan").Select("DISTINCT car_name").Where("company_name=?","Skoda").Find(&s)

		var names = []string{}
		for _,v:=range s{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if strings.ToLower(route)=="jaguar"{
		var s []Sedan
		db.Debug().Table("sedan").Select("DISTINCT car_name").Where("company_name=?","Jaguar").Find(&s)

		var names = []string{}
		for _,v:=range s{
			names = append(names,v.CarName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}

}


func variantname(w http.ResponseWriter,r *http.Request){

	c:=fmt.Sprint(r.URL)
	var route string
	//temp:="still in process...!"
	route=string(c)[1:]
	fmt.Println(route)

	if route=="Redi-Go" {
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Redi-Go").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)

	}else if route=="Go"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Go").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Tiago"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Tiago").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Nano"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Nano").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Beat"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Beat").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Sail"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Sail").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Alto%20K10"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Alto K10").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Celerio"{
		var h []Hatchback
		db.Debug().Table("hatchback").Select("DISTINCT variant_name").Where("car_name=?","Celerio").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}

	if route=="A3" {
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","A3").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)

	}else if route=="A4"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","A4").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Verna"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","Verna").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Elantra"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","Elantra").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Rapid"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","Rapid").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="Octavia"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","Octavia").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="XE"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","XE").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}else if route=="XF"{
		var h []Sedan
		db.Debug().Table("sedan").Select("DISTINCT variant_name").Where("car_name=?","XF").Find(&h)

		var names = []string{}
		for _,v:=range h{
			names = append(names,v.VariantName)
		}
		t.ExecuteTemplate(w,"routecar.html",names)
	}

}

func showdata(w http.ResponseWriter,r *http.Request){
	/*fmt.Println(r.Referer())
	fmt.Println(strings.SplitAfter(r.Referer(),"/")[3])
	fmt.Println(len(strings.SplitAfter(r.Referer(),"/")))*/

	c:=fmt.Sprint(r.URL)

	var route string
	//temp:="still in process...!"
	route=string(c)[1:]
	fmt.Println(route)

	r1:=strings.Replace(route,"%20"," ",-1)   //to replace that %20 back to " " (space)
	//fmt.Println("r1     ",r1)

	/*rows, err := db.Debug().Table("spec").Select(" spec.*,hatchback.variant_name").Joins("inner join hatchback on spec.specs_id = hatchback.s_id").Where("hatchback.variant_name=?",r1).Rows()
	if err!=nil{
		fmt.Fprintln(w,"Error Fetching data")
		fmt.Println("Error Fetching data")
	}
	for rows.Next() {
		rows.Scan(&t1.VariantName, &t1.Length)


	}*/

	var t1 HatchSpec
	db.Raw("SELECT spec.*, hatchback.variant_name FROM spec INNER JOIN hatchback ON hatchback.s_id = spec.specs_id Where hatchback.variant_name=?",r1 ).Scan(&t1)

	fmt.Println(t1.VariantName,t1.Length)

	//var temp =[2]string{"asda","dadd"}
	t.ExecuteTemplate(w,"showdata.html",t1)

}


func showdata1(w http.ResponseWriter,r *http.Request){
	/*fmt.Println(r.Referer())
	fmt.Println(strings.SplitAfter(r.Referer(),"/")[3])
	fmt.Println(len(strings.SplitAfter(r.Referer(),"/")))*/

	c:=fmt.Sprint(r.URL)

	var route string
	//temp:="still in process...!"
	route=string(c)[1:]
	fmt.Println(route)

	r1:=strings.Replace(route,"%20"," ",-1)   //to replace that %20 back to " " (space)
	//fmt.Println("r1     ",r1)

	/*rows, err := db.Debug().Table("spec").Select(" spec.*,hatchback.variant_name").Joins("inner join hatchback on spec.specs_id = hatchback.s_id").Where("hatchback.variant_name=?",r1).Rows()
	if err!=nil{
		fmt.Fprintln(w,"Error Fetching data")
		fmt.Println("Error Fetching data")
	}
	for rows.Next() {
		rows.Scan(&t1.VariantName, &t1.Length)


	}*/

	var t1 SedanSpec
	db.Debug().Raw("SELECT spec.*, sedan.variant_name FROM spec INNER JOIN sedan ON sedan.s_id = spec.specs_id Where sedan.variant_name=?",r1 ).Scan(&t1)

	fmt.Println(t1.VariantName,t1.Length)

	//var temp =[2]string{"asda","dadd"}
	t.ExecuteTemplate(w,"showdata.html",t1)

}
