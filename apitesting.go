package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)

type person struct{
	ID string       `json:"Id,omitempty"`
	Name string     `json:name,omitempty`
	Age int      `json:age,omitempty`
	Mobile int      `json:mobile,omitempty`
	Address string  `json:address,omitempty`
	Hobbies []string`json:hobbies,omitempty`
}

var people []person

func main() {

	people=append(people,person{ID:"001",Name:"Aman",Age:21,Mobile:9811399095,Address:"G.Noida",Hobbies:[]string{"Athletics","Reading"}})
	people=append(people,person{ID:"002",Name:"Ambuj",Age:22,Mobile:7876872634,Address:"Haryana",Hobbies:[]string{"Travel","Reading"}})
	people=append(people,person{ID:"003",Name:"Shubham",Age:23,Mobile:9834246435,Address:"Shahadra",Hobbies:[]string{"Khokho","coding"}})
	people=append(people,person{ID:"004",Name:"Mukul",Age:22,Mobile:9325435395,Address:"Rohini",Hobbies:[]string{"TT","Handball"}})
	people=append(people,person{ID:"005",Name:"Kiran",Age:21,Mobile:7653325495,Address:"Pune",Hobbies:[]string{"Cricket","Badminton"}})


	router:=mux.NewRouter()
	router.HandleFunc("/getall",getall).Methods("GET")
	router.HandleFunc("/getperson/{Id}",getperson).Methods("GET")
	router.HandleFunc("/create/{Id}",create).Methods("POST")
	router.HandleFunc("/deleterec/{Id}",deleterec).Methods("DELETE")

	err:=http.ListenAndServe(":8080",router)
	check(err)
}

func getall(w http.ResponseWriter,r *http.Request){
	for _,v:=range people{
		json.NewEncoder(w).Encode(v)
	}

}

func getperson(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)      //gives the route variables of current url  //is a map[string]string

	for _,v:=range people{
		if v.ID==params["Id"]{
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(person{})
}

func create(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	var p1 person
	json.NewDecoder(r.Body).Decode(&p1)
    if p1.ID!=params["Id"]{
    	w.Write([]byte("Enter same ID.....!"))
    	return
	}
	p1.ID=params["Id"]
	people=append(people,p1)

	for _,v:=range people{
		json.NewEncoder(w).Encode(v)
	}

}

func deleterec(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)      //gives the route variables of current url  //is a map[string]string

	for i,v:=range people{
		if v.ID==params["Id"]{
			people=append(people[:i],people[i+1:]...)
			break
		}
	}

	for _,v:=range people{
		json.NewEncoder(w).Encode(v)
	}

}

func check(err error){
	if err!=nil{
		log.Fatal(err)
	}
}
