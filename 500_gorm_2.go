package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Student struct{
	gorm.Model
	FirstName string
	LastName string
	//Notes Notes
}

type Notes struct{
	gorm.Model
	Name string
	Author string
	OwnerID uint
}

func main() {

	db,err:=gorm.Open("mysql","root:password@/gorm_db?charset=utf8&parseTime=True&loc=Local")
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

	db.SingularTable(true)
	db.DropTableIfExists(&Student{},&Notes{})
	db.CreateTable(&Student{},&Notes{})

	student:=Student{FirstName:"Aman",LastName:"Patel"}
	db.Create(&student)
	student1:=Student{FirstName:"Aman",LastName:"Garg"}
	db.Create(&student1)

	notes:=Notes{Name:"Rick",Author:"narakson",OwnerID:1}
	db.Create(&notes)
	notes1:=Notes{Name:"LeBlanc",Author:"daffo",OwnerID:1}
	db.Create(&notes1)

	db.Model(&Student{}).Related(&Notes{})
	/*var p []Student
	db.Debug().Find(&p,"first_name=?","Aman")                            // remember Normal MySQL
    for _,v:=range p{
		fmt.Println(v)

	}*/
}

