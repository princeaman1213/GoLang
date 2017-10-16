package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Id     int `json:"id"`
}

type Usercontroller struct{
	session *mgo.Session
}

func Newusercontroller(s *mgo.Session) *Usercontroller{
	return &Usercontroller{s}
}

func getsession() *mgo.Session{
	s,err:=mgo.Dial("mongodb://localhost")
	if err!=nil{
		panic(err)
	}
	return s
}

func main() {

	r:=httprouter.New()
	uc:=Newusercontroller(getsession())
	r.GET("/user/:id",uc.Getuser)
	r.GET("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8000",r)

}

func (uc Usercontroller) Getuser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	//u:=User{Name:"Aman",Gender:"Male",Age:21,Id:"002"}

	id:=p.ByName("id")
	/*if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
		return
	}*/

	//oid:=bson.ObjectIdHex(id)
	u:=User{}
    fmt.Println(id)
	err:=uc.session.DB("mongo1").C("Users").FindId(id).One(&u)
    if err!=nil{
    	w.WriteHeader(http.StatusNotFound)
    	return
	}
    fmt.Println("user Is:",u)
	uj,err:=json.Marshal(u)
	if err!=nil{
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc Usercontroller) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := User{}

	//json.NewDecoder(r.Body).Decode(&u)

	u.Id = 005
	u.Name="Aman"
	u.Age=21
	u.Gender="male"
    fmt.Println("u is :",u)
	err:=uc.session.DB("mongo1").C("Users").Insert(&u)
    if err!=nil{
    	panic(err)
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
	fmt.Println("uj is :", string(uj))
}

func (uc Usercontroller) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: only write code to delete user

	id:=p.ByName("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid:=bson.ObjectIdHex(id)

	err:=uc.session.DB("mongo1").C("Users").RemoveId(oid)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}