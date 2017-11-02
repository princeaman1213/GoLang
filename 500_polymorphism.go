package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Cat struct {
	Id    int
	Name  string
	Toy   Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	Id   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	Id        int
	Name      string
	OwnerId   int
	OwnerType string
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
	db.DropTableIfExists(&Cat{},&Dog{},&Toy{})
	db.CreateTable(&Cat{},&Dog{},&Toy{})

	toy1:=Toy{Id:1,Name:"blocks",OwnerId:1,OwnerType:"dog"}
	db.Create(&toy1)

	toy2:=Toy{Id:2,Name:"ball",OwnerId:2,OwnerType:"dog"}
	db.Create(&toy2)

	toy3:=Toy{Id:3,Name:"aeroplane",OwnerId:3,OwnerType:"cat"}
	db.Create(&toy3)

	dog1:=Dog{Id:1,Name:"Bruno",Toy:toy1}
	db.Create(dog1)
	dog2:=Dog{Id:2,Name:"Tommy",Toy:toy2}
	db.Create(dog2)


	var dog []Dog
	db.Debug().Preload("Town1").Find(&dog)

	for _,r:=range dog{

			fmt.Println(r)

	}

}

