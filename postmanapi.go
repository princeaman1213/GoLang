package main

import
(
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_"encoding/json"
	_"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
)

var db *sql.DB
var err error

type UserDetail struct{
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
	Address string `json:"address,omitempty"`
}

func InsertData(w http.ResponseWriter, req *http.Request){

 req.Header.Set("Content-Type","application/json")
  var user UserDetail
	_ = json.NewDecoder(req.Body).Decode(&user)
	fmt.Println(user)
	_,err:=db.Exec(`INSERT INTO emptable VALUES(?,?,?,?)`,user.ID,user.Name,user.Position,user.Address)
	checkError(err)
}



func DeleteData(w http.ResponseWriter, req *http.Request){
	req.Header.Set("Content-Type","application/json")
	var id string
	_ = json.NewDecoder(req.Body).Decode(&id)
	_,err:= db.Exec(`DELETE FROM emptable WHERE ID = ?`,id)
	checkError(err)

}


func FetchData(w http.ResponseWriter, req *http.Request){
	req.Header.Set("Content-Type","application/json")
	var user UserDetail
	params:=mux.Vars(req)
	res,err:= db.Query(`SELECT * FROM emptable WHERE ID = ? `,params["id"])
	checkError(err)
	for res.Next(){
		err:= res.Scan(&user.ID,&user.Name,&user.Position,&user.Address)
		checkError(err)
	}

		json.NewEncoder(w).Encode(user)
}



func UpdateData(w http.ResponseWriter, req *http.Request){
	req.Header.Set("Content-Type","application/json")
	var user UserDetail
	_ = json.NewDecoder(req.Body).Decode(&user)

	_,err:= db.Exec(`UPDATE emptable SET NAME = ?, POSITION = ?, ADDRESS = ? WHERE ID = ?`,user.Name,user.Position,user.Address,user.ID)
	checkError(err)


}


func checkError(e error){
	if e!=nil{
		log.Fatalln(e)
	}
}


func main(){
	//Connecting Database
	db,err = sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/test1")
	checkError(err)
	defer db.Close()
	err = db.Ping()
	checkError(err)

	//Routing URLs
	router := mux.NewRouter()
	router.HandleFunc("/admin",InsertData).Methods("POST")
	router.HandleFunc("/admin/{id}",FetchData).Methods("GET")
	router.HandleFunc("/admin",UpdateData).Methods("PUT")
    router.HandleFunc("/admin",DeleteData).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",router))

}
