package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"fmt"
	"io"
)

var err error
var db *sql.DB

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	db,err=sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/test1")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/rec", records)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func records(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT Name FROM table1;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {

	stmt,err:=db.Prepare(`CREATE TABLE customer1(name VARCHAR(10),age INT(2));`)  //customer table (name varchar) is also there
	check(err)
	defer stmt.Close()

	//Exec executes a prepared statement with the given arguments
	exe,err:=stmt.Exec()
	check(err)

	//RowsAffected returns the number of rows affected by an
	// update, insert, or delete
	n,err:=exe.RowsAffected()
	check(err)

	fmt.Fprintln(w,"Created Table !",n)

}

func insert(w http.ResponseWriter, req *http.Request) {

	stmt,err:=db.Prepare(`INSERT INTO table1 VALUES("ID3","James",676738673,"delhi","Associate");`)
	check(err)

	exe, err := stmt.Exec()
	check(err)

	n, err := exe.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)

}

func read(w http.ResponseWriter, req *http.Request) {

	rows,err:=db.Query(`SELECT * FROM table1;`)
	check(err)

	var s,str1,str2,str3,str4,str5 string
	s="Records from customer Table :\n"

	for rows.Next(){
		err=rows.Scan(&str1,&str2,&str3,&str4,&str5)
		check(err)
		s+=str1+"  "+str2+"  "+str3+"  "+str4+"  "+str5+"\n"
	}

	fmt.Fprintln(w,s)

}

func update(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`UPDATE customer1 SET name="Jimmy" WHERE name="James";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)

}

func delete(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`DELETE FROM customer1 WHERE name="jimmy";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)

}

func drop(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")

}
