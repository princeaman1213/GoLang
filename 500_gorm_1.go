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
	db.Create(&owner1)

	var owner3 []Owner
	//db.Debug().Set("gorm:query_option", "FOR UPDATE").First(&owner3, 1)      //used to fetch the record

	//db.Find(&owner3)

  //  Search(db).Find(&owner3)                       // Search func is used here to apply additional conditions
	//fmt.Println(owner3)

	db.Scopes(Search).Find(&owner3)                //Scope is used here to apply additional conditions
	fmt.Println(owner3)

	//db.Debug().Model(&Owner{}).ModifyColumn("first_name", "CHAR(22)")

	//owner.FirstName="Abc"
	//db.Debug().Save(&owner)
	//db.Debug().Delete(&owner)
	//db.Model(&Owner{}).Update("first_name", "new-name")
	//owner1:=Owner{FirstName:"Shubham",LastName:"Garg"}
	//db.Create(&owner1)
	//var owner1 []Owner

	/*tx := db.Begin()
	tx1:=db.Begin()
	owner:=Owner{FirstName:"Raman",LastName:"Bindal"}
	err=tx.Create(&owner).Error

	if err!=nil{
		tx.Rollback()
	}else {
		tx.Commit()
		owner:=Owner{FirstName:"Nirbheek",LastName:"Banga"}
		err=tx1.Create(&owner).Error
		if err!=nil{
			tx1.Rollback()
		}else{
			tx1.Commit()
		}
	}*/


	//db.Debug().Model(owner1).Where("id=?",1).Update("first_name","Akash")

	//db.Where("first_name = ?", "Aman").First(&owner1)                         //use var owner1 Owner
	//db.Where("first_name in (?)", []string{"Aman", "Avinash"}).Find(&owner1)     //use var owner1 []Owner
	//fmt.Println(owner1)

}

func Search(db *gorm.DB) *gorm.DB {
	return db.Where("id = ?", 2)
}
