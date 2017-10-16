package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("postgres", "postgres://aman:password@localhost/test1?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	http.HandleFunc("/books",books)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8000",nil)

}

func books(w http.ResponseWriter,r *http.Request){
	isbn:="Jayne Austen"
	rows, err := db.Query("SELECT * FROM books WHERE author = $1",isbn)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks:=make([]Book,0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}

	for _,v:=range bks{
		fmt.Fprint(w,v.isbn)
		fmt.Fprint(w,v.title)
		fmt.Fprint(w,v.author)
		fmt.Fprint(w,v.price)
		fmt.Fprintln(w,"")
	}

}
