package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Place struct {
	gorm.Model
	Name        string
	Town        Town           //    `gorm:"ForeignKey:TownID"`
	TownID      int             //Place belongs to Town
}

type Town struct {
	gorm.Model
	Name string
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
	db.DropTableIfExists(&Place{},&Town{})
	db.CreateTable(&Place{},&Town{})

	place:=Place{Name:"NOIDA",Town:Town{Name:"GBN"}}
	db.Create(&place)
	place1:=Place{Name:"DELHI",Town:Town{Name:"shahadra"}}
	db.Create(&place1)

	/*town:=Town{Name:"gbn"}
	db.Create(&town)
	town1:=Town{Name:"shahadra"}
	db.Create(&town1)*/

	//var user2 User
	//db.Find(&user2)
	//for i, _ := range user2 {
	//db.Model(&user).Related(&profile)
	//}
	var places []Place
	db.Debug().Preload("Town").Find(&places)
	//db.Debug().Model(&user).Related(&profile)
	//db.Debug().Raw("SELECT place.name, town.name FROM place INNER JOIN town ON town.id = place.town_id").Scan(&places)
	fmt.Println(places)
	//db.Model(&)
	/*var p []Student
	db.Debug().Find(&p,"first_name=?","Aman")                            // remember Normal MySQL
    for _,v:=range p{
		fmt.Println(v)

	}*/
}

