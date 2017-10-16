package main

import (
	"net/http"
	"log"
	//"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"html/template"
	"io"

)

var t *template.Template
var err error
var db *sql.DB
//var people []person

func init(){
	t=template.Must(template.ParseFiles("insertapi.gohtml","getpersonapi.gohtml","getall.gohtml"))
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


	//router:=mux.NewRouter()
	http.HandleFunc("/",index)
	http.HandleFunc("/getall",getall)
	http.HandleFunc("/getperson",getperson)      //get person by id
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/deleterec",deleterec)  //.Methods("DELETE") (how to use this along in browser and not postman)

	err:=http.ListenAndServe(":8000",nil)
	check(err)
}

func index(w http.ResponseWriter,r *http.Request){
	//t.ExecuteTemplate(w,"getall.gohtml",nil)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`<a href="/getall">Getall</a><br>
	<a href="/getperson">GetPerson</a><br>
	<a href="/insert">Insert Record</a><br>
	<a href="/deleterec">Delete Record</a><br>`)
}

func getall(w http.ResponseWriter,r *http.Request){

	rows, err := db.Query(`SELECT * FROM table2;`)
	check(err)
	defer rows.Close()
	// data to be used in query
	var s,str1,str2,str3,str4,str5 string
	s = "RETRIEVED RECORDS:\n"

	/*data:=make([]string,0)
	var i int*/
	// query
	for rows.Next() {
		err=rows.Scan(&str1,&str2,&str3,&str4,&str5)
		check(err)
		s+=str1+"  "+str2+"  "+str3+"  "+str4+"  "+str5+"\t"+"\n"

		/*data=append(data,str1+"  "+str2+"  "+str3+"  "+str4+"  "+str5+"\t"+"\n")
		i++*/
	}
	/*io.Copy(w,strings.NewReader(s))

	for _,v:=range data{
		io.Copy(w,strings.NewReader(v))
	}*/

	w.Header().Set("Content-Type","text/html; charset=utf-8")
	fmt.Fprintln(w,s+`<h1><a href="/">Go to index</a></h1>`)
	//t.ExecuteTemplate(w,"getall.gohtml",nil)
	//io.Copy(w,strings.NewReader(s))
}

func getperson(w http.ResponseWriter,r *http.Request){            //by EmpId

	v1:=r.FormValue("empid")

	//params:=mux.Vars(r)
	rows,err:=db.Query(`SELECT * FROM table2 WHERE EmpId = ?;`,v1)
	check(err)

	//params:=mux.Vars(r)      //gives the route variables of current url  //is a map[string]string

	var s,str1,str2,str3,str4,str5 string


	for rows.Next(){
		s="Selected Record from customer Table :\n"
		err=rows.Scan(&str1,&str2,&str3,&str4,&str5)
		check(err)
		//if str1==params["Id"]{
		s+=str1+"  "+str2+"  "+str3+"  "+str4+"  "+str5+"\n"
		//}
	}

	t.ExecuteTemplate(w,"getpersonapi.gohtml",s)
	//fmt.Fprintln(w,s)
}

func insert(w http.ResponseWriter, r *http.Request) {

	/*http.SetCookie(w,&http.Cookie{
		Name:"insert",
		Value:"insert val",
		Expires:time.Now().Add(time.Second*20),
		HttpOnly:true,
		Path:"/",
		//Domain:"/",
	})*/

	v1:=r.FormValue("empid")
	v2:=r.FormValue("name")
	v3:=r.FormValue("mobile")
	v4:=r.FormValue("address")
	v5:=r.FormValue("position")
     fmt.Println(v1,v2,v3,v4,v5)               //checking
     s:="INSERTED RECORD"
	//params:=mux.Vars(r)

	/*if v1==""{
		return
	}*/

	/*rows1,err1:=db.Query(`SELECT * FROM table2 WHERE EmpId = ?;`,v1)
	check(err1)
	if rows1.Next(){
		fmt.Fprintln(w,"Duplicate EmpId")
		return
	}*/

	query:="INSERT INTO table2(EmpId,Name,Mobile,Address,Position) VALUES(?,?,?,?,?);"
	stmt,err:=db.Prepare(query)
	check(err)

	exe, err := stmt.Exec(v1,v2, v3,v4,v5)
	check(err)

	t.ExecuteTemplate(w,"insertapi.gohtml",nil)

	n, err := exe.RowsAffected()
	check(err)

		fmt.Fprintln(w,s,n)

}

func deleterec(w http.ResponseWriter,r *http.Request){
	v1:=r.FormValue("empid")
	//params:=mux.Vars(r)

	_, err := db.Query(`DELETE FROM table2 WHERE EmpId=?;`,v1)
	check(err)
	//defer stmt.Close()

	//r1, err := stmt.Exec()
	//check(err)

	//n, err := r1.RowsAffected()
	//check(err)
	t.ExecuteTemplate(w,"getpersonapi.gohtml",nil)
	fmt.Fprintln(w, "DELETED RECORD")

}

func check(err error){
	if err!=nil{
		log.Fatal(err)
	}
}
