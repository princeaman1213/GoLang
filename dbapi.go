package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"html/template"
	"io"
)

/*type person struct{
	ID string       `json:"Id,omitempty"`
	Name string     `json:name,omitempty`
	Age int      `json:age,omitempty`
	Mobile int      `json:mobile,omitempty`
	Address string  `json:address,omitempty`
	Hobbies []string`json:hobbies,omitempty`
}*/
var t *template.Template
var err error
var db *sql.DB
//var people []person

func init(){
	t=template.Must(template.ParseFiles("getall.gohtml"))
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


	/*people=append(people,person{ID:"001",Name:"Aman",Age:21,Mobile:9811399095,Address:"G.Noida",Hobbies:[]string{"Athletics","Reading"}})
	people=append(people,person{ID:"002",Name:"Ambuj",Age:22,Mobile:7876872634,Address:"Haryana",Hobbies:[]string{"Travel","Reading"}})
	people=append(people,person{ID:"003",Name:"Shubham",Age:23,Mobile:9834246435,Address:"Shahadra",Hobbies:[]string{"Khokho","coding"}})
	people=append(people,person{ID:"004",Name:"Mukul",Age:22,Mobile:9325435395,Address:"Rohini",Hobbies:[]string{"TT","Handball"}})
	people=append(people,person{ID:"005",Name:"Kiran",Age:21,Mobile:7653325495,Address:"Pune",Hobbies:[]string{"Cricket","Badminton"}})
*/

	router:=mux.NewRouter()
	router.HandleFunc("/",index)
	router.HandleFunc("/getall",getall).Methods("GET")
	router.HandleFunc("/getperson/{Id}",getperson).Methods("GET")      //get person by id
	router.HandleFunc("/insert/{Id}", insert)
	router.HandleFunc("/deleterec/{Id}",deleterec)  //.Methods("DELETE") (how to use this along in browser and not postman)
	//router.HandleFunc("/update/{Id}",deleterec)

	err:=http.ListenAndServe(":8080",router)
	check(err)
}

func index(w http.ResponseWriter,r *http.Request){
	//t.ExecuteTemplate(w,"getall.gohtml",nil)
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`<a href="/getall">Getall</a><br>`)
}

func getall(w http.ResponseWriter,r *http.Request){


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

func getperson(w http.ResponseWriter,r *http.Request){


	/*for _,v:=range people{
		if v.ID==params["Id"]{
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(person{})*/
	params:=mux.Vars(r)
	rows,err:=db.Query(`SELECT * FROM table1 WHERE EmpId = ?;`,params["Id"])
	check(err)

	//params:=mux.Vars(r)      //gives the route variables of current url  //is a map[string]string

	var s,str1,str2,str3,str4,str5 string
	s="Selected Record from customer Table :\n"

	for rows.Next(){
		err=rows.Scan(&str1,&str2,&str3,&str4,&str5)
		check(err)
		//if str1==params["Id"]{
			s+=str1+"  "+str2+"  "+str3+"  "+str4+"  "+str5+"\n"
		//}
	}

	fmt.Fprintln(w,s)

}

func insert(w http.ResponseWriter, r *http.Request) {
     params:=mux.Vars(r)
	_,err:=db.Query(`INSERT INTO table1 VALUES(?,"Ambuj","676738673","delhi","Developer");`,params["Id"])
	check(err)

	//exe, err := stmt.Exec()
	//check(err)

	//n, err := exe.RowsAffected()
	//check(err)

	fmt.Fprintln(w, "INSERTED RECORD")

}

func deleterec(w http.ResponseWriter,r *http.Request){
	//params:=mux.Vars(r)      //gives the route variables of current url  //is a map[string]string
//id:=params["Id"]
	/*for i,v:=range people{
		if v.ID==params["Id"]{
			people=append(people[:i],people[i+1:]...)
			break
		}
	}

	for _,v:=range people{
		json.NewEncoder(w).Encode(v)
	}*/
	params:=mux.Vars(r)

	_, err := db.Query(`DELETE FROM table1 WHERE EmpId=?;`,params["Id"])
	check(err)
	//defer stmt.Close()

	//r1, err := stmt.Exec()
	//check(err)

	//n, err := r1.RowsAffected()
	//check(err)

	fmt.Fprintln(w, "DELETED RECORD")

}

func check(err error){
	if err!=nil{
		log.Fatal(err)
	}
}
