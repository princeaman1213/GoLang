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

/*var db *gorm.DB
var err error*/

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
	db.DropTableIfExists(&Owner{},&Book{},&Author{})
	db.CreateTable(&Owner{},&Book{},&Author{})

	owner:=Owner{FirstName:"Aman",LastName:"Patel"}
	db.Create(&owner)
	owner1:=Owner{FirstName:"Shubham",LastName:"Garg"}
	//db.Create(&owner1)

	//db.Debug().Model(&Owner{}).Where("id = ?",1).Update("last_name","Patel")   //updating with callbacks
	//db.Model(&Owner{}).Where("id = ?",1).Updates(map[string]interface{}{"first_name": "hello","last_name":"newname"})
	//db.Debug().Model(&Owner{}).Where("id = ?",1).UpdateColumn("first_name", "hello")         //without callbaks

	//db.Debug().Model(&Owner{}).Where("id = ?",2).Delete(&Owner{})                      //deleting with callback

	tx:=db.Begin()                                          //transaction
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	err1:=tx.Create(&owner1).Error

	if err1!=nil{
		tx.Rollback()
		fmt.Println(tx.Error)
	}else{
		tx.Commit()
	}


	var owner3 []Owner

	db.Scopes(Search).Find(&owner3)                //Scope is used here to apply additional conditions
	fmt.Println(owner3)


}

func Search(db *gorm.DB) *gorm.DB {
	return db.Where("id = ?", 1)
}
