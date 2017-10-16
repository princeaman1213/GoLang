package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Place1 struct {
	gorm.Model
	Name        string
	Town1        []Town1             `gorm:"many2many:place1_town1;"`
	//Town1ID      int
}

type Town1 struct {
	gorm.Model
	Name string
	Place1 []Place1                   `gorm:"many2many:place1_town1;"`
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
	db.DropTableIfExists(&Place1{},&Town1{})
	db.CreateTable(&Place1{},&Town1{})

	t1:=Town1{Name:"gbn"}
	t2:=Town1{Name:"sec-62"}

	place:=Place1{Name:"NOIDA",Town1:[]Town1{t1,t2}}
	db.Create(&place)
	place1:=Place1{Name:"DELHI",Town1:[]Town1{{Name:"shahadra"}}}
	db.Create(&place1)
	place2:=Place1{Name:"Mumbai",Town1:[]Town1{{Name:"pune"}}}
	db.Create(&place2)
	/*town:=Town1{Name:"gbn"}
	db.Create(&town)
	town1:=Town1{Name:"shahadra"}
	db.Create(&town1)*/

	//var user2 User
	//db.Find(&user2)
	//for i, _ := range user2 {
	//db.Model(&user).Related(&profile)
	//}
	var places []Place1
	db.Debug().Preload("Town1").Find(&places)
	//db.Debug().Model(&user).Related(&profile)
	//db.Debug().Raw("SELECT place.name, town.name FROM place INNER JOIN town ON town.id = place.town_id").Scan(&places)
	for _,r:=range places{
		for _,r1:=range r.Town1{
			fmt.Println(r.ID,r.Name,r1.Name)
		}
	}

	//db.Model(&)
	/*var p []Student
	db.Debug().Find(&p,"first_name=?","Aman")                            // remember Normal MySQL
    for _,v:=range p{
		fmt.Println(v)

	}*/
}

