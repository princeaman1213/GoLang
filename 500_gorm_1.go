package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
)

type Owner struct{
	gorm.Model
	FirstName string
	LastName string
	Books []Book
}

type Book struct{
	gorm.Model
	Name string
	PublishDate time.Time
	OwnerID uint   `sql:"index"`
	Authors []Author  `gorm:"many2many:books_authors"`
}

type Author struct{
	gorm.Model
	Firstname string
	LastName string

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
	//db.DropTableIfExists(&Owner{},&Book{},&Author{})
	//db.CreateTable(&Owner{},&Book{},&Author{})

	//owner:=Owner{FirstName:"Aman",LastName:"Patel"}
	//db.Create(&owner)
	//owner1:=Owner{FirstName:"Aman",LastName:"Garg"}
	//db.Create(&owner1)
	//owner.FirstName="Abc"
	//db.Debug().Save(&owner)
	//db.Debug().Delete(&owner)
	//db.Model(&Owner{}).Update("first_name", "new-name")
	//owner1:=Owner{FirstName:"Shubham",LastName:"Garg"}
	//db.Create(&owner1)
	var owner1 []Owner

	//db.Debug().Model(owner1).Where("id=?",1).Update("first_name","Akash")

	//db.Where("first_name = ?", "Aman").First(&owner1)                         //use var owner1 Owner
	db.Where("first_name in (?)", []string{"Aman", "Avinash"}).Find(&owner1)     //use var owner1 []Owner
	fmt.Println(owner1)

}

